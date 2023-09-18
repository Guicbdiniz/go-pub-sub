package main

import (
	"log"

	"github.com/Guicbdiniz/go-pub-sub/internal/pubsub"
)

func main() {
	message := "Hello World!"

	// Create a new PubSub agent
	pubSub := pubsub.CreatePubSub()

	// Create a queue
	pubSub.CreateQueue("my_queue")

	// Subscribe to queue
	sub, err := pubSub.Subscribe("my_queue")
	if err != nil {
		log.Fatalf("err not nil when subscribing to topic, %v", err)
	}

	// Publish a message to the queue
	err = pubSub.Publish("my_queue", message)
	if err != nil {
		log.Fatalf("err not nil when publishing message, %v", err)
	}

	// Print received message from subscription
	receivedMessage := <-sub
	log.Printf("Received message: %v", receivedMessage)
}
