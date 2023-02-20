package event

import (
	"testing"
	"unsafe"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestQueryGet(t *testing.T) {
	var e Event

	httpEventSym.MockData{
		EventId: 3,
		Queries: map[string]string{
			"name":     "taubyte",
			"password": "qwerty",
			"empty":    "",
		}}.Mock()

	_, err := e.Query().Get("name")
	if err == nil {
		t.Error("Expected error")
		return
	}

	e = 3
	empty, err := e.Query().Get("empty")
	if err != nil {
		t.Error(err)
		return
	}
	if empty != "" {
		t.Error("Expected empty string")
		return
	}

	httpEventSym.GetHttpEventQueryValueByName = func(eventId uint32, key string, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return 1
	}
	_, err = e.Query().Get("password")
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestQueryList(t *testing.T) {
	var e Event = 3

	m := httpEventSym.MockData{
		EventId: uint32(e),
	}.Mock()

	// Empty
	queries, err := e.Query().List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(queries) != 0 {
		t.Error("Expected no queries got", queries)
		return
	}

	m.Queries = map[string]string{
		"name":     "taubyte",
		"password": "qwerty",
	}
	m.Mock()
	e = 0
	_, err = e.Query().List()
	if err == nil {
		t.Error("Expected error")
		return
	}

	httpEventSym.GetHttpEventRequestQueryKeys = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		return 1
	}

	e = 3
	_, err = e.Query().List()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Conversion error
	httpEventSym.GetHttpEventRequestQueryKeysSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	httpEventSym.GetHttpEventRequestQueryKeys = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		d := unsafe.Slice(bufPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}

	_, err = e.Query().List()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
