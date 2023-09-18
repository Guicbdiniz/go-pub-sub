package queue

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Guicbdiniz/go-pub-sub/internal/utils/filelogger"
)

type Queue struct {
	name        string
	subs        []chan string
	logFilePath string
}

// CreateQueue creates a Pub/Sub queue with a name and no subscribers.
func CreateQueue(name string) *Queue {
	return &Queue{name: name, subs: []chan string{}, logFilePath: ""}
}

// AddLogger creates a logger file to log every message sent to the queue.
func (q *Queue) AddLogger(logDirPath string) error {
	filePath, err := filelogger.CreateLoggerFile(logDirPath, q.name)
	if err != nil {
		return fmt.Errorf("error while adding logger to queue, %w", err)
	}
	q.logFilePath = filePath
	return nil
}

// GetName gets the queue name.
func (q *Queue) GetName() string {
	return q.name
}

// Publish sends the passed data to all subscribers and logs the data if logger is available.
func (q *Queue) Publish(data string) {
	if q.logFilePath != "" {
		err := q.logData(data)
		if err != nil {
			log.Printf("Error while logging data: %v", err)
		}
	}
	for _, ch := range q.subs {
		go (func() {
			ch <- data
		})()
	}
}

// Subscribe adds a new subscriber which will receive all published data from the queue.
func (q *Queue) Subscribe() <-chan string {
	ch := make(chan string)
	q.subs = append(q.subs, ch)
	return ch
}

func (q *Queue) logData(data string) error {
	file, err := os.OpenFile(q.logFilePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("error while opening log file from queue, %w", err)
	}
	defer file.Close()

	stringLine := fmt.Sprintf("%s %s\n", time.Now().Format("1000-01-01 00:00:00"), data)

	_, err = file.WriteString(stringLine)
	if err != nil {
		return fmt.Errorf("error while writing log to file, %w", err)
	}

	return nil
}
