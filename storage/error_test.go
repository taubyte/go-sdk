package storage_test

import (
	"testing"
	"unsafe"

	symbols "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/storage"
)

func TestStorageErrors(t *testing.T) {
	symbols.StorageMockData{StorageId: 2}.Mock()

	symbols.StorageGet = func(storageName string, idPtr *uint32) (error errno.Error) {
		return 1
	}

	_, err := storage.Get("fail")
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageDeleteFile = func(storageId uint32, fileName string, version, all uint32) (error errno.Error) {
		return 1
	}

	symbols.StorageListFilesSize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
		return 1
	}

	var testStorage storage.Storage
	_, err = testStorage.ListFiles()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageListFilesSize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
		*sizePtr = uint32(12)
		return 0
	}

	symbols.StorageListFiles = func(storageId uint32, bufPtr *byte) (error errno.Error) {
		return 1
	}
	_, err = testStorage.ListFiles()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageCidSize = func(storageId uint32, fileName string, idPtr *uint32) (error errno.Error) {
		return 1
	}

	_, err = testStorage.Cid("failFile")
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageCidSize = func(storageId uint32, fileName string, idPtr *uint32) (error errno.Error) {
		return 0
	}

	symbols.StorageCid = func(cidPtr *byte, idPtr *uint32) (error errno.Error) {
		return 1
	}

	_, err = testStorage.Cid("failFile")
	if err == nil {
		t.Error("Expected error")
		return
	}

	var file storage.File
	err = file.Delete()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageGetFile = func(storageId uint32, fileName string, version uint32, fdPtr *uint32) (error errno.Error) {
		return 1
	}

	_, err = file.GetFile()
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = file.DeleteAllVersions()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageAddFile = func(storageId uint32, fileName string, versionPtr *uint32, bufPtr *byte, bufLen, overWrite uint32) (error errno.Error) {
		return 1
	}

	_, err = file.Add([]byte("test"), false)
	if err == nil {
		t.Error("Expected error")
		return
	}

}
func TestVersionsError(t *testing.T) {
	var file storage.File

	symbols.StorageListVersionsSize = func(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}

	symbols.StorageListVersions = func(storageId uint32, fileName string, versionPtr *byte) (error errno.Error) {
		d := unsafe.Slice(versionPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}

	_, err := file.Versions()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageListVersions = func(storageId uint32, fileName string, versionPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = file.Versions()
	if err == nil {
		t.Error("Expected error")
		return
	}

}

func TestCurrentVersion(t *testing.T) {
	var file storage.File

	symbols.StorageCurrentVersionSize = func(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}

	symbols.StorageCurrentVersion = func(filename string, versionPtr *byte) (error errno.Error) {
		d := unsafe.Slice(versionPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}

	_, err := file.CurrentVersion()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageCurrentVersion = func(filename string, versionPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = file.CurrentVersion()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageCurrentVersionSize = func(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 0
		return 0
	}

	_, err = file.CurrentVersion()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageCurrentVersionSize = func(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error) {
		return 1
	}

	_, err = file.CurrentVersion()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestListFiles(t *testing.T) {
	symbols.StorageMockData{
		Files: "file/testFile1/abc/file/testFile2/2",
	}.Mock()

	var storage storage.Storage

	_, err := storage.ListFiles()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestRemainingCapacity(t *testing.T) {
	var testStorage storage.Storage
	symbols.StorageUsedSize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 10
		return 0
	}

	symbols.StorageUsed = func(storageId uint32, usedPtr *byte) (error errno.Error) {
		d := unsafe.Slice(usedPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}

	symbols.StorageCapacity = func(storageId uint32, capacityPtr *byte) (error errno.Error) {
		d := unsafe.Slice(capacityPtr, 10)
		copy(d, []byte("10"))
		return 0
	}

	_, err := testStorage.RemainingCapacity()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageUsed = func(storageId uint32, usedPtr *byte) (error errno.Error) {
		return 0
	}

	symbols.StorageCapacitySize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 10
		return 0
	}

	symbols.StorageCapacity = func(storageId uint32, capacityPtr *byte) (error errno.Error) {
		d := unsafe.Slice(capacityPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}

	_, err = testStorage.RemainingCapacity()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageCapacity = func(storageId uint32, capacityPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = testStorage.RemainingCapacity()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageCapacitySize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
		return 1
	}

	_, err = testStorage.RemainingCapacity()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageUsed = func(storageId uint32, usedPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = testStorage.RemainingCapacity()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageUsedSize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
		return 1
	}

	_, err = testStorage.RemainingCapacity()
	if err == nil {
		t.Error("Expected error")
		return
	}

}
