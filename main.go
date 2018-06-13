package main

import (
	"context"
	"errors"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

func main() {
	projectID := os.Getenv("PROJECT_ID")
	subscriptionName := os.Getenv("SUBSCRIPTION_NAME")

	ctx := context.Background()

	log.Printf("Creating pubsub client connection for project: %s", projectID)
	client, err := pubsub.NewClient(ctx, projectID)
	fatalOnErr(err)
	log.Println("Client created successfully")

	log.Printf("Getting subscription for: %s", subscriptionName)
	sub := client.Subscription(subscriptionName)
	exists, err := sub.Exists(ctx)
	fatalOnErr(err)
	if !exists {
		fatalOnErr(errors.New("Subscription doesn't exist"))
	}

	sub.ReceiveSettings.MaxOutstandingMessages = 1
	err = sub.Receive(ctx, func(mCtx context.Context, message *pubsub.Message) {
		message.Ack()
		messageData := message.Data
		log.Println(string(messageData))
	})
}

func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
