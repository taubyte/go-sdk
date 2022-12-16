package storage

import (
	"bytes"
	"io"
	"testing"
	"unsafe"

	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
)

func TestFileRead(t *testing.T) {
	sf := &StorageFile{}

	storageSym.StorageReadFile = func(storageId uint32, fd uint32, buf *byte, bufSize uint32, count *uint32) (error errno.Error) {
		return errno.ErrorCap
	}

	_, err := io.ReadAll(sf)
	if err == nil {
		t.Error("Expected error")
		return
	}

	storageSym.StorageCloseFile = func(storageId uint32, fd uint32) (error errno.Error) {
		return errno.ErrorCap
	}

	err = sf.Close()
	if err == nil {
		t.Error("Expected error")
		return
	}

	storageId := uint32(5)
	fileId := uint32(42)
	sf = &StorageFile{storage: Storage(storageId), fd: fileId}
	fileContent := []byte("Hello from the other side")
	b := bytes.NewBuffer([]byte(fileContent))
	storageSym.StorageReadFile = func(storageId uint32, fd uint32, buf *byte, bufSize uint32, countPtr *uint32) (error errno.Error) {
		if fd != sf.fd || Storage(storageId) != sf.storage {
			return errno.ErrorCap
		}

		data := unsafe.Slice(buf, bufSize)
		count, err := b.Read(data)
		if err != nil {
			if err == io.EOF {
				return errno.ErrorEOF
			}
			return errno.ErrorHttpReadBody
		}

		*countPtr = uint32(count)
		return 0
	}

	body, err := io.ReadAll(sf)
	if err != nil {
		t.Error(err)
		return
	}

	if string(body) != string(fileContent) {
		t.Errorf("Got content: `%s`, expected: `%s`", string(body), string(fileContent))
		return
	}

	storageSym.StorageCloseFile = func(storageId uint32, fd uint32) (error errno.Error) {
		if fd != sf.fd || Storage(storageId) != sf.storage {
			return errno.ErrorCap
		}

		return 0
	}

	err = sf.Close()
	if err != nil {
		t.Error(err)
		return
	}
}
