package event

import (
	"testing"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestHttpPath(t *testing.T) {
	httpEventSym.MockData{Path: "/test/v1", EventId: 1}.Mock()

	var e Event
	_, err := e.Path()
	if err == nil {
		t.Error("Expected error")
		return
	}

	httpEventSym.GetHttpEventPath = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return 1
	}

	e = 1
	_, err = e.Path()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
