package database

import (
	"testing"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
	"github.com/taubyte/go-sdk/errno"
)

func TestGetEmpty(t *testing.T) {
	testId := uint32(3)
	testKey := "testKey"
	testData := map[string][]byte{
		testKey: {},
	}
	databaseSym.MockGet(testId, testData)

	_, err := Database(testId).Get(testKey)
	if err != nil {
		t.Error("expected Error")
		return
	}
}

func TestGetError(t *testing.T) {
	testId := uint32(3)
	testKey := "testKey"
	testData := map[string][]byte{
		testKey: []byte("Hello, world"),
	}
	databaseSym.MockGet(testId, testData)

	_, err := Database(testId).Get("otherKey")
	if err == nil {
		t.Error("expected Error")
		return
	}

	_, err = Database(5).Get(testKey)
	if err == nil {
		t.Error("expected Error")
		return
	}

	databaseSym.DatabaseGet = func(databaseId uint32, key string, data *byte) (error errno.Error) {
		return 1
	}
	_, err = Database(testId).Get(testKey)
	if err == nil {
		t.Error("expected Error")
		return
	}
}
