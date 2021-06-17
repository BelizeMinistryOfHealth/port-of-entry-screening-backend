package integration

import (
	"bz.moh.epi/poebackend"
	"bz.moh.epi/poebackend/models"
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/cloudevents/sdk-go/v2/event/datacodec/json"
	"testing"
	"time"
)

func TestCreateTopic(t *testing.T) {
	ctx := context.Background()
	topic := "TEST_TOPIC"
	proj := "epi-belize"
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		t.Fatalf("pubsub.NewClient failed: %v", err)
	}
	pubsubTopic := poebackend.CreateTopic(ctx, client, topic)
	if pubsubTopic == nil {
		t.Fatalf("CreateTopic failed: %v", err)
	}
	t.Logf("topic: %v", pubsubTopic)
}

func TestPublishMessage(t *testing.T) {
	ctx := context.Background()
	topic := "arrival_created"
	proj := "epi-belize-staging"
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		t.Fatalf("pubsub.NewClient failed: %v", err)
	}
	person := models.Person{
		ID: "11111",
		PersonalInfo: models.PersonalInfo{
			FirstName: "Raheem",
			LastName:  "Fernandez",
			FullName:  "Raheem Fernandez",
		},
		PortOfEntry: "PGIA",
		Arrival: models.Arrival{
			ArrivalInfo: models.ArrivalInfo{
				DateOfArrival: time.Now(),
			},
		},
	}

	encodedMsg, _ := json.Encode(ctx, person)

	msg := &pubsub.Message{
		Data: encodedMsg,
	}
	pubresult := client.Topic(topic).Publish(ctx, msg)
	_, err = pubresult.Get(ctx)
	if err != nil {
		t.Fatalf("Publishing error: %v", err)
	}
}
