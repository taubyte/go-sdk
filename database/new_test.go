package database

import (
	"testing"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
)

func TestNewError(t *testing.T) {
	testName := "someDatabase"
	databaseSym.MockNew(18, testName)

	_, err := New(testName + "other")
	if err == nil {
		t.Error("Expected an error")
		return
	}
}
