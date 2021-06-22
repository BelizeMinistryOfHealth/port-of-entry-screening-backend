package poebackend

import (
	"bz.moh.epi/poebackend/handlers"
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firesearch"
	"bz.moh.epi/poebackend/repository/firestore"
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event/datacodec/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var server Server                              //nolint:gochecknoglobals
var personFiresearchService firesearch.Service //nolint:gochecknoglobals

func init() {
	backendBaseURL := "https://us-east1-epi-belize.cloudfunctions.net"
	server = Server{
		BackendBaseURL: backendBaseURL,
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	personFiresearchService = firesearch.CreateFiresearchService("Persons Index", "persons_index", "NA")
}

func enableCors() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Referer, Connection")
			w.Header().Set("responseType", "*")
			f(w, r)
		}
	}
}

// HandlerEcho is an echo endpoint for testing purposes
func HandlerEcho(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		panic("simple hello echo failed")
	}
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub is a sample cloud function that subscribes to PubSub
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	//name := string(m.Data) // Automatically decoded from base64.
	//if name == "" {
	//	name = "World"
	//}
	var person models.Person
	if err := json.Decode(ctx, m.Data, &person); err != nil {
		log.Printf("Error decoding person: %v", err)
		return nil
	}

	log.Printf("Hello, %v!", person)
	return nil
}

// PersonsHook is triggered when a new record is inserted in the persons collection
func PersonsHook(ctx context.Context, event models.FirestorePersonEvent) error {
	personStore := firesearch.PersonStore{Service: personFiresearchService}
	result, err := handlers.PersonCreated(ctx, event, personStore)
	if err != nil {
		log.WithError(err).Info("creating user index failed")
		return fmt.Errorf("new person liatener failed: %w", err)
	}
	log.WithFields(log.Fields{
		"result":  result,
		"context": ctx,
	}).Info("created person event")
	return nil
}

// PersonDeletedListener is the function triggered when a person is deleted
func PersonDeletedListener(ctx context.Context, event models.FirestorePersonEvent) error {
	personStore := firesearch.PersonStore{
		Service: personFiresearchService,
	}
	personID := event.OldValue.Fields.ID.StringValue
	projectID := os.Getenv("PROJECT_ID")
	firestoreDb, err := firestore.CreateFirestoreDB(ctx, projectID)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "AccessKeyFn",
			"message": "error creating firestore db connection",
		}).WithError(err)
		return fmt.Errorf("PersonDeletedListener failed: %w", err)
	}
	addressStoreService := firestore.CreateAddressStoreService(firestoreDb, "addresses")
	arrivalStoreService := firestore.CreateArrivalsStoreService(firestoreDb, "arrivals")
	args := handlers.PersonDeletedArgs{
		PersonFiresearchStore: personStore,
		ArrivalStoreService:   arrivalStoreService,
		AddressStoreService:   addressStoreService,
	}
	if err := handlers.PersonDeleted(ctx, args, personID); err != nil {
		log.WithFields(log.Fields{
			"event":    event,
			"personID": personID,
		}).Error(err)

		return fmt.Errorf("PersonDeletedListener failed: %w", err)
	}
	return nil
}

// PersonUpdatedListener is triggered when a record is updated in the persons collection
func PersonUpdatedListener(ctx context.Context, event models.FirestorePersonEvent) error {
	personStore := firesearch.PersonStore{Service: personFiresearchService}
	log.WithFields(log.Fields{
		"event": event,
	}).Info("updating person record")
	result, err := handlers.PersonCreated(ctx, event, personStore)
	if err != nil {
		log.WithError(err).Info("update person failed")
		return fmt.Errorf("update Person Failed: %w", err)
	}
	log.WithFields(log.Fields{
		"result":  result,
		"context": ctx,
	}).Info("update person event")
	return nil
}

// ScreeningListener Cloud Function triggered when a screening is created or updated
func ScreeningListener(ctx context.Context, event models.FirestoreScreeningEvent) error {
	log.WithFields(log.Fields{
		"event": event,
	}).Info("screening event")
	projectID := os.Getenv("PROJECT_ID")
	firestoreDB, err := firestore.CreateFirestoreDB(ctx, projectID)
	if err != nil {
		log.WithFields(log.Fields{
			"event": event,
		}).WithError(err).Info("failed to create firestore connection")
		return fmt.Errorf("failed to create firestore connection: %w", err)
	}

	personStore := firestore.CreatePersonService(firestoreDB, "persons")
	arrivalStore := firestore.CreateArrivalsStoreService(firestoreDB, "arrivals")
	addressStore := firestore.CreateAddressStoreService(firestoreDB, "addresses")
	godataErr := handlers.ScreeningEventHandler(ctx, event, personStore, arrivalStore, addressStore)
	if godataErr != nil {
		return fmt.Errorf("godata persistence failed: %w", godataErr)
	}

	return nil
}

// AccessKeyFn is the function that returns a firesearch access key
func AccessKeyFn(w http.ResponseWriter, r *http.Request) {
	Chainz(accessKeyHandler, enableCors())(w, r)
}

func accessKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		projectID := os.Getenv("PROJECT_ID")
		firestoreDb, err := firestore.CreateFirestoreDB(r.Context(), projectID)
		if err != nil {
			log.WithFields(log.Fields{
				"handler": "AccessKeyFn",
				"message": "error creating firestore db connection",
			}).WithError(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		handlers.AccessKeyHandler(*firestoreDb, w, r)
	}
}

// ArrivalsStatAccessKeyFn is the function that returns a firesearch access key
func ArrivalsStatAccessKeyFn(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"headers": r.Header,
	}).Info("requested access key for arrivals index")
	Chainz(arrivalsStatAccessKeyHandler, enableCors())(w, r)
}

func arrivalsStatAccessKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		projectID := os.Getenv("PROJECT_ID")
		firestoreDb, err := firestore.CreateFirestoreDB(r.Context(), projectID)
		if err != nil {
			log.WithFields(log.Fields{
				"handler": "ArrivalsStatAccessKeyFn",
				"message": "error creating firestore db connection",
			}).WithError(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		handlers.ArrivalsStatAccessKeyHandler(*firestoreDb, w, r)
	}
}

// RegistrationFn is the REST endpoint for registering a traveller
func RegistrationFn(w http.ResponseWriter, r *http.Request) {
	Chainz(registrationFn, enableCors())(w, r)
}

func registrationFn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		projectID := os.Getenv("PROJECT_ID")
		firestoreDb, err := firestore.CreateFirestoreDB(r.Context(), projectID)
		if err != nil {
			log.WithFields(log.Fields{
				"handler": "AccessKeyFn",
				"message": "error creating firestore db connection",
			}).WithError(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		personStoreService := firestore.CreatePersonService(firestoreDb, "persons")
		addressStoreService := firestore.CreateAddressStoreService(firestoreDb, "addresses")
		arrivalStoreService := firestore.CreateArrivalsStoreService(firestoreDb, "arrivals")
		args := handlers.RegistrationArgs{
			PersonStoreService:  personStoreService,
			ArrivalStoreService: arrivalStoreService,
			AddressStoreService: addressStoreService,
		}
		handlers.RegistrationHandler(args, w, r)
	}
}

// ArrivalsListener subscribes to events on the arrivals collection
func ArrivalsListener(ctx context.Context, event models.FirestoreArrivalEvent) error {
	firesearchService := firesearch.CreateFiresearchService("Arrivals Stat Index", "arrivals_stat_index", "NA")
	arrivalsStore := firesearch.ArrivalsStore{Service: firesearchService}
	res, err := handlers.ArrivalStatEvent(ctx, event, arrivalsStore)
	if err != nil {
		return fmt.Errorf("failure in the arrivals listener: %w", err)
	}
	log.WithFields(log.Fields{
		"event": event,
		"stat":  res.ArrivalStat,
	}).Info("successfully processed an arrival event")
	return nil
}

func ArrivalsDeletedListener(ctx context.Context, event models.FirestoreArrivalEvent) error {
	firesearchService := firesearch.CreateFiresearchService("Arrivals Stat Index", "arrivals_stat_index", "NA")
	arrivalsStore := firesearch.ArrivalsStore{Service: firesearchService}
	ID := event.OldValue.Fields.ID.StringValue
	if err := handlers.ArrivalDeleted(ctx, arrivalsStore, ID); err != nil {
		return fmt.Errorf("failed to delete arrival: %w", err)
	}
	return nil
}

// GetServer exposes Server to modify some settings
func GetServer() *Server {
	return &server
}
