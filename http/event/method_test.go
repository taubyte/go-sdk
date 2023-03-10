package event

import (
	"testing"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestHttpMethod(t *testing.T) {
	httpEventSym.MockData{Method: "POST", EventId: 1}.Mock()

	var e Event
	_, err := e.Method()
	if err == nil {
		t.Error("Expected error")
		return
	}

	httpEventSym.GetHttpEventMethod = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return 1
	}

	e = 1
	_, err = e.Method()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
