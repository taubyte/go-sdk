package common

import (
	"testing"
)

func TestEventType(t *testing.T) {
	e := EventType(0)
	if e.String() != "EventTypeUndefined" {
		t.Error("Wrong type")
	}

	e = EventType(1)
	if e.String() != "EventTypeHttp" {
		t.Error("Wrong type")
	}

	e = EventType(2)
	if e.String() != "EventTypePubsub" {
		t.Error("Wrong type")
	}

	e = EventType(3)
	if e.String() != "EventTypeP2P" {
		t.Error("Wrong type")
	}

	e = EventType(4)
	if e.String() != "EventTypeUndefined" {
		t.Error("Wrong type")
	}
}
