package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestHttpUserAgent(t *testing.T) {
	symbols.MockData{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/517.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/5447.36",
		EventId:   1,
	}.Mock()

	var e HttpEvent
	_, err := e.UserAgent()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.GetHttpEventUserAgent = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return 1
	}

	e = 1
	_, err = e.UserAgent()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
