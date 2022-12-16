package client

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/http/client"

	"github.com/taubyte/go-sdk/errno"
)

func TestResponseClose(t *testing.T) {
	symbols.MockData{ClientId: 1, RequestId: 1}.Mock()

	request := HttpRequest{id: 1, client: 1}
	response, err := request.Do()
	if err != nil {
		t.Error(err)
		return
	}

	symbols.CloseHttpResponseBody = func(clientId uint32, requestId uint32) (error errno.Error) {
		return 1
	}
	err = response.Body().Close()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Basic read error
	symbols.ReadHttpResponseBody = func(clientId, requestId uint32, buf *byte, bufSize uint32, countPtr *uint32) (error errno.Error) {
		return 1
	}
	_, err = response.Body().Read(make([]byte, 10))
	if err == nil {
		t.Error("Expected error")
		return
	}
}
