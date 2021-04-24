package poebackend

import (
	"cloud.google.com/go/pubsub"
	"context"
	"testing"
)

func TestCreateTopic(t *testing.T) {
	ctx := context.Background()
	topic := "TEST_TOPIC"
	proj := "epi-belize"
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		t.Fatalf("pubsub.NewClient failed: %v", err)
	}
	pubsubTopic := CreateTopic(ctx, client, topic)
	if pubsubTopic == nil {
		t.Fatalf("CreateTopic failed: %v", err)
	}
	t.Logf("topic: %v", pubsubTopic)
}
