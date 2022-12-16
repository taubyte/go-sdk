package client

import (
	"fmt"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
)

func Body(data []byte) HttpRequestOption {
	return func(r HttpRequest) error {
		if err := r.Body().Set(data); err != nil {
			return err
		}

		return nil
	}
}

func (r *HttpRequestBody) Set(data []byte) error {
	var dataPtr *byte
	if len(data) != 0 {
		dataPtr = &data[0]
	}

	err := httpClientSym.SetHttpRequestBody(uint32(r.client), r.id, dataPtr, uint32(len(data)))
	if err != 0 {
		return fmt.Errorf("Set Http Request body failed with %s", err)
	}

	return nil
}

// TODO: implement
// func (r *HttpRequestBody) Get() ([]byte, error) {
// 	var size uint32
// 	err := httpClientSym.GetHttpRequestBodySize(uint32(r.client), r.id, &size)
// 	if err != 0 {
// 		return nil, fmt.Errorf("Get Http Request body size failed with %s", err)
// 	}

// 	data := make([]byte, size)
// 	var dataPtr *byte
// 	if len(data) != 0 {
// 		dataPtr = &data[0]
// 	}

// 	err = httpClientSym.GetHttpRequestBody(uint32(r.client), r.id, dataPtr, &size)
// 	if err != 0 {
// 		return nil, fmt.Errorf("Get Http Request body failed with %s", err)
// 	}

// 	return data, nil
// }

// func (r *HttpRequestBody) Write(io.Reader?) (err error) {
// 	return
// }
