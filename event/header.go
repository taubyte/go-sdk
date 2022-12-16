package event

import (
	"errors"
	"fmt"

	eventSym "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/utils/codec"
)

func (e HttpEvent) Headers() HttpEventHeaders {
	return HttpEventHeaders(e)
}

func (e HttpEventHeaders) Set(key, value string) error {
	if len(key) == 0 {
		return errors.New("Cannot set header with empty key")
	}

	err := eventSym.EventHttpHeaderAdd(uint32(e), key, value)
	if err != 0 {
		return fmt.Errorf("Setting header failed with: %s", err)
	}
	return nil
}

func (e HttpEventHeaders) Get(key string) (string, error) {
	var size uint32
	err := eventSym.GetHttpEventHeaderByNameSize(uint32(e), &size, key)
	if err != 0 {
		return "", fmt.Errorf("Getting Header Size By Name for key:`%s` Failed with:%s", key, err)
	}
	if size == 0 {
		return "", nil
	}

	headers := make([]byte, size)
	err = eventSym.GetHttpEventHeaderByName(uint32(e), key, &headers[0], uint32(len(headers)))
	if err != 0 {
		return "", fmt.Errorf("Getting Header By Name for key:`%s` Failed with:%s", key, err)
	}

	return string(headers), nil
}

func (e HttpEventHeaders) List() ([]string, error) {
	var size uint32
	err := eventSym.GetHttpEventRequestHeaderKeysSize(uint32(e), &size)
	if err != 0 {
		return nil, fmt.Errorf("Getting all header keys size failed with: %s", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	header := make([]byte, size)
	if err := eventSym.GetHttpEventRequestHeaderKeys(uint32(e), &header[0]); err != 0 {
		return nil, fmt.Errorf("Getting all header keys failed with: %s", err)
	}

	var headers []string
	err0 := codec.Convert(header).To(&headers)
	if err0 != nil || len(headers) == 0 {
		return nil, fmt.Errorf("Converting header slice to string slice failed with: %s", err0)
	}

	return headers, nil
}
