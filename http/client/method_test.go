package client

import (
	"testing"

	httpClientSym "github.com/taubyte/go-sdk-symbols/http/client"
	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestMethodOption(t *testing.T) {
	symbols.MockData{ClientId: 1}.Mock()

	httpClientSym.SetHttpRequestMethod = func(clientId uint32, requestId, method uint32) (error errno.Error) {
		return 1
	}

	_, err := HttpClient(1).Request("", Method("POST"))
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestMethodSet(t *testing.T) {
	symbols.MockData{ClientId: 1, RequestId: 1}.Mock()

	r := &HttpRequest{id: 1, client: 1}
	err := r.Method().Set("POST")
	if err != nil {
		t.Error(err)
		return
	}

	r = &HttpRequest{id: 0, client: 1}
	err = r.Method().Set("POST")
	if err == nil {
		t.Error("Expected error")
		return
	}

	r = &HttpRequest{id: 1, client: 0}
	err = r.Method().Set("POST")
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Unknown
	r = &HttpRequest{id: 1, client: 1}
	err = r.Method().Set("SomethingElse")
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestMethodGet(t *testing.T) {
	m := symbols.MockData{ClientId: 1, RequestId: 1}.Mock()
	*m.RequestMethod = "PUT"

	r := &HttpRequest{id: 1, client: 1}
	method, err := r.Method().Get()
	if err != nil {
		t.Error(err)
		return
	}
	if method != *m.RequestMethod {
		t.Errorf("Expected %s, got %s", *m.RequestMethod, method)
		return
	}

	r = &HttpRequest{id: 0, client: 1}
	_, err = r.Method().Get()
	if err == nil {
		t.Error("Expected error")
		return
	}

	r = &HttpRequest{id: 1, client: 0}
	_, err = r.Method().Get()
	if err == nil {
		t.Error("Expected error")
		return
	}

	r = &HttpRequest{id: 1, client: 1}
	*m.RequestMethod = "SomethingElse"
	method, err = r.Method().Get()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
