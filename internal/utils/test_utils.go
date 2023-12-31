package utils

import (
	"os"
	"testing"
)

// CheckTestError is a test helper to confirm that an error is nil.
//
// If error is not nil, test will fail.
func CheckTestError(t *testing.T, err error, message string) {
	t.Helper()

	if err != nil {
		t.Fatalf("%s: %s", message, err.Error())
	}
}

// AssertEqual compares two values, failing if they are not equal.
func AssertEqual[V comparable](t *testing.T, expected V, actual V, errorMessage string) {
	t.Helper()

	if expected != actual {
		t.Fatalf("%s. Expected %v, got %v", errorMessage, expected, actual)
	}
}

// RemoveDataDirectory removes directory passed.
//
// It should be used when a data directory is created in the test.
func RemoveDataDirectory(t *testing.T, dirPath string) {
	t.Helper()

	err := os.RemoveAll(dirPath)
	if err != nil {
		t.Fatalf("err should be nil when deleting directory. Expected nil, got %v", err)
	}
}
