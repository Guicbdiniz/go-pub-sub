package pubsub

import (
	"errors"
	"fmt"
	"log"

	"github.com/Guicbdiniz/go-pub-sub/internal/utils/filelogger"

	"github.com/Guicbdiniz/go-pub-sub/internal/queue"
)

type PubSub struct {
	queues        []*queue.Queue
	loggerDirPath string
}

// CreatePubSub creates a new Publish/Subscribe system with no queues and no logger attached.
func CreatePubSub() *PubSub {
	pubSub := PubSub{
		queues:        []*queue.Queue{},
		loggerDirPath: "",
	}
	return &pubSub
}

// AddLogger attaches a directory to the Pub/Sub system, creating it if necessary.
func (pubSub *PubSub) AddLogger() error {
	loggerDirName := "data"
	loggerDirPath, err := filelogger.CreateLoggerDir(loggerDirName)
	if err != nil {
		return fmt.Errorf("error captured while creating logger directory, %w", err)
	}

	pubSub.loggerDirPath = loggerDirPath
	return nil
}

// GetLoggerDirectoryPath gets the logger directory attached to the Pub/Sub system.
func (pubSub *PubSub) GetLoggerDirectoryPath() string {
	return pubSub.loggerDirPath
}

// CreateQueue adds new queue to the Pub/Sub system with the passed name.
func (pubSub *PubSub) CreateQueue(name string) *queue.Queue {
	queue := queue.CreateQueue(name)
	if pubSub.loggerDirPath != "" {
		err := queue.AddLogger(pubSub.loggerDirPath)
		if err != nil {
			log.Printf("Error while adding logger to queue: %v", err)
		}
	}
	pubSub.queues = append(pubSub.queues, queue)
	return queue
}

// Publish sends data to all the subscribers from the queue with the passed name (if it exists).
func (pubSub *PubSub) Publish(queueName string, data string) error {
	q, err := getQueue(pubSub.queues, queueName)
	if err != nil {
		return err
	}
	go q.Publish(data)
	return nil
}

// Subscribe adds a new subscriber to the queue with the passed topic.
func (pubSub *PubSub) Subscribe(topic string) (<-chan string, error) {
	q, err := getQueue(pubSub.queues, topic)
	if err != nil {
		return nil, err
	}
	return q.Subscribe(), nil
}

func getQueue(queues []*queue.Queue, name string) (*queue.Queue, error) {
	for _, q := range queues {
		if q.GetName() == name {
			return q, nil
		}
	}
	return nil, errors.New("no queue with specified name found")
}
