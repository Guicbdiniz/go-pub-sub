package filelogger

import (
	"os"
	"testing"

	"github.com/Guicbdiniz/go-pub-sub/internal/utils"
)

func TestCreateLoggerDir(t *testing.T) {
	loggerDirPath, err := CreateLoggerDir("test")
	if err != nil {
		t.Fatalf("err should be nil when logger directory creation is called. Expected nil, got %v", err)
	}

	_, err = os.Stat(loggerDirPath)
	if err != nil {
		t.Fatalf("err should be nil when checking the logger directory. Expected nil, got %v", err)
	}

	_, err = CreateLoggerDir("test")
	if err != nil {
		t.Fatalf("err should be nil when logger directory creation is called even if the directory exists. Expected nil, got %v", err)
	}

	utils.RemoveDataDirectory(t, loggerDirPath)
}

func TestCreateLoggerFile(t *testing.T) {
	loggerDirPath, err := CreateLoggerDir("test")
	if err != nil {
		t.Fatalf("err should be nil when logger directory creation is called. Expected nil, got %v", err)
	}

	loggerFilePath, err := CreateLoggerFile(loggerDirPath, "first")
	if err != nil {
		t.Fatalf("err should be nil when logger file creation is called. Expected nil, got %v", err)
	}

	_, err = os.Stat(loggerFilePath)
	if err != nil {
		t.Fatalf("err should be nil when checking the logger file. Expected nil, got %v", err)
	}

	_, err = CreateLoggerFile(loggerDirPath, "first")
	if err != nil {
		t.Fatalf("err should be nil when logger file creation is called even if the file exists. Expected nil, got %v", err)
	}

	utils.RemoveDataDirectory(t, loggerDirPath)
}
