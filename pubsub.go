package poebackend

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

// CreateTopic creates a PubSub Topic
func CreateTopic(ctx context.Context, c *pubsub.Client, topic string) *pubsub.Topic {
	t := c.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return t
	}
	t, err = c.CreateTopic(ctx, topic)
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}

	return t
}
