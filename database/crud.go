package database

import (
	"encoding/binary"
	"errors"
	"fmt"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
	"github.com/taubyte/go-sdk/utils/sized"
)

// Entry is a key with its value.
type Entry struct {
	Key   string
	Value []byte
}

// Batch queues puts and deletes on the host; nothing is written until Commit.
type Batch uint32

// Has reports whether key exists.
func (d Database) Has(key string) (bool, error) {
	_, err := d.Get(key)
	if errors.Is(err, ErrNotFound) {
		return false, nil
	}
	return err == nil, err
}

// Batch starts a batch of writes on the database.
func (d Database) Batch() (Batch, error) {
	var id uint32
	if err := databaseSym.DatabaseBatchNew(uint32(d), &id); err != 0 {
		return 0, fmt.Errorf("Database Batch failed with: %s", err)
	}
	return Batch(id), nil
}

// Put queues a write of data at key.
func (b Batch) Put(key string, data []byte) error {
	var p *byte
	if len(data) != 0 {
		p = &data[0]
	}
	if err := databaseSym.DatabaseBatchPut(uint32(b), key, p, uint32(len(data))); err != 0 {
		return fmt.Errorf("Batch Put failed with: %s", err)
	}
	return nil
}

// Delete queues the removal of key.
func (b Batch) Delete(key string) error {
	if err := databaseSym.DatabaseBatchDelete(uint32(b), key); err != 0 {
		return fmt.Errorf("Batch Delete failed with: %s", err)
	}
	return nil
}

// Commit applies the queued operations in order and frees the batch.
func (b Batch) Commit() error {
	if err := databaseSym.DatabaseBatchCommit(uint32(b)); err != 0 {
		return fmt.Errorf("Batch Commit failed with: %s", err)
	}
	return nil
}

// Entries returns every key under prefix with its value.
func (d Database) Entries(prefix string) ([]Entry, error) {
	data, sizeErr, dataErr := sized.Read(
		func(size *uint32) errno.Error { return databaseSym.DatabaseListEntriesSize(uint32(d), prefix, size) },
		func(data *byte) errno.Error { return databaseSym.DatabaseListEntries(uint32(d), prefix, data) },
	)
	if sizeErr != 0 {
		return nil, fmt.Errorf("Database entries size failed with: %s", sizeErr)
	}
	if dataErr != 0 {
		return nil, fmt.Errorf("Database Entries failed with: %s", dataErr)
	}
	entries := []Entry{}
	// each entry: key, 0x00, value length (uint32 LE), value
	for i := 0; i < len(data); {
		end := i
		for end < len(data) && data[end] != 0 {
			end++
		}
		if end+5 > len(data) {
			return nil, errors.New("Database Entries: malformed data")
		}
		n := int(binary.LittleEndian.Uint32(data[end+1:]))
		v := end + 5
		if n < 0 || v+n > len(data) {
			return nil, errors.New("Database Entries: malformed data")
		}
		entries = append(entries, Entry{Key: string(data[i:end]), Value: data[v : v+n]})
		i = v + n
	}
	return entries, nil
}

func decodeKeys(data []byte) ([]string, error) {
	keys := make([]string, 0)
	if err := codec.Convert(data).To(&keys); err != nil {
		return nil, fmt.Errorf("Failed converting bytes to []string with: %s", err)
	}
	return keys, nil
}

// Match returns the keys under prefix that match the regular expression pattern.
func (d Database) Match(prefix, pattern string) ([]string, error) {
	data, sizeErr, dataErr := sized.Read(
		func(size *uint32) errno.Error { return databaseSym.DatabaseListMatchSize(uint32(d), prefix, pattern, size) },
		func(data *byte) errno.Error { return databaseSym.DatabaseListMatch(uint32(d), prefix, pattern, data) },
	)
	if sizeErr != 0 {
		return nil, fmt.Errorf("Database match size failed with: %s", sizeErr)
	}
	if dataErr != 0 {
		return nil, fmt.Errorf("Database Match failed with: %s", dataErr)
	}
	if len(data) == 0 {
		return []string{}, nil
	}
	return decodeKeys(data)
}

// ListPage returns at most limit keys under prefix, in order, after cursor
// (exclusive; "" starts at the first key). Pass the last key as the next cursor.
func (d Database) ListPage(prefix, cursor string, limit uint32) ([]string, error) {
	data, sizeErr, dataErr := sized.Read(
		func(size *uint32) errno.Error { return databaseSym.DatabaseListPageSize(uint32(d), prefix, cursor, limit, size) },
		func(data *byte) errno.Error { return databaseSym.DatabaseListPage(uint32(d), prefix, cursor, limit, data) },
	)
	if sizeErr != 0 {
		return nil, fmt.Errorf("Database page size failed with: %s", sizeErr)
	}
	if dataErr != 0 {
		return nil, fmt.Errorf("Database ListPage failed with: %s", dataErr)
	}
	if len(data) == 0 {
		return []string{}, nil
	}
	return decodeKeys(data)
}

// Count returns the number of keys under prefix.
func (d Database) Count(prefix string) (uint32, error) {
	var n uint32
	if err := databaseSym.DatabaseCount(uint32(d), prefix, &n); err != 0 {
		return 0, fmt.Errorf("Database Count failed with: %s", err)
	}
	return n, nil
}
