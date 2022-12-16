package event

import (
	"strings"
	"testing"

	eventSym "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/errno"
)

func checkZeroSize(t *testing.T, err error) {
	if strings.Contains(err.Error(), errno.ErrorZeroSize.String()) == false {
		t.Error("Expected ZeroSize error")
	}
}

func checkNotZeroSize(t *testing.T, err error) {
	if strings.Contains(err.Error(), errno.ErrorZeroSize.String()) == true {
		t.Error("Got ZeroSize error expected a normal error")
	}
}

func TestHttpZero(t *testing.T) {
	var e HttpEvent

	eventSym.GetHttpEventMethodSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return 0
	}
	_, err := e.Method()
	checkZeroSize(t, err)

	eventSym.GetHttpEventHostSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return 0
	}
	_, err = e.Host()
	checkZeroSize(t, err)

	eventSym.GetHttpEventPathSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return 0
	}
	_, err = e.Path()
	checkZeroSize(t, err)

	eventSym.GetHttpEventUserAgentSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return 0
	}
	_, err = e.UserAgent()
	checkZeroSize(t, err)

	eventSym.GetHttpEventMethodSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return errno.ErrorCap
	}
	_, err = e.Method()
	checkNotZeroSize(t, err)

	eventSym.GetHttpEventHostSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return errno.ErrorCap
	}
	_, err = e.Host()
	checkNotZeroSize(t, err)

	eventSym.GetHttpEventPathSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return errno.ErrorCap
	}
	_, err = e.Path()
	checkNotZeroSize(t, err)

	eventSym.GetHttpEventUserAgentSize = func(eventId uint32, size *uint32) (error errno.Error) {
		return errno.ErrorCap
	}
	_, err = e.UserAgent()
	checkNotZeroSize(t, err)
}
