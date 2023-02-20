package event

import (
	"testing"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestHttpHost(t *testing.T) {
	httpEventSym.MockData{Host: "hal.computers.com", EventId: 1}.Mock()

	var e Event
	_, err := e.Host()
	if err == nil {
		t.Error("Expected error")
		return
	}

	httpEventSym.GetHttpEventHost = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return 1
	}

	e = 1
	_, err = e.Host()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
