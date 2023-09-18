package pubsub

import (
	"testing"

	"github.com/Guicbdiniz/go-pub-sub/internal/utils"
)

func TestPubSubCreation(t *testing.T) {
	pubSub := CreatePubSub()
	utils.AssertEqual[int](t, 0, len(pubSub.queues), "PubSub should start with 0 queues created.")
}

func TestQueueCreation(t *testing.T) {
	pubSub := CreatePubSub()
	queue := pubSub.CreateQueue("first")
	utils.AssertEqual[string](t, "first", queue.GetName(), "Queue should have the initial name.")
	utils.AssertEqual[int](t, 1, len(pubSub.queues), "PubSub should update queues as one is created")
}

func TestGetQueue(t *testing.T) {
	pubSub := CreatePubSub()
	pubSub.CreateQueue("first")
	pubSub.CreateQueue("second")
	utils.AssertEqual[int](t, 2, len(pubSub.queues), "PubSub should update queues when created.")

	foundQueue, err := getQueue(pubSub.queues, "first")
	utils.AssertEqual[string](t, "first", foundQueue.GetName(), "Found queue should have same name as searched")
	utils.CheckTestError(t, err, "No error should be returned when queue is found")

	foundQueue, err = getQueue(pubSub.queues, "third")
	if foundQueue != nil {
		t.Fatalf("Found queue should be nil when passed name is not found. Expected nil, got %v", foundQueue)
	}
	if err == nil {
		t.Fatalf("err should not be nil when passed name is not found. Expected anything but nil, got %v", err)
	}
}

func TestPublish(t *testing.T) {
	initialMessage := "Hello World!"
	pubSub := CreatePubSub()
	pubSub.CreateQueue("first")

	err := pubSub.Publish("first", initialMessage)
	if err != nil {
		t.Fatalf("err should be nil when Publish is called with existing queue. Expected nil, got %v", err)
	}
	sub := pubSub.queues[0].Subscribe()
	receivedMessage := <-sub
	utils.AssertEqual(t, initialMessage, receivedMessage, "Received message should be equal published message")

	err = pubSub.Publish("second", initialMessage)
	if err == nil {
		t.Fatalf("err should not be nil when Publish is called with not existing queue. Expected anything but nil, got %v", err)
	}
}

func TestSubscribe(t *testing.T) {
	initialMessage := "Hello World!"
	pubSub := CreatePubSub()
	pubSub.CreateQueue("first")

	sub, err := pubSub.Subscribe("first")
	if err != nil {
		t.Fatalf("err should be nil when Subscribe is called with existing queue. Expected nil, got %v", err)
	}
	go pubSub.queues[0].Publish(initialMessage)
	receivedMessage := <-sub
	utils.AssertEqual(t, initialMessage, receivedMessage, "Received message should be equal published message")

	sub, err = pubSub.Subscribe("second")
	if sub != nil {
		t.Fatalf("sub should be nil when Subscribe is called with not existing queue. Expected nil, got %v", sub)
	}
	if err == nil {
		t.Fatalf("err should not be nil when Subscribe is called with not existing queue. Expected anything but nil, got %v", err)
	}
}

func TestAddLogger(t *testing.T) {
	pubSub := CreatePubSub()
	err := pubSub.AddLogger()
	if err != nil {
		t.Fatalf("err should be nil when adding logger to pub sub system. Expected nil, got %v", err)
	}
	utils.RemoveDataDirectory(t, pubSub.GetLoggerDirectoryPath())
}
