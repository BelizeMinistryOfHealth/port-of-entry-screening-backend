package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Example of how to use pubsub.
func main() {
	// Create a new subscription.
	proj := "epi-belize"
	topic := "TEST_TOPIC"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		panic(fmt.Sprintf("pubsub.NewClient failed: %v", err))
	}
	//pubsubTopic := poebackend.CreateTopic(ctx, client, topic)
	pubsubTopic := client.Topic(topic)
	subscriber := "test_subscriber"
	if err := create(ctx, client, subscriber, pubsubTopic); err != nil {
		log.Fatal(err)
	}
	// Pull messages via the subscription.
	if err := pullMsgs(ctx, client, subscriber, pubsubTopic, true); err != nil {
		log.Fatal(err)
	}
}

func create(ctx context.Context, client *pubsub.Client, name string, topic *pubsub.Topic) error {
	// [START pubsub_create_pull_subscription]
	sub, err := client.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("create error: %w", err)
	}
	fmt.Printf("Created subscription: %v\n", sub)
	// [END pubsub_create_pull_subscription]
	return nil
}

func pullMsgs(ctx context.Context, client *pubsub.Client, name string, topic *pubsub.Topic, testPublish bool) error {

	if testPublish {
		// Publish 10 messages on the topic.
		var results []*pubsub.PublishResult
		for i := 0; i < 10; i++ {
			res := topic.Publish(ctx, &pubsub.Message{
				Data: []byte(fmt.Sprintf("hello world #%d", i)),
			})
			results = append(results, res)
		}

		// Check that all messages were published.
		for _, r := range results {
			_, err := r.Get(ctx)
			if err != nil {
				return fmt.Errorf("pullMsgs error: %w", err)
			}
		}
	}
	var mu sync.Mutex
	received := 0
	sub := client.Subscription(name)
	cctx, cancel := context.WithCancel(ctx)
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		mu.Lock()
		defer mu.Unlock()
		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		return fmt.Errorf("pullMsgs error: %w", err)
	}
	return nil
}
