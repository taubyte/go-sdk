package common

type EventType uint32

const (
	EventTypeUndefined EventType = iota
	EventTypeHttp
	EventTypePubsub
	EventTypeP2P
)

func (et EventType) String() string {
	switch et {
	case EventTypeUndefined:
		return "EventTypeUndefined"
	case EventTypeHttp:
		return "EventTypeHttp"
	case EventTypePubsub:
		return "EventTypePubsub"
	case EventTypeP2P:
		return "EventTypeP2P"
	default:
		return "EventTypeUndefined"
	}
}
