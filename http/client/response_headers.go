package client

import (
	"fmt"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/utils/codec"
)

func (resp *HttpResponse) Headers() *HttpResponseHeaders {
	return &HttpResponseHeaders{
		request: resp.request,
	}
}

func (h *HttpResponseHeaders) Get(key string) ([]string, error) {
	r := h.request

	var size uint32
	if err := httpClientSym.GetHttpResponseHeaderSize(uint32(r.client), r.id, key, &size); err != 0 {
		return nil, fmt.Errorf("Get header size failed with: %s", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	byteHeaders := make([]byte, size)
	if err := httpClientSym.GetHttpResponseHeader(uint32(r.client), r.id, key, &byteHeaders[0]); err != 0 {
		return nil, fmt.Errorf("Getting header for key: %s failed with error: %s", key, err)
	}

	var headers []string
	err := codec.Convert(byteHeaders).To(&headers)
	if err != nil || len(headers) == 0 {
		return nil, fmt.Errorf("Conversion Get(`%s`) failed with: %s", key, err)
	}

	return headers, nil
}

func (h *HttpResponseHeaders) List() ([]string, error) {
	r := h.request

	var size uint32
	if err := httpClientSym.GetHttpResponseHeaderKeysSize(uint32(r.client), r.id, &size); err != 0 {
		return nil, fmt.Errorf("Getting all header keys size failed with: %s", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	header := make([]byte, size)
	if err := httpClientSym.GetHttpResponseHeaderKeys(uint32(r.client), r.id, &header[0], uint32(len(header))); err != 0 {
		return nil, fmt.Errorf("Getting all header keys failed with: %s", err)
	}

	var value []string
	err := codec.Convert(header).To(&value)
	if err != nil || len(value) == 0 {
		return nil, fmt.Errorf("Converting header slice to string slice failed with: %s", err)
	}

	return value, nil
}

func (h *HttpResponseHeaders) GetAll() (map[string][]string, error) {
	r := h.request

	var size uint32
	if err := httpClientSym.GetHttpResponseHeaderKeysSize(uint32(r.client), r.id, &size); err != 0 {
		return nil, fmt.Errorf("Getting all header keys size failed with: %s", err)
	}
	if size == 0 {
		return map[string][]string{}, nil
	}

	header := make([]byte, size)
	if err := httpClientSym.GetHttpResponseHeaderKeys(uint32(r.client), r.id, &header[0], uint32(len(header))); err != 0 {
		return nil, fmt.Errorf("Getting all header keys failed with: %s", err)
	}

	var _value []string
	err := codec.Convert(header).To(&_value)
	if err != nil || len(_value) == 0 {
		return nil, fmt.Errorf("Converting header slice to string slice failed with: %s", err)
	}

	var err0 error
	value := make(map[string][]string, len(_value))
	for _, header := range _value {
		value[header], err0 = h.Get(header)
		if err0 != nil {
			return value, err0
		}
	}

	return value, nil
}
