package filelogger

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateLoggerDir creates a directory next to the current executable file
// to keep the logs from each queue and returns its path.
//
// If the directory already exists, the function simply gets its path.
func CreateLoggerDir(dirName string) (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("error while getting current path to create logger, %w", err)
	}

	currentDir := filepath.Dir(executablePath)
	loggerDirPath := filepath.Join(currentDir, dirName)

	_, err = os.Stat(loggerDirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(loggerDirPath, 0755)
		if err != nil {
			return "", fmt.Errorf("error while creating logger directory, %w", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("error while getting status of current path, %w", err)
	}

	return loggerDirPath, nil
}

// CreateLoggerFile creates a file in the directory path passed
// to keep the logs from a queue.
//
// If the file already exists, the function simply gets its path.
func CreateLoggerFile(dirPath string, fileName string) (string, error) {
	filePath := filepath.Join(dirPath, fileName)
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filePath)
			if err != nil {
				return "", fmt.Errorf("error while creating logger file, %w", err)
			}
			defer file.Close()
		} else {
			return "", fmt.Errorf("error while getting status of logger file, %w", err)
		}
	}
	return filePath, nil
}
