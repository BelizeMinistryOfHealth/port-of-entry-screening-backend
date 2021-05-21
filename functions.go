package poebackend

import (
	"bz.moh.epi/poebackend/handlers"
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firesearch"
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

func PersonsHook(ctx context.Context, event models.FirestorePersonEvent) error {
	personStore := firesearch.PersonStore{Service: personFiresearchService}
	result, err := handlers.PersonCreated(ctx, event, personStore)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"result":  result,
		"context": ctx,
	}).Info("created person event")
	return nil
}

func PersonDeletedListener(ctx context.Context, event models.FirestorePersonEvent) error {
	personStore := firesearch.PersonStore{
		Service: personFiresearchService,
	}
	personID := event.OldValue.Fields.ID.StringValue
	if err := handlers.PersonDeleted(ctx, personStore, personID); err != nil {
		log.WithFields(log.Fields{
			"event":    event,
			"personID": personID,
		}).Error(err)

		return fmt.Errorf("PersonDeletedListener failed: %w", err)
	}
	return nil
}

// GetServer exposes Server to modify some settings
func GetServer() *Server {
	return &server
}
