package client

import (
	"fmt"
	"io"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/errno"
)

func (resp *HttpResponse) Body() *HttpResponseBody {
	return &HttpResponseBody{
		request: resp.request,
	}
}

func (resp *HttpResponseBody) Read(p []byte) (int, error) {
	var counter uint32
	err := httpClientSym.ReadHttpResponseBody(uint32(resp.request.client), resp.request.id, &p[0], uint32(len(p)), &counter)
	if err != 0 {
		if err == errno.ErrorEOF {
			return int(counter), io.EOF
		} else {
			return 0, fmt.Errorf("Reading HTTP response body failed with: %s", err)
		}
	}

	return int(counter), nil
}

func (resp *HttpResponseBody) Close() error {
	clientId := uint32(resp.request.client)
	reqId := resp.request.id

	err := httpClientSym.CloseHttpResponseBody(clientId, reqId)
	if err != (0) {
		return fmt.Errorf("Closing HTTP response body failed with: %s", err)
	}
	return nil
}
