package convert

import (
	"fmt"
	"strings"
)

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

func MethodUintToString(_method uint32) (method string, err error) {
	switch HttpMethod(_method) {
	case GET:
		method = "GET"
	case POST:
		method = "POST"
	case PUT:
		method = "PUT"
	case DELETE:
		method = "DELETE"
	case HEAD:
		method = "HEAD"
	case OPTIONS:
		method = "OPTIONS"
	case PATCH:
		method = "PATCH"
	case TRACE:
		method = "TRACE"
	case CONNECT:
		method = "CONNECT"
	default:
		return "", fmt.Errorf("Method `%d` not supported", _method)
	}

	return
}

func MethodStringToUint(_method string) (methodId uint32, err error) {
	var method HttpMethod
	switch strings.ToUpper(_method) {
	case "GET":
		method = GET
	case "POST":
		method = POST
	case "PUT":
		method = PUT
	case "DELETE":
		method = DELETE
	case "HEAD":
		method = HEAD
	case "OPTIONS":
		method = OPTIONS
	case "PATCH":
		method = PATCH
	case "TRACE":
		method = TRACE
	case "CONNECT":
		method = CONNECT
	default:
		return 0, fmt.Errorf("Method %s not supported", _method)
	}

	return uint32(method), nil
}
