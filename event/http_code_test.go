package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/event"
)

func TestHttpCode(t *testing.T) {
	m := &symbols.MockData{EventId: 1}
	m.MockReturnCode()

	var e HttpEvent
	err := e.Return(503)
	if err == nil {
		t.Error("expected and error")
		return
	}
}
