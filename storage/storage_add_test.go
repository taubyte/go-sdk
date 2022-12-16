package storage

import (
	"testing"

	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
)

type addTest struct {
	storage  uint32
	filename string
}

func TestAdd(t *testing.T) {
	var add addTest
	var testData = "testingData Here"

	storageSym.StorageAddFile = func(storageId uint32, fileName string, versionPtr *uint32, bufPtr *byte, bufLen, overWrite uint32) (error errno.Error) {
		add = addTest{
			storage:  storageId,
			filename: fileName,
		}
		return 0
	}

	db, err := New("testStorage")
	if err != nil {
		t.Error(err)
		return
	}

	_file := db.File("testFile")
	_, err = _file.Add([]byte(testData), false)
	if err != nil {
		t.Error(err)
		return
	}

	if Storage(add.storage) != db || add.filename != _file.name {
		t.Errorf("Expected values are not equal %d != %d, %s != %s", add.storage, uint32(db), _file.name, add.filename)
	}
}

func TestAddEmpty(t *testing.T) {
	db, err := New("testStorage")
	if err != nil {
		t.Error(err)
		return
	}

	_file := db.File("testFile")
	_, err = _file.Add([]byte{}, false)
	if err == nil {
		t.Error("This should of failed passing in empty data")
		return
	}
}
