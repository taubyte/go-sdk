package client

import (
	"fmt"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
)

func (r *HttpRequest) Do() (*HttpResponse, error) {
	err := httpClientSym.DoHttpRequest(uint32(r.client), uint32(r.id))
	if err != 0 {
		return nil, fmt.Errorf("Executing HTTP request failed with: %s", err)
	}

	return &HttpResponse{r}, nil
}

func (r *HttpRequest) Method() *HttpRequestMethod {
	return &HttpRequestMethod{
		id:     r.id,
		client: r.client,
	}
}

func (r *HttpRequest) Headers() *HttpRequestHeaders {
	return &HttpRequestHeaders{
		id:     r.id,
		client: r.client,
	}
}

func (r *HttpRequest) Body() *HttpRequestBody {
	return &HttpRequestBody{
		id:     r.id,
		client: r.client,
	}
}
