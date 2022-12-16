package database

import (
	"testing"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
)

func TestDeleteError(t *testing.T) {
	testId := uint32(3)
	testKey := "someKey"
	databaseSym.MockDelete(testId, testKey, map[string][]byte{})

	err := Database(5).Delete(testKey)
	if err == nil {
		t.Error("expected Error")
		return
	}

	err = Database(testId).Delete("otherKey")
	if err == nil {
		t.Error("expected Error")
		return
	}
}
