package storage

import (
	"testing"

	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
)

type testList struct {
}

func TestList(t *testing.T) {
	storageSym.StorageListVersionsSize = func(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error) {
		return 0
	}

	storageSym.StorageListVersions = func(storageId uint32, fileName string, versionPtr *byte) (error errno.Error) {
		return 0
	}

	db, err := New("testdb")
	if err != nil {
		t.Error(err)
		return
	}

	file := db.File("testFile")
	_, err = file.Versions()
	if err == nil {
		t.Error("Expected error listing empty versions")
		return
	}
}
