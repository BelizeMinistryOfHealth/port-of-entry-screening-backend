package fns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"bz.epi.covid.screen/arrivals/domain"
	"bz.epi.covid.screen/arrivals/persistence"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

type ArrivalList struct {
	Arrivals []domain.Arrival `json:"arrivals"`
	NextPage string           `json:"next_page"`
}

func doAuth(clientId, secret string, r *http.Request) error {
	if len(clientId) == 0 && len(secret) == 0 {
		return fmt.Errorf("unauthorized: no credentials provided")
	}

	if secret != os.Getenv(strings.ToTitle(clientId)) {
		return fmt.Errorf("unauthorized: wrong credentials provided")
	}
	return nil
}

func Arrivals(w http.ResponseWriter, r *http.Request) {
	projectId := os.Getenv("PROJECT_ID")
	collection := os.Getenv("DB_COLLECTION")
	clientId := r.Header.Get("client_id")
	secret := r.Header.Get("client_secret")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, client_id, client_secret")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	log.WithFields(log.Fields{
		"headers": r.Header,
	}).Info("Got a request for arrivals")

	if len(collection) == 0 {
		http.Error(w, "could not find the collection", http.StatusInternalServerError)
		return
	}
	dbClient := persistence.CreateClient(r.Context(), projectId)
	switch method := r.Method; method {
	case http.MethodPost:
		err := doAuth(clientId, secret, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		_, err = UpsertArrival(dbClient, collection, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "ok")
	case http.MethodGet:
		err := doAuth(clientId, secret, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		arrivals, err := ListArrivals(dbClient, collection, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var nextPage string
		if len(arrivals) == 25 {
			nextPage = arrivals[len(arrivals)-1].Id
		}
		arrivalList := ArrivalList{
			Arrivals: arrivals,
			NextPage: nextPage,
		}
		resp, err := json.Marshal(arrivalList)
		if err != nil {
			http.Error(w, "marshalling response failed", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, string(resp))
	default:
		fmt.Fprintf(w, "ok")
	}

}

func ListArrivals(dbClient *persistence.FirestoreClient, collection string, r *http.Request) ([]domain.Arrival, error) {
	layout := "2006-01-02"
	now := time.Now()
	after := now.Format(layout)

	//after := r.URL.Query()["next_page"][0]
	if len(r.URL.Query().Get("next_page")) > 0 {
		after = r.URL.Query().Get("next_page")
	}

	// Verify that the format is `YYYY-MM-DD`
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if !re.MatchString(after) {
		return nil, fmt.Errorf("wrong format provided for `next_page. Expected format `YYYY-MM-DD`")
	}

	cursor, err := url.QueryUnescape(after)
	if err != nil {
		return nil, fmt.Errorf("could not url decode the query string")
	}

	arrivals, err := dbClient.List(collection, cursor)

	if err != nil {
		return nil, fmt.Errorf("error retrieving arrivals: %v", err)
	}

	return arrivals, nil
}

func UpsertArrival(dbClient *persistence.FirestoreClient, collection string, r *http.Request) ([]domain.Arrival, error) {
	var newArrivals []domain.NewArrival

	// Read body
	body := r.Body
	b, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not parse the body posted: %v", err)
	}

	err = json.Unmarshal(b, &newArrivals)
	if err != nil {
		log.WithFields(log.Fields{
			"body":  string(b),
			"error": err,
		}).Error("could not unmarshall the body posted")
		return nil, fmt.Errorf("could not unmarshall the body posted: %v", err)
	}

	arrivals := domain.HydrateCompanions(newArrivals)

	err = dbClient.UpsertArrival(collection, arrivals)
	if err != nil {
		return nil, fmt.Errorf("error saving newArrivals information: %v", err)
	}

	return arrivals, nil

}

func FindByPortOfEntry(w http.ResponseWriter, r *http.Request) {
	projectId := os.Getenv("PROJECT_ID")
	collection := os.Getenv("DB_COLLECTION")
	clientId := r.Header.Get("client_id")
	secret := r.Header.Get("client_secret")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, client_id, client_secret")

	if r.Method != http.MethodGet {
		return
	}

	if len(collection) == 0 {
		http.Error(w, "could not find the collection", http.StatusInternalServerError)
		return
	}

	if len(clientId) == 0 && len(secret) == 0 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if secret != os.Getenv(strings.ToTitle(clientId)) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the query string
	// And exit if the query string for port of entry is not present
	q := r.URL.Query()
	portEntry := q["port_entry"][0]

	dbClient := persistence.CreateClient(r.Context(), projectId)

	arrivals, err := dbClient.FindByPortOfEntry(collection, portEntry)

	if err != nil {
		http.Error(w, "query failed", http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(arrivals)
	if err != nil {
		http.Error(w, "marshalling response failed", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(resp))
}

type FindByNameReq struct {
	Name string `json:"name"`
}

func FindByName(w http.ResponseWriter, r *http.Request) {
	projectId := os.Getenv("PROJECT_ID")
	collection := os.Getenv("DB_COLLECTION")
	clientId := r.Header.Get("client_id")
	secret := r.Header.Get("client_secret")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, client_id, client_secret")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodPost {
		return
	}

	if len(collection) == 0 {
		http.Error(w, "could not find the collection", http.StatusInternalServerError)
		return
	}

	if len(clientId) == 0 && len(secret) == 0 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if secret != os.Getenv(strings.ToTitle(clientId)) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "could not parse the body posted", http.StatusInternalServerError)
		return
	}

	var nameReq FindByNameReq
	json.Unmarshal(b, &nameReq)

	dbClient := persistence.CreateClient(r.Context(), projectId)

	names := strings.Split(strings.Trim(nameReq.Name, ""), " ")
	var arrivals []domain.Arrival

	if len(names) == 1 {
		// If only one name is provided, we assume it is the last name
		arrivals, err = dbClient.FindByLastName(collection, strings.Title(strings.Trim(nameReq.Name, "")))
	}

	if len(names) == 2 {
		// If both names are provided we search for the full name
		arrivals, err = dbClient.FindByFullName(collection, strings.Title(strings.Trim(nameReq.Name, "")))
	}

	if err != nil {
		http.Error(w, "query failed", http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(arrivals)
	if err != nil {
		http.Error(w, "marshalling response failed", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(resp))
}
