package event

import (
	"testing"
	"unsafe"

	symbols "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestHeaderSet(t *testing.T) {
	var headers HttpEventHeaders

	symbols.EventHttpHeaderAdd = func(eventId uint32, key, val string) (error errno.Error) {
		return 0
	}

	err := headers.Set("content-type", "json")
	if err != nil {
		t.Error(err)
		return
	}

	err = headers.Set("", "json")
	if err == nil {
		t.Error("Expected an error")
		return
	}

	symbols.EventHttpHeaderAdd = func(eventId uint32, key, val string) (error errno.Error) {
		return errno.ErrorHeaderNotFound
	}

	err = headers.Set("content-type", "json")
	if err == nil {
		t.Error("Expected an error")
		return
	}
}

func TestHeaderGet(t *testing.T) {
	var headers HttpEventHeaders
	symbols.MockData{EventId: 1}.Mock()

	_, err := headers.Get("Content-Type")
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Size 0
	headers = HttpEventHeaders(1)
	contentType, err := headers.Get("Content-Type")
	if err != nil {
		t.Error(err)
		return
	}

	if contentType != "" {
		t.Error("Expected empty string")
		return
	}

	err = headers.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
		return
	}

	// Get error
	symbols.GetHttpEventHeaderByName = func(eventId uint32, key string, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return 1
	}

	contentType, err = headers.Get("Content-Type")
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestHeaderList(t *testing.T) {
	var headers HttpEventHeaders
	symbols.MockData{EventId: 1}.Mock()

	_, err := headers.List()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Size 0
	headers = HttpEventHeaders(1)
	list, err := headers.List()
	if err != nil {
		t.Error(err)
		return
	}

	if len(list) != 0 {
		t.Error("Expected empty list")
		return
	}

	err = headers.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
		return
	}

	// Get error
	symbols.GetHttpEventRequestHeaderKeys = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		return 1
	}
	list, err = headers.List()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Conversion error
	symbols.GetHttpEventRequestHeaderKeysSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	symbols.GetHttpEventRequestHeaderKeys = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		d := unsafe.Slice(bufPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}

	list, err = headers.List()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
