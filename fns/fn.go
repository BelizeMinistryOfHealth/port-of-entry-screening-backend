package fns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"bz.epi.covid.screen/arrivals/domain"
	"bz.epi.covid.screen/arrivals/persistence"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func Ingest(w http.ResponseWriter, r *http.Request) {
	projectId := os.Getenv("PROJECT_ID")
	collection := os.Getenv("DB_COLLECTION")
	clientId := r.Header.Get("client_id")
	secret := r.Header.Get("client_secret")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, client_id, client_secret")

	if r.Method != http.MethodPost {
		return
	}

	if len(collection) == 0 {
		http.Error(w, "could not find the collection", http.StatusInternalServerError)
		return
	}

	var arrivals []domain.Arrival
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "could not parse the body posted", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &arrivals)
	if err != nil {
		log.WithFields(log.Fields{
			"body":  string(b),
			"error": err,
		}).Error("could not unmarshall the body posted")
		http.Error(w, "could not unmarshall the body posted", http.StatusInternalServerError)
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

	dbClient := persistence.CreateClient(r.Context(), projectId)

	err = dbClient.UpsertArrival(collection, arrivals)
	if err != nil {
		http.Error(w, "error saving arrivals information", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "ok")
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
