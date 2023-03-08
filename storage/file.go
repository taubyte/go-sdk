package storage

import (
	"fmt"
	"io"

	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
)

// GetFile grabs the the file from the storage
// Returns the file and an error
func (f *File) GetFile() (*StorageFile, error) {
	var fd uint32
	err := storageSym.StorageGetFile(uint32(f.storage), f.name, f.version, &fd)
	if err != 0 {
		return nil, fmt.Errorf("Getting file: `%s` failed with: %s", f.name, err)
	}

	return &StorageFile{storage: f.storage, fd: fd}, nil
}

// Read reads the given bytes
// Returns an int of how much was read and an error
func (file *StorageFile) Read(p []byte) (int, error) {
	var counter uint32
	err := storageSym.StorageReadFile(uint32(file.storage), file.fd, &p[0], uint32(len(p)), &counter)
	if err != 0 {
		if err == errno.ErrorEOF {
			return int(counter), io.EOF
		} else {
			return 0, fmt.Errorf("Reading file failed with: %s", err)
		}
	}

	return int(counter), nil
}

// Close closes the current file
// Returns an error
func (file *StorageFile) Close() error {
	err := storageSym.StorageCloseFile(uint32(file.storage), file.fd)
	if err != 0 {
		return fmt.Errorf("Closing file failed with: %s", err)
	}

	return nil
}
