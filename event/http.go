package event

import (
	"fmt"

	eventSym "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/errno"
)

func (e HttpEvent) Method() (string, error) {
	var size uint32
	err := eventSym.GetHttpEventMethodSize(uint32(e), &size)
	if err != 0 || size == 0 {
		if err == 0 {
			err = errno.ErrorZeroSize
		}
		return "", fmt.Errorf("Getting HTTP method size failed with: %s", err)
	}

	method := make([]byte, size)
	err = eventSym.GetHttpEventMethod(uint32(e), &method[0], uint32(len(method)))
	if err != 0 {
		return "", fmt.Errorf("Getting HTTP method failed with: %s", err)
	}
	return string(method), nil
}

func (e HttpEvent) Write(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	var n uint32
	err := eventSym.EventHttpWrite(uint32(e), &data[0], uint32(len(data)), &n)
	if err != 0 {
		return int(n), fmt.Errorf("Writing HTTP reply failed with: %s", err)
	}
	return int(n), nil
}

func (e HttpEvent) Return(code int) error {
	err := eventSym.EventHttpRetCode(uint32(e), uint32(code))
	if err != 0 {
		return fmt.Errorf("Writing return code failed with: %s", err)
	}
	return nil
}

func (e HttpEvent) Host() (string, error) {
	var size uint32
	err := eventSym.GetHttpEventHostSize(uint32(e), &size)
	if err != 0 || size == 0 {
		if err == 0 {
			err = errno.ErrorZeroSize
		}
		return "", fmt.Errorf("Getting HTTP request host size failed with: %s", err)
	}

	host := make([]byte, size)
	err = eventSym.GetHttpEventHost(uint32(e), &host[0], uint32(len(host)))
	if err != 0 {
		return "", fmt.Errorf("Getting HTTP request host failed with: %s", err)
	}
	return string(host), nil
}

func (e HttpEvent) Path() (string, error) {
	var size uint32
	err := eventSym.GetHttpEventPathSize(uint32(e), &size)
	if err != 0 || size == 0 {
		if err == 0 {
			err = errno.ErrorZeroSize
		}
		return "", fmt.Errorf("Getting HTTP request path size failed with: %s", err)
	}

	path := make([]byte, size)
	err = eventSym.GetHttpEventPath(uint32(e), &path[0], uint32(len(path)))
	if err != 0 {
		return "", fmt.Errorf("Getting HTTP request path failed with: %s", err)
	}
	return string(path), nil
}

func (e HttpEvent) UserAgent() (string, error) {
	var size uint32
	err := eventSym.GetHttpEventUserAgentSize(uint32(e), &size)
	if err != 0 || size == 0 {
		if err == 0 {
			err = errno.ErrorZeroSize
		}
		return "", fmt.Errorf("Getting HTTP request User Agent size failed with: %s", err)
	}

	userAgent := make([]byte, size)
	err = eventSym.GetHttpEventUserAgent(uint32(e), &userAgent[0], uint32(len(userAgent)))
	if err != 0 {
		return "", fmt.Errorf("Getting HTTP request User Agent failed with: %s", err)
	}
	return string(userAgent), nil
}
