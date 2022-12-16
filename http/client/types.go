package client

type HttpClient uint32
type HttpRequest struct {
	id     uint32
	client HttpClient
}

type HttpRequestMethod HttpRequest

type HttpResponse struct {
	request *HttpRequest
}

type HttpResponseBody HttpResponse

type HttpRequestOption func(HttpRequest) error

type HttpMethod uint32

const (
	UNKNOWN HttpMethod = iota
	GET
	POST
	PUT
	DELETE
	HEAD
	OPTIONS
	PATCH
	TRACE
	CONNECT
)

type HttpRequestHeaders HttpRequest
type HttpRequestBody HttpRequest
