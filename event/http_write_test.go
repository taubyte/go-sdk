package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/event"
)

func TestHttpWrite(t *testing.T) {
	symbols.MockData{EventId: 1}.Mock()

	var e HttpEvent
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
