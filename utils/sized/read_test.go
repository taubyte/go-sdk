package sized_test

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
	globalSym "github.com/taubyte/go-sdk-symbols/globals"
	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/database"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/globals/str"
	"github.com/taubyte/go-sdk/storage"
	"github.com/taubyte/go-sdk/utils/codec"
)

const maxAttempts = 5

type sizeFn = func(*uint32) errno.Error
type dataFn = func(*byte) errno.Error

// call is one SDK call that reads a value in two host steps.
type call struct {
	name    string
	value   []byte
	want    string
	install func(sizeFn, dataFn)
	run     func() (any, error)
}

func calls(t *testing.T) []call {
	var list []byte
	if err := codec.Convert([]string{"a", "b"}).To(&list); err != nil {
		t.Fatal(err)
	}
	listWant := "[a b]"
	entry := append([]byte("k\x00"), 1, 0, 0, 0, 'v')
	db, st := database.Database(1), storage.Storage(1)

	return []call{
		{"Get", []byte("value"), "value", func(s sizeFn, d dataFn) {
			databaseSym.DatabaseGetSize = func(_ uint32, _ string, p *uint32) errno.Error { return s(p) }
			databaseSym.DatabaseGet = func(_ uint32, _ string, p *byte) errno.Error { return d(p) }
		}, func() (any, error) { v, err := db.Get("k"); return string(v), err }},
		{"List", list, listWant, func(s sizeFn, d dataFn) {
			databaseSym.DatabaseListSize = func(_ uint32, _ string, p *uint32) errno.Error { return s(p) }
			databaseSym.DatabaseList = func(_ uint32, _ string, p *byte) errno.Error { return d(p) }
		}, func() (any, error) { return db.List("p") }},
		{"Entries", entry, "[{k [118]}]", func(s sizeFn, d dataFn) {
			databaseSym.DatabaseListEntriesSize = func(_ uint32, _ string, p *uint32) errno.Error { return s(p) }
			databaseSym.DatabaseListEntries = func(_ uint32, _ string, p *byte) errno.Error { return d(p) }
		}, func() (any, error) { return db.Entries("p") }},
		{"Match", list, listWant, func(s sizeFn, d dataFn) {
			databaseSym.DatabaseListMatchSize = func(_ uint32, _, _ string, p *uint32) errno.Error { return s(p) }
			databaseSym.DatabaseListMatch = func(_ uint32, _, _ string, p *byte) errno.Error { return d(p) }
		}, func() (any, error) { return db.Match("p", "x") }},
		{"ListPage", list, listWant, func(s sizeFn, d dataFn) {
			databaseSym.DatabaseListPageSize = func(_ uint32, _, _ string, _ uint32, p *uint32) errno.Error { return s(p) }
			databaseSym.DatabaseListPage = func(_ uint32, _, _ string, _ uint32, p *byte) errno.Error { return d(p) }
		}, func() (any, error) { return db.ListPage("p", "", 2) }},
		{"GlobalGet", []byte("value"), "value", func(s sizeFn, d dataFn) {
			globalSym.GetGlobalValueSize = func(_ string, _, _ uint32, p *uint32) errno.Error { return s(p) }
			globalSym.GetGlobalValue = func(_ string, _, _ uint32, p *byte) errno.Error { return d(p) }
		}, func() (any, error) { v, err := str.Get("n"); return v.Value(), err }},
		{"GlobalGetOrCreate", []byte("value"), "value", func(s sizeFn, d dataFn) {
			globalSym.GetOrCreateGlobalValueSize = func(_ string, _, _ uint32, p *uint32) errno.Error { return s(p) }
			globalSym.GetGlobalValue = func(_ string, _, _ uint32, p *byte) errno.Error { return d(p) }
		}, func() (any, error) { v, err := str.GetOrCreate("n"); return v.Value(), err }},
		{"ListFiles", []byte("/file/f/1"), "1", func(s sizeFn, d dataFn) {
			storageSym.StorageListFilesSize = func(_ uint32, p *uint32) errno.Error { return s(p) }
			storageSym.StorageListFiles = func(_ uint32, p *byte) errno.Error { return d(p) }
		}, func() (any, error) { v, err := st.ListFiles(); return len(v), err }},
		{"Versions", list, listWant, func(s sizeFn, d dataFn) {
			storageSym.StorageListVersionsSize = func(_ uint32, _ string, p *uint32) errno.Error { return s(p) }
			storageSym.StorageListVersions = func(_ uint32, _ string, p *byte) errno.Error { return d(p) }
		}, func() (any, error) { return st.File("f").Versions() }},
		{"RemainingCapacity", []byte("13"), "7", func(s sizeFn, d dataFn) {
			storageSym.StorageUsedSize = func(_ uint32, p *uint32) errno.Error { return s(p) }
			storageSym.StorageUsed = func(_ uint32, p *byte) errno.Error { return d(p) }
			storageSym.StorageCapacitySize = func(_ uint32, p *uint32) errno.Error { *p = 2; return 0 }
			storageSym.StorageCapacity = func(_ uint32, p *byte) errno.Error {
				copy(unsafe.Slice(p, 2), "20")
				return 0
			}
		}, func() (any, error) { return st.RemainingCapacity() }},
	}
}

// host mimics the real host: the first short answers report a size too small
// for value, and the data step then refuses with ErrorSizeMismatch.
type host struct {
	value            []byte
	short            int
	sizeErr, dataErr errno.Error
	reported         uint32
	sizes, datas     int
}

func (h *host) size(p *uint32) errno.Error {
	h.sizes++
	if h.sizeErr != 0 {
		return h.sizeErr
	}
	h.reported = uint32(len(h.value))
	if h.sizes <= h.short {
		h.reported = 1
	}
	*p = h.reported
	return 0
}

func (h *host) data(p *byte) errno.Error {
	h.datas++
	if h.dataErr != 0 {
		return h.dataErr
	}
	if h.reported < uint32(len(h.value)) {
		return errno.ErrorSizeMismatch
	}
	copy(unsafe.Slice(p, h.reported), h.value)
	return 0
}

func TestSettlesAfterMismatch(t *testing.T) {
	for _, c := range calls(t) {
		c := c
		t.Run(c.name, func(t *testing.T) {
			h := &host{value: c.value, short: maxAttempts - 1}
			c.install(h.size, h.data)
			got, err := c.run()
			if err != nil {
				t.Fatal(err)
			}
			if s := fmt.Sprint(got); s != c.want {
				t.Fatalf("got %s, want %s", s, c.want)
			}
			if h.sizes != maxAttempts || h.datas != maxAttempts {
				t.Fatalf("size calls %d, data calls %d, want %d each", h.sizes, h.datas, maxAttempts)
			}
		})
	}
}

func TestGivesUpAtTheCap(t *testing.T) {
	for _, c := range calls(t) {
		c := c
		t.Run(c.name, func(t *testing.T) {
			h := &host{value: c.value, short: 1 << 30}
			c.install(h.size, h.data)
			_, err := c.run()
			if err == nil || !strings.Contains(err.Error(), errno.ErrorSizeMismatch.String()) {
				t.Fatalf("expected a size mismatch error, got %v", err)
			}
			if h.sizes != maxAttempts || h.datas != maxAttempts {
				t.Fatalf("size calls %d, data calls %d, want %d each", h.sizes, h.datas, maxAttempts)
			}
		})
	}
}

func TestOtherErrorsAreNotRetried(t *testing.T) {
	for _, c := range calls(t) {
		c := c
		t.Run(c.name+"/data", func(t *testing.T) {
			h := &host{value: c.value, dataErr: errno.ErrorDatabaseGetFailed}
			c.install(h.size, h.data)
			if _, err := c.run(); err == nil {
				t.Fatal("expected error")
			}
			if h.sizes != 1 || h.datas != 1 {
				t.Fatalf("size calls %d, data calls %d, want 1 each", h.sizes, h.datas)
			}
		})
		t.Run(c.name+"/size", func(t *testing.T) {
			h := &host{value: c.value, sizeErr: errno.ErrorDatabaseGetFailed}
			c.install(h.size, h.data)
			if _, err := c.run(); err == nil {
				t.Fatal("expected error")
			}
			if h.sizes != 1 || h.datas != 0 {
				t.Fatalf("size calls %d, data calls %d, want 1 and 0", h.sizes, h.datas)
			}
		})
	}
}
