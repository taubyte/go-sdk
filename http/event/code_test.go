package event

import (
	"testing"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
)

func TestHttpCode(t *testing.T) {
	m := &httpEventSym.MockData{EventId: 1}
	m.MockReturnCode()

	var e Event
	err := e.Return(503)
	if err == nil {
		t.Error("expected and error")
		return
	}
}
