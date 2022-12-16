package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/common"
)

func TestTypeBasic(t *testing.T) {
	var e Event

	symbols.MockEventType(0, common.EventTypeUndefined)
	if e.Type() != common.EventTypeUndefined {
		t.Error("Expected unknown event type", e.Type())
		return
	}

	symbols.MockEventType(0, common.EventTypeHttp)
	_, err := e.HTTP()
	if e.Type() != common.EventTypeHttp || err != nil {
		t.Error("Expected http event type got", e.Type())
		return
	}

	symbols.MockEventType(0, common.EventTypePubsub)
	_, err = e.PubSub()
	if e.Type() != common.EventTypePubsub || err != nil {
		t.Error("Expected pubsub event type", e.Type())
		return
	}

	symbols.MockEventType(0, common.EventTypeP2P)
	_, err = e.P2P()
	if e.Type() != common.EventTypeP2P || err != nil {
		t.Error("Expected pubsub event type", e.Type())
		return
	}
}

func TestTypeError(t *testing.T) {
	var e Event

	symbols.MockEventType(1, common.EventTypeHttp)
	_, err := e.HTTP()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockEventType(1, common.EventTypeP2P)
	_, err = e.P2P()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockEventType(1, common.EventTypePubsub)
	_, err = e.PubSub()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
