package event

import (
	"errors"
	"fmt"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/utils/codec"
)

func (e Event) Headers() EventHeaders {
	return EventHeaders(e)
}

func (e EventHeaders) Set(key, value string) error {
	if len(key) == 0 {
		return errors.New("Cannot set header with empty key")
	}

	err := httpEventSym.EventHttpHeaderAdd(uint32(e), key, value)
	if err != 0 {
		return fmt.Errorf("Setting header failed with: %s", err)
	}
	return nil
}

func (e EventHeaders) Get(key string) (string, error) {
	var size uint32
	err := httpEventSym.GetHttpEventHeaderByNameSize(uint32(e), &size, key)
	if err != 0 {
		return "", fmt.Errorf("Getting Header Size By Name for key:`%s` Failed with:%s", key, err)
	}
	if size == 0 {
		return "", nil
	}

	headers := make([]byte, size)
	err = httpEventSym.GetHttpEventHeaderByName(uint32(e), key, &headers[0], uint32(len(headers)))
	if err != 0 {
		return "", fmt.Errorf("Getting Header By Name for key:`%s` Failed with:%s", key, err)
	}

	return string(headers), nil
}

func (e EventHeaders) List() ([]string, error) {
	var size uint32
	err := httpEventSym.GetHttpEventRequestHeaderKeysSize(uint32(e), &size)
	if err != 0 {
		return nil, fmt.Errorf("Getting all header keys size failed with: %s", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	header := make([]byte, size)
	if err := httpEventSym.GetHttpEventRequestHeaderKeys(uint32(e), &header[0]); err != 0 {
		return nil, fmt.Errorf("Getting all header keys failed with: %s", err)
	}

	var headers []string
	err0 := codec.Convert(header).To(&headers)
	if err0 != nil || len(headers) == 0 {
		return nil, fmt.Errorf("Converting header slice to string slice failed with: %s", err0)
	}

	return headers, nil
}
