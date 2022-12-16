package client

import (
	"testing"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestRequestDo(t *testing.T) {
	_requestId := uint32(6)
	client := HttpClient(3)
	request := HttpRequest{id: _requestId, client: client}

	httpClientSym.DoHttpRequest = func(clientId uint32, requestId uint32) (error errno.Error) {
		if clientId != uint32(client) {
			return 1
		}
		if requestId != _requestId {
			return 1
		}
		return 0
	}

	resp, err := request.Do()
	if err != nil {
		t.Error(err)
		return
	}
	if uint32(resp.request.id) != _requestId {
		t.Errorf("Response request id `%d` does not match set `%d`", uint32(resp.request.id), _requestId)
		return
	}

	httpClientSym.DoHttpRequest = func(clientId uint32, requestId uint32) (error errno.Error) {
		return 1
	}
	_, err = request.Do()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
