package storage

import (
	"testing"

	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
)

func TestNew(t *testing.T) {
	storageSym.StorageNew = func(storageName string, idPtr *uint32) (error errno.Error) {
		return errno.ErrorCap
	}

	_, err := New("someStorage")
	if err == nil {
		t.Error("Expected error")
		return
	}

	var newId uint32 = 5
	storageSym.StorageNew = func(storageName string, idPtr *uint32) (error errno.Error) {
		*idPtr = newId
		return 0
	}

	storage, err := New("someStorage")
	if err != nil {
		t.Error(err)
		return
	}
	if uint32(storage) != newId {
		t.Errorf("Got storage: `%d`, expected: `%d`", uint32(storage), newId)
		return
	}
}
