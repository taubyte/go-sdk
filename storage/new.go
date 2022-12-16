package storage

import (
	"fmt"

	storageSym "github.com/taubyte/go-sdk-symbols/storage"
)

// New creates a new storage
// Returns a storage and an error
func New(storageName string) (Storage, error) {
	var id uint32
	err := storageSym.StorageNew(storageName, &id)
	if err != 0 {
		return 0, fmt.Errorf("Failed to create storage: `%s` with: %s, is it configured?", storageName, err)
	}

	return Storage(id), nil
}
