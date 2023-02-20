package event

import (
	"testing"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
)

func TestHttpWrite(t *testing.T) {
	httpEventSym.MockData{EventId: 1}.Mock()

	var e Event
	toWrite := []byte("Hello, world")
	_, err := e.Write(toWrite)
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Empty
	e = 1
	_, err = e.Write(nil)
	if err != nil {
		t.Error(err)
		return
	}
}
