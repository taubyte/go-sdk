package event

import (
	"errors"

	eventSym "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/common"
	p2p "github.com/taubyte/go-sdk/p2p/event"
	"github.com/taubyte/go-sdk/pubsub"
)

func (e Event) Type() common.EventType {
	var et uint32
	eventSym.GetEventType(uint32(e), &et)

	return common.EventType(et)
}

func (e Event) HTTP() (HttpEvent, error) {
	if e.Type() != common.EventTypeHttp {
		return 0, errors.New("Not an http event")
	}

	return HttpEvent(e), nil
}

func (e Event) PubSub() (pubsub.PubSubEvent, error) {
	if e.Type() != common.EventTypePubsub {
		return 0, errors.New("Not a pubsub event")
	}

	return pubsub.PubSubEvent(e), nil
}

// Gives a p2pEvent if the current function is called through p2p
func (e Event) P2P() (p2p.Event, error) {
	if e.Type() != common.EventTypeP2P {
		return 0, errors.New("Not a p2p event")
	}

	return p2p.Event(e), nil
}
