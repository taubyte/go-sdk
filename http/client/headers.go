package client

import (
	"fmt"
	"strings"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/utils/codec"
)

func Headers(headers map[string][]string) HttpRequestOption {
	return func(r HttpRequest) error {
		for k, v := range headers {
			if err := r.Headers().Set(k, v...); err != nil {
				return err
			}
		}

		return nil
	}
}

func (r *HttpRequestHeaders) Set(key string, values ...string) error {
	// Returning nil here, because you can't set a header to empty
	if len(values) == 0 {
		return nil
	}

	var byteValues []byte
	// Ignoring error here as values is always a non-empty string slice
	codec.Convert(values).To(&byteValues)

	err := httpClientSym.SetHttpRequestHeader(uint32(r.client), r.id, key, &byteValues[0], uint32(len(byteValues)))
	if err != 0 {
		return fmt.Errorf("Set header `%s`:`%s`, failed with: %s", key, strings.Join(values, "; "), err)
	}

	return nil
}

func (r *HttpRequestHeaders) Add(key, value string) error {
	err := httpClientSym.AddHttpRequestHeader(uint32(r.client), r.id, key, value)
	if err != 0 {
		return fmt.Errorf("Add header `%s`:`%s`, failed with: %s", key, value, err)
	}

	return nil
}

func (r *HttpRequestHeaders) Get(key string) ([]string, error) {
	var size uint32
	if err := httpClientSym.GetHttpRequestHeaderSize(uint32(r.client), r.id, key, &size); err != 0 {
		return nil, fmt.Errorf("Get header size failed with: %s", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	byteHeaders := make([]byte, size)
	if err := httpClientSym.GetHttpRequestHeader(uint32(r.client), r.id, key, &byteHeaders[0]); err != 0 {
		return nil, fmt.Errorf("Getting header for key: %s failed with error: %s", key, err)
	}

	var headers []string
	err := codec.Convert(byteHeaders).To(&headers)
	if err != nil || len(headers) == 0 {
		return nil, fmt.Errorf("Conversion Get(`%s`) failed with: %s", key, err)
	}

	return headers, nil
}

func (r *HttpRequestHeaders) List() ([]string, error) {
	var size uint32
	if err := httpClientSym.GetHttpRequestHeaderKeysSize(uint32(r.client), r.id, &size); err != 0 {
		return nil, fmt.Errorf("Getting all header keys size failed with: %s", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	header := make([]byte, size)
	if err := httpClientSym.GetHttpRequestHeaderKeys(uint32(r.client), r.id, &header[0], uint32(len(header))); err != 0 {
		return nil, fmt.Errorf("Getting all header keys failed with: %s", err)
	}

	var value []string
	err := codec.Convert(header).To(&value)
	if err != nil || len(value) == 0 {
		return nil, fmt.Errorf("Converting header slice to string slice failed with: %s", err)
	}

	return value, nil
}

func (r *HttpRequestHeaders) GetAll() (map[string][]string, error) {
	var size uint32
	if err := httpClientSym.GetHttpRequestHeaderKeysSize(uint32(r.client), r.id, &size); err != 0 {
		return nil, fmt.Errorf("Getting all header keys size failed with: %s", err)
	}
	if size == 0 {
		return map[string][]string{}, nil
	}

	header := make([]byte, size)
	if err := httpClientSym.GetHttpRequestHeaderKeys(uint32(r.client), r.id, &header[0], uint32(len(header))); err != 0 {
		return nil, fmt.Errorf("Getting all header keys failed with: %s", err)
	}

	var _value []string
	err := codec.Convert(header).To(&_value)
	if err != nil || len(_value) == 0 {
		return nil, fmt.Errorf("Converting header slice to string slice failed with: %s", err)
	}

	var err0 error
	value := make(map[string][]string, len(_value))
	for _, h := range _value {
		value[h], err0 = r.Get(h)
		if err0 != nil {
			return value, err0
		}
	}

	return value, nil
}
