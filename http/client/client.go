package client

import (
	"fmt"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
)

func New() (HttpClient, error) {
	var clientId uint32
	err := httpClientSym.NewHttpClient(&clientId)
	if err != 0 {
		return 0, fmt.Errorf("Creating a new http client failed with: %s", err)
	}

	return HttpClient(clientId), nil
}

func (c HttpClient) Request(url string, options ...HttpRequestOption) (HttpRequest, error) {
	var requestId uint32
	err := httpClientSym.NewHttpRequest(uint32(c), &requestId)
	if err != 0 {
		return HttpRequest{}, fmt.Errorf("Creating a new http request failed with: %s", err)
	}

	r := HttpRequest{client: c, id: requestId}
	err = httpClientSym.SetHttpRequestURL(uint32(c), requestId, url)
	if err != 0 {
		return HttpRequest{}, fmt.Errorf("Setting request url failed with: %s", err)
	}

	for _, option := range options {
		if err := option(r); err != nil {
			return HttpRequest{}, err
		}
	}

	return r, nil
}
