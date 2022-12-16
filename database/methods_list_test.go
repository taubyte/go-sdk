package database

import (
	"testing"
	"unsafe"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
	"github.com/taubyte/go-sdk/errno"
)

func TestListEmpty(t *testing.T) {
	testId := uint32(5)
	testKey := "someKey"
	testData := map[string][]byte{}
	databaseSym.MockList(testId, testKey, testData)

	arr, err := Database(testId).List(testKey)
	if err != nil {
		t.Error(err)
		return
	}
	if len(arr) != 0 {
		t.Errorf("Got %d expected 0", len(arr))
		return
	}
}

func TestListError(t *testing.T) {
	testId := uint32(5)
	testKey := "someKey"
	testData := map[string][]byte{
		testKey + "one": {},
		testKey + "two": {},
		"somekeyone":    {},
	}
	databaseSym.MockList(testId, testKey, testData)

	_, err := Database(0).List(testKey)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = Database(testId).List("otherKey")
	if err == nil {
		t.Error("Expected error")
		return
	}

	databaseSym.DatabaseList = func(databaseId uint32, key string, data *byte) (error errno.Error) {
		return 1
	}

	_, err = Database(testId).List(testKey)
	if err == nil {
		t.Error("Expected error")
		return
	}

	databaseSym.DatabaseListSize = func(databaseId uint32, key string, size *uint32) (error errno.Error) {
		*size = 12
		return 0
	}

	databaseSym.DatabaseList = func(databaseId uint32, key string, data *byte) (error errno.Error) {
		d := unsafe.Slice(data, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}

	_, err = Database(testId).List(testKey + "two")
	if err == nil {
		t.Error("Expected error")
		return
	}
}
