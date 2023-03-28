package event

import (
	"strings"
	"testing"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/errno"
)

func checkZeroSize(t *testing.T, err error) {
	if !strings.Contains(err.Error(), errno.ErrorZeroSize.String()) {
		t.Error("Expected ZeroSize error")
	}
}

func checkNotZeroSize(t *testing.T, err error) {
	if strings.Contains(err.Error(), errno.ErrorZeroSize.String()) {
		t.Error("Got ZeroSize error expected a normal error")
	}
}

func TestHttpZero(t *testing.T) {
	var e Event

	httpEventSym.GetHttpEventMethodSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return 0
	}
	_, err := e.Method()
	checkZeroSize(t, err)

	httpEventSym.GetHttpEventHostSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return 0
	}
	_, err = e.Host()
	checkZeroSize(t, err)

	httpEventSym.GetHttpEventPathSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return 0
	}
	_, err = e.Path()
	checkZeroSize(t, err)

	httpEventSym.GetHttpEventUserAgentSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return 0
	}
	_, err = e.UserAgent()
	checkZeroSize(t, err)

	httpEventSym.GetHttpEventMethodSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return errno.ErrorCap
	}
	_, err = e.Method()
	checkNotZeroSize(t, err)

	httpEventSym.GetHttpEventHostSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return errno.ErrorCap
	}
	_, err = e.Host()
	checkNotZeroSize(t, err)

	httpEventSym.GetHttpEventPathSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return errno.ErrorCap
	}
	_, err = e.Path()
	checkNotZeroSize(t, err)

	httpEventSym.GetHttpEventUserAgentSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return errno.ErrorCap
	}
	_, err = e.UserAgent()
	checkNotZeroSize(t, err)
}
