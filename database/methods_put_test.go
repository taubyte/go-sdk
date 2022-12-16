package database

import (
	"testing"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
)

func TestPutEmpty(t *testing.T) {
	testId := uint32(6)
	putData := map[string][]byte{}
	databaseSym.MockPut(testId, putData)

	testKey := "someKey"
	err := Database(testId).Put(testKey, []byte{})
	if err != nil {
		t.Error(err)
		return
	}

	_, ok := putData["otherKey"]
	if ok == true {
		t.Error("Test logic failed, otherKey should be nil")
	}

	_, ok = putData[testKey]
	if ok == false {
		t.Errorf("Got value `%s` expected nil", putData[testKey])
		return
	}
}

func TestPutError(t *testing.T) {
	putData := map[string][]byte{}
	databaseSym.MockPut(6, putData)

	testKey := "someKey"
	err := Database(8).Put(testKey, []byte{})
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, ok := putData[testKey]
	if ok == true {
		t.Errorf("Got value `%s` expected nil", putData[testKey])
		return
	}
}
