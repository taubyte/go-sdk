package storage

import (
	"testing"

	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
)

type testStorage struct {
	storageName string
	filename    string
	storage     uint32
	version     uint32
}

func TestGet(t *testing.T) {
	var testGet testStorage
	var expectedName = "pass"
	storageSym.StorageGet = func(storageName string, idPtr *uint32) (error errno.Error) {
		testGet.storageName = expectedName

		return 0
	}

	_, err := Get("someStorage")
	if err != nil {
		t.Error(err)
		return
	}

	if testGet.storageName != "pass" {
		t.Errorf("Expected name did not match %s != %s", testGet.storageName, expectedName)
	}
}

func TestFileGet(t *testing.T) {
	var testFile testStorage
	storage, err := New("testStorage")
	if err != nil {
		t.Error(err)
		return
	}

	storageSym.StorageGetFile = func(storageId uint32, fileName string, version uint32, fdPtr *uint32) (error errno.Error) {
		testFile = testStorage{
			filename: fileName,
			storage:  storageId,
			version:  version,
		}

		return 0
	}

	_file := storage.File("testFile")

	_, err = _file.GetFile()
	if err != nil {
		t.Error(err)
		return
	}

	if Storage(testFile.storage) != storage || testFile.version != _file.version || testFile.filename != _file.name {
		t.Errorf("Expected values are not met. %d != %d, %d != %d, %s != %s", testFile.storage, uint32(storage), testFile.version, _file.version, testFile.filename, _file.name)
	}

}
