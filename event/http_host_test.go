package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestHttpHost(t *testing.T) {
	symbols.MockData{Host: "hal.computers.com", EventId: 1}.Mock()

	var e HttpEvent
	_, err := e.Host()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.GetHttpEventHost = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return 1
	}

	e = 1
	_, err = e.Host()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
