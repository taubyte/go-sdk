package event

import (
	"testing"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestHttpUserAgent(t *testing.T) {
	httpEventSym.MockData{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/517.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/5447.36",
		EventId:   1,
	}.Mock()

	var e Event
	_, err := e.UserAgent()
	if err == nil {
		t.Error("Expected error")
		return
	}

	httpEventSym.GetHttpEventUserAgent = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return 1
	}

	e = 1
	_, err = e.UserAgent()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
