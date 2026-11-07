package database

import (
	"encoding/binary"
	"errors"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

func putBytes(b []byte, data *byte) { copy(unsafe.Slice(data, len(b)), b) }

func keysBytes(t *testing.T, keys []string) []byte {
	var b []byte
	if err := codec.Convert(keys).To(&b); err != nil {
		t.Fatal(err)
	}
	return b
}

func TestBatch(t *testing.T) {
	var ops []string
	databaseSym.DatabaseBatchNew = func(id uint32, b *uint32) errno.Error { *b = 9; return 0 }
	databaseSym.DatabaseBatchPut = func(b uint32, k string, d *byte, n uint32) errno.Error {
		ops = append(ops, "put "+k+" "+string(unsafe.Slice(d, n)))
		return 0
	}
	databaseSym.DatabaseBatchDelete = func(b uint32, k string) errno.Error { ops = append(ops, "del "+k); return 0 }
	databaseSym.DatabaseBatchCommit = func(b uint32) errno.Error { ops = append(ops, "commit"); return 0 }

	b, err := Database(1).Batch()
	if err != nil || b != 9 {
		t.Fatal(b, err)
	}
	if b.Put("a", []byte("x")) != nil || b.Delete("b") != nil || b.Commit() != nil {
		t.Fatal("unexpected error")
	}
	if want := []string{"put a x", "del b", "commit"}; !reflect.DeepEqual(ops, want) {
		t.Fatal(ops)
	}

	databaseSym.DatabaseBatchPut = func(uint32, string, *byte, uint32) errno.Error { return 1 }
	databaseSym.DatabaseBatchDelete = func(uint32, string) errno.Error { return 1 }
	databaseSym.DatabaseBatchCommit = func(uint32) errno.Error { return 1 }
	databaseSym.DatabaseBatchNew = func(uint32, *uint32) errno.Error { return 1 }
	if b.Put("a", nil) == nil || b.Delete("a") == nil || b.Commit() == nil {
		t.Fatal("expected errors")
	}
	if _, err := Database(1).Batch(); err == nil {
		t.Fatal("expected error")
	}
}

func TestEntries(t *testing.T) {
	var enc []byte
	for _, e := range []Entry{{"k1", []byte("v1")}, {"k2", nil}, {"k3", []byte("v3")}} {
		enc = append(enc, e.Key...)
		enc = append(enc, 0)
		enc = binary.LittleEndian.AppendUint32(enc, uint32(len(e.Value)))
		enc = append(enc, e.Value...)
	}
	databaseSym.DatabaseListEntriesSize = func(id uint32, p string, s *uint32) errno.Error { *s = uint32(len(enc)); return 0 }
	databaseSym.DatabaseListEntries = func(id uint32, p string, d *byte) errno.Error { putBytes(enc, d); return 0 }
	got, err := Database(1).Entries("k")
	if err != nil || len(got) != 3 || got[0].Key != "k1" || string(got[0].Value) != "v1" || len(got[1].Value) != 0 || string(got[2].Value) != "v3" {
		t.Fatal(got, err)
	}

	enc = enc[:len(enc)-1] // truncated value
	if _, err := Database(1).Entries("k"); err == nil {
		t.Fatal("expected malformed error")
	}

	databaseSym.DatabaseListEntriesSize = func(uint32, string, *uint32) errno.Error { return 0 }
	if got, err := Database(1).Entries("k"); err != nil || len(got) != 0 {
		t.Fatal(got, err)
	}
	databaseSym.DatabaseListEntriesSize = func(uint32, string, *uint32) errno.Error { return 1 }
	if _, err := Database(1).Entries("k"); err == nil {
		t.Fatal("expected error")
	}
}

func TestMatch(t *testing.T) {
	keys := []string{"u/1", "u/22"}
	enc := keysBytes(t, keys)
	databaseSym.DatabaseListMatchSize = func(id uint32, p, pat string, s *uint32) errno.Error {
		if p != "u/" || pat != `\d$` {
			return 1
		}
		*s = uint32(len(enc))
		return 0
	}
	databaseSym.DatabaseListMatch = func(id uint32, p, pat string, d *byte) errno.Error { putBytes(enc, d); return 0 }
	got, err := Database(1).Match("u/", `\d$`)
	if err != nil || !reflect.DeepEqual(got, keys) {
		t.Fatal(got, err)
	}
	if _, err := Database(1).Match("u/", "("); err == nil {
		t.Fatal("expected error")
	}
	databaseSym.DatabaseListMatchSize = func(uint32, string, string, *uint32) errno.Error { return 0 }
	if got, err := Database(1).Match("u/", "x"); err != nil || len(got) != 0 {
		t.Fatal(got, err)
	}
}

func TestListPage(t *testing.T) {
	all := []string{"a", "b", "c", "d", "e"}
	page := func(prefix, cursor string, limit uint32) []byte {
		var out []string
		for _, k := range all {
			if strings.HasPrefix(k, prefix) && k > cursor && uint32(len(out)) < limit {
				out = append(out, k)
			}
		}
		if len(out) == 0 {
			return nil
		}
		return keysBytes(t, out)
	}
	databaseSym.DatabaseListPageSize = func(id uint32, p, c string, l uint32, s *uint32) errno.Error {
		*s = uint32(len(page(p, c, l)))
		return 0
	}
	databaseSym.DatabaseListPage = func(id uint32, p, c string, l uint32, d *byte) errno.Error {
		putBytes(page(p, c, l), d)
		return 0
	}
	var seen []string
	cursor := ""
	for {
		got, err := Database(1).ListPage("", cursor, 2)
		if err != nil {
			t.Fatal(err)
		}
		if len(got) == 0 {
			break
		}
		seen = append(seen, got...)
		cursor = got[len(got)-1]
	}
	if !reflect.DeepEqual(seen, all) {
		t.Fatal(seen)
	}
	databaseSym.DatabaseListPageSize = func(uint32, string, string, uint32, *uint32) errno.Error { return errno.ErrorZeroSize }
	if _, err := Database(1).ListPage("", "", 0); err == nil {
		t.Fatal("expected error")
	}
}

func TestCount(t *testing.T) {
	databaseSym.DatabaseCount = func(id uint32, p string, n *uint32) errno.Error { *n = 7; return 0 }
	if n, err := Database(1).Count("p"); err != nil || n != 7 {
		t.Fatal(n, err)
	}
	databaseSym.DatabaseCount = func(uint32, string, *uint32) errno.Error { return 1 }
	if _, err := Database(1).Count("p"); err == nil {
		t.Fatal("expected error")
	}
}

func TestNotFoundAndHas(t *testing.T) {
	databaseSym.DatabaseGetSize = func(uint32, string, *uint32) errno.Error { return errno.ErrorDatabaseKeyNotFound }
	if _, err := Database(1).Get("x"); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if ok, err := Database(1).Has("x"); ok || err != nil {
		t.Fatal(ok, err)
	}
	databaseSym.DatabaseGetSize = func(uint32, string, *uint32) errno.Error { return 0 }
	if ok, err := Database(1).Has("x"); !ok || err != nil {
		t.Fatal(ok, err)
	}
	databaseSym.DatabaseGetSize = func(uint32, string, *uint32) errno.Error { return errno.ErrorDatabaseGetFailed }
	if _, err := Database(1).Has("x"); err == nil {
		t.Fatal("expected error")
	}
}
