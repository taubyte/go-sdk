package storage

import (
	"testing"

	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
)

type deleteTest struct {
	storage uint32
	name    string
	version uint32
	all     uint32
}

func TestDelete(t *testing.T) {
	var deleted deleteTest
	db, err := New("testStorage")
	if err != nil {
		t.Error(err)
	}

	file := db.File("testFile")

	storageSym.StorageDeleteFile = func(storageId uint32, fileName string, version, all uint32) (error errno.Error) {
		deleted = deleteTest{
			storage: storageId,
			name:    fileName,
			version: version,
			all:     all,
		}

		return 0
	}

	err = file.Delete()
	if err != nil {
		t.Error(err)
	}

	if Storage(deleted.storage) != db || deleted.name != file.name {
		t.Errorf("Expected delete values are wrong %d != %d, %s != %s", deleted.storage, uint32(db), deleted.name, file.name)
	}
}
