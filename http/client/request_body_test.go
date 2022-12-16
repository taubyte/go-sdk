package client

import (
	"testing"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestBodyOption(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()

	httpClientSym.SetHttpRequestBody = func(clientId, requestId uint32, data *byte, dataSize uint32) (error errno.Error) {
		return 1
	}

	_, err := HttpClient(1).Request("", Body([]byte("Hello, world")))
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestBodyError(t *testing.T) {
	m := symbols.MockData{ClientId: 1, RequestId: 1}.Mock()

	testSet := "Hello, world!"
	r := &HttpRequest{id: 1, client: 1}
	err := r.Body().Set([]byte(testSet))
	if err != nil {
		t.Error(err)
		return
	}

	if string(m.RequestBody) != testSet {
		t.Errorf("Expected %s, got %s", testSet, string(m.RequestBody))
		return
	}

	r = &HttpRequest{id: 0, client: 1}
	err = r.Body().Set([]byte(testSet))
	if err == nil {
		t.Error("Expected error")
		return
	}

	r = &HttpRequest{id: 1, client: 0}
	err = r.Body().Set([]byte(testSet))
	if err == nil {
		t.Error("Expected error")
		return
	}
}
