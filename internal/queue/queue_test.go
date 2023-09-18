package queue

import (
	"path/filepath"
	"testing"

	"github.com/Guicbdiniz/go-pub-sub/internal/utils/filelogger"

	"github.com/Guicbdiniz/go-pub-sub/internal/utils"
)

func TestCreateQueue(t *testing.T) {
	queue := CreateQueue("test")
	utils.AssertEqual[string](t, "test", queue.name, "Queue should keep initial name in creation")
	utils.AssertEqual[int](t, 0, len(queue.subs), "Queue should have no subs in creation")
}

func TestPublishSubscribe(t *testing.T) {
	message := "Hello World!"
	queue := CreateQueue("test")
	sub := queue.Subscribe()

	go queue.Publish(message)

	received := <-sub

	utils.AssertEqual[string](t, message, received, "Received message should be the same as sent")
}

func TestAddLogger(t *testing.T) {
	queueName := "test_queue"
	loggerDirPath, err := filelogger.CreateLoggerDir("test")
	defer utils.RemoveDataDirectory(t, loggerDirPath)
	if err != nil {
		t.Fatalf("err should be nil when creating logger directory. Expected nil, got %v", err)
	}
	queue := CreateQueue(queueName)
	err = queue.AddLogger(loggerDirPath)
	if err != nil {
		t.Fatalf("err should be nil when adding logger to file. Expected nil, got %v", err)
	}

	utils.AssertEqual[string](t, filepath.Join(loggerDirPath, queue.name), queue.logFilePath, "Log file path from queue should be created correctly")
}
