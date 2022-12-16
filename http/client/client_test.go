package client

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestClientError(t *testing.T) {
	symbols.NewHttpClient = func(clientId *uint32) (error errno.Error) {
		return 1
	}

	_, err := New()
	if err == nil {
		t.Error("Expected Error")
		return
	}
}

func TestRequestError(t *testing.T) {
	m := symbols.MockData{
		ClientId:   5,
		RequestId:  3,
		RequestUrl: "www.google.com",
	}.Mock()

	_, err := HttpClient(2).Request(m.RequestUrl)
	if err == nil {
		t.Error("Expected Error")
	}

	_, err = HttpClient(m.ClientId).Request(m.RequestUrl + "/hello")
	if err == nil {
		t.Error("Expected Error")
	}
}

func TestRequestOptionError(t *testing.T) {
	m := symbols.MockData{
		ClientId:   5,
		RequestId:  3,
		RequestUrl: "www.google.com",
	}.Mock()

	new := func(opts ...HttpRequestOption) error {
		_, err := HttpClient(m.ClientId).Request(m.RequestUrl, opts...)
		return err
	}

	symbols.SetHttpRequestHeader = func(clientId, requestId uint32, key string, valuesPtr *byte, valuesSize uint32) (error errno.Error) {
		return 1
	}
	if new(Headers(map[string][]string{"hello": {"world"}})) == nil {
		t.Error("Expected error")
		return
	}

	symbols.SetHttpRequestMethod = func(clientId, requestId, method uint32) (error errno.Error) {
		return 1
	}
	if new(Method("GET")) == nil {
		t.Error("Expected error")
		return
	}

	symbols.SetHttpRequestBody = func(clientId, requestId uint32, data *byte, dataSize uint32) (error errno.Error) {
		return 1
	}
	if new(Body([]byte("Hello, world"))) == nil {
		t.Error("Expected error")
		return
	}
}
