package poebackend

import (
	"bz.moh.epi/poebackend/models"
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
)

// ArrivalCreated is the event triggered when an arrival is created in Firestore
func ArrivalCreated(ctx context.Context, e models.FirestoreArrivalEvent) error {
	log.Print("got request at ArrivalCreated")
	log.Printf("got event: %v", e)
	return nil
}

// CloudEventHandler handles cloud events
func CloudEventHandler(ctx context.Context, cloudEvents cloudevents.Event) error {
	log.Print("got request at ArrivalCreated")
	log.Printf("got event: %v", cloudEvents)
	return nil
}
