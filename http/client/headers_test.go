package client

import (
	"testing"
	"unsafe"

	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestHeadersSet(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()
	headers := &HttpRequestHeaders{}

	if headers.Set("fruits", "apple", "orange") == nil {
		t.Error("Expected error")
		return
	}

	if headers.Set("fruits") != nil {
		return
	}
}

func TestHeadersAdd(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()
	headers := &HttpRequestHeaders{}

	if headers.Add("fruits", "apple") == nil {
		t.Error("Expected error")
		return
	}
}

func TestHeadersGet(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()
	headers := &HttpRequestHeaders{}

	_, err := headers.Get("fruits")
	if err == nil {
		t.Error("Expected error")
		return
	}

	headers.client = 1
	err = headers.Set("fruits")
	if err != nil {
		t.Error(err)
		return
	}

	// Empty
	_, err = headers.Get("fruits")
	if err != nil {
		t.Error(err)
		return
	}

	err = headers.Set("fruits", "banana")
	if err != nil {
		t.Error(err)
		return
	}

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
	symbols.GetHttpRequestHeaderSize = func(clientId, requestId uint32, key string, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	symbols.GetHttpRequestHeader = func(clientId, requestId uint32, key string, headerPtr *byte) (error errno.Error) {
		d := unsafe.Slice(headerPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}
	_, err = headers.Get("fruits")
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.GetHttpRequestHeader = func(clientId, requestId uint32, key string, headerPtr *byte) (error errno.Error) {
		return 1
	}
	_, err = headers.Get("fruits")
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestHeadersList(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()
	headers := &HttpRequestHeaders{}

	_, err := headers.List()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Empty
	headers.client = 1
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
	symbols.GetHttpRequestHeaderKeysSize = func(clientId, requestId uint32, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	symbols.GetHttpRequestHeaderKeys = func(clientId, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
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
	symbols.GetHttpRequestHeaderKeys = func(clientId, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
		return 1
	}
	_, err = headers.List()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestHeadersGetAll(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()
	headers := &HttpRequestHeaders{}

	_, err := headers.GetAll()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Empty
	headers.client = 1
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
	symbols.GetHttpRequestHeaderKeysSize = func(clientId, requestId uint32, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	symbols.GetHttpRequestHeaderKeys = func(clientId, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
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
	symbols.GetHttpRequestHeaderKeys = func(clientId, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
		return 1
	}
	_, err = headers.GetAll()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockData{ClientId: 1}.Mock()
	err = headers.Set("fruits", "apple")
	if err != nil {
		t.Error(err)
		return
	}
	symbols.GetHttpRequestHeader = func(clientId, requestId uint32, key string, headerPtr *byte) (error errno.Error) {
		return 1
	}
	_, err = headers.GetAll()
	if err == nil {
		t.Error("Expected error")
		return
	}

}
