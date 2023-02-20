package event

import (
	"errors"

	eventSym "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/common"
	http "github.com/taubyte/go-sdk/http/event"
	p2p "github.com/taubyte/go-sdk/p2p/event"
	pubsub "github.com/taubyte/go-sdk/pubsub/event"
)

func (e Event) Type() common.EventType {
	var et uint32
	eventSym.GetEventType(uint32(e), &et)

	return common.EventType(et)
}

func (e Event) HTTP() (http.Event, error) {
	if e.Type() != common.EventTypeHttp {
		return 0, errors.New("Not an http event")
	}

	return http.Event(e), nil
}

func (e Event) PubSub() (pubsub.Event, error) {
	if e.Type() != common.EventTypePubsub {
		return 0, errors.New("Not a pubsub event")
	}

	return pubsub.Event(e), nil
}

// Gives a p2pEvent if the current function is called through p2p
func (e Event) P2P() (p2p.Event, error) {
	if e.Type() != common.EventTypeP2P {
		return 0, errors.New("Not a p2p event")
	}

	return p2p.Event(e), nil
}
