package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestHttpMethod(t *testing.T) {
	symbols.MockData{Method: "POST", EventId: 1}.Mock()

	var e HttpEvent
	_, err := e.Method()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.GetHttpEventMethod = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return 1
	}

	e = 1
	_, err = e.Method()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
