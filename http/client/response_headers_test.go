package client

import (
	"testing"
	"unsafe"

	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestResponseHeadersGet(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()
	headers := &HttpResponseHeaders{request: &HttpRequest{}}

	_, err := headers.Get("fruits")
	if err == nil {
		t.Error("Expected error")
		return
	}

	headers.request.client = 1
	// Empty
	_, err = headers.Get("fruits")
	if err != nil {
		t.Error(err)
		return
	}

	symbols.MockData{ClientId: 1, ResponseHeaders: map[string][]string{"fruits": {"banana"}}}.Mock()
	_headers, err := headers.Get("fruits")
	if err != nil {
		t.Error(err)
		return
	}
	if len(_headers) != 1 {
		t.Errorf("Expected 1 header got: %d", len(_headers))
		return
	}

	// Conversion error
	symbols.GetHttpResponseHeaderSize = func(clientId, requestId uint32, key string, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	symbols.GetHttpResponseHeader = func(clientId, requestId uint32, key string, headerPtr *byte) (error errno.Error) {
		d := unsafe.Slice(headerPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}
	h, err := headers.Get("fruits")
	if err == nil {
		t.Error("Expected error", h)
		return
	}

	symbols.GetHttpResponseHeader = func(clientId, requestId uint32, key string, headerPtr *byte) (error errno.Error) {
		return 1
	}
	_, err = headers.Get("fruits")
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestResponseHeadersList(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()
	headers := &HttpResponseHeaders{request: &HttpRequest{}}

	_, err := headers.List()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Empty
	headers.request.client = 1
	_headers, err := headers.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(_headers) != 0 {
		t.Errorf("Expected no headers got %v", _headers)
		return
	}

	// Conversion error
	symbols.GetHttpResponseHeaderKeysSize = func(clientId, requestId uint32, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	symbols.GetHttpResponseHeaderKeys = func(clientId, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
		d := unsafe.Slice(headerPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}
	_, err = headers.List()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Basic error
	symbols.GetHttpResponseHeaderKeys = func(clientId, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
		return 1
	}
	_, err = headers.List()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestResponseHeadersGetAll(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()
	headers := &HttpResponseHeaders{request: &HttpRequest{}}

	_, err := headers.GetAll()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Empty
	headers.request.client = 1
	_headers, err := headers.GetAll()
	if err != nil {
		t.Error(err)
		return
	}
	if len(_headers) != 0 {
		t.Error("Expected headers to be empty")
		return
	}

	// Conversion error
	symbols.GetHttpResponseHeaderKeysSize = func(clientId, requestId uint32, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	symbols.GetHttpResponseHeaderKeys = func(clientId, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
		d := unsafe.Slice(headerPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}
	_, err = headers.GetAll()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Basic errors
	symbols.GetHttpResponseHeaderKeys = func(clientId, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
		return 1
	}
	_, err = headers.GetAll()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockData{ClientId: 1, ResponseHeaders: map[string][]string{"fruits": {"apple"}}}.Mock()
	symbols.GetHttpResponseHeader = func(clientId, requestId uint32, key string, headerPtr *byte) (error errno.Error) {
		return 1
	}
	_, err = headers.GetAll()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
