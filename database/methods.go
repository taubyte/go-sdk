package database

import (
	"fmt"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
	"github.com/taubyte/go-sdk/utils/codec"
)

// Put writes the key and data into the database.
// Returns an error
func (d Database) Put(key string, data []byte) error {
	var dataPtr *byte
	if len(data) != 0 {
		dataPtr = &data[0]
	}

	err := databaseSym.DatabasePut(uint32(d), key, dataPtr, uint32(len(data)))
	if err != 0 {
		return fmt.Errorf("Database Put failed with: %s", err)
	}

	return nil
}

// Get retrieves the given key from the database.
// Returns the data of the key and an error.
func (d Database) Get(key string) ([]byte, error) {
	var size uint32
	err := databaseSym.DatabaseGetSize(uint32(d), key, &size)
	if err != 0 {
		return nil, fmt.Errorf("Database get size failed with: %s", err)
	}
	if size == 0 {
		return []byte{}, nil
	}

	data := make([]byte, size)
	err = databaseSym.DatabaseGet(uint32(d), key, &data[0])
	if err != 0 {
		return nil, fmt.Errorf("Database Get failed with: %s", err)
	}

	return data, nil
}

// Delete removes an entry from the database
// Retuns an error
func (d Database) Delete(key string) error {
	err := databaseSym.DatabaseDelete(uint32(d), key)
	if err != 0 {
		return fmt.Errorf("Deleting file: %s failed with: %s", key, err)
	}

	return nil
}

// Close closes the database
// Returns an error
func (d Database) Close() error {
	err := databaseSym.DatabaseClose(uint32(d))
	if err != 0 {
		return fmt.Errorf("Closing database: %d failed with %d", d, err)
	}

	return nil
}

// List uses the prefix given to get a list of keys that use the prefix in the database
// Returns all keys found and an error
func (d Database) List(prefix string) ([]string, error) {
	var size uint32
	err := databaseSym.DatabaseListSize(uint32(d), prefix, &size)
	if err != 0 {
		return nil, fmt.Errorf("Database get list size failed with: %s", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	data := make([]byte, size)
	err = databaseSym.DatabaseList(uint32(d), prefix, &data[0])
	if err != 0 {
		return nil, fmt.Errorf("Database List failed with: %s", err)
	}

	keys := make([]string, 0)
	err0 := codec.Convert(data).To(&keys)
	if err0 != nil || len(keys) == 0 {
		return nil, fmt.Errorf("Failed converting bytes to []string with: %s", err0)
	}

	return keys, nil
}
