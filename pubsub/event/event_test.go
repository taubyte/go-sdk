package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/pubsub/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestEventData(t *testing.T) {
	m := symbols.MockData{EventId: 1}.Mock()

	var e Event

	// Wrong event
	_, err := e.Data()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Zero size
	e = 1
	data, err := e.Data()
	if err != nil {
		t.Error(err)
		return
	}
	if data != nil {
		t.Error("Expected nil data")
		return
	}

	// Data error
	m.EventData = []byte("Hello, world")
	m.Mock()

	symbols.GetMessageData = func(eventId uint32, buf *byte) (error errno.Error) {
		return 1
	}

	_, err = e.Data()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestEventChannel(t *testing.T) {
	m := symbols.MockData{EventId: 1}.Mock()

	var e Event

	// Wrong event
	_, err := e.Channel()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Zero size
	e = 1
	_, err = e.Channel()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Data error
	m.Channel = "someChannel"
	m.Mock()

	symbols.GetMessageChannel = func(eventId uint32, buf *byte) (error errno.Error) {
		return 1
	}

	_, err = e.Channel()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
