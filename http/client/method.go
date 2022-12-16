package client

import (
	"fmt"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/utils/convert"
)

func Method(method string) HttpRequestOption {
	return func(r HttpRequest) error {
		return r.Method().Set(method)
	}
}

func (r *HttpRequestMethod) Set(method string) error {
	_method, err := convert.MethodStringToUint(method)
	if err != nil {
		return err
	}

	err0 := httpClientSym.SetHttpRequestMethod(uint32(r.client), r.id, _method)
	if err0 != 0 {
		return fmt.Errorf("Setting HTTP request method failed with: %s", err0)
	}

	return nil
}

func (r *HttpRequestMethod) Get() (string, error) {
	var method uint32
	err := httpClientSym.GetHttpRequestMethod(uint32(r.client), r.id, &method)
	if err != 0 {
		return "", fmt.Errorf("Getting HTTP request method failed with: %s", err)
	}
	_method, err0 := convert.MethodUintToString(method)
	if err0 != nil {
		return "", err0
	}

	return _method, nil
}
