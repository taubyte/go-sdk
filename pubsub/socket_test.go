package pubsub

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/pubsub"
	"github.com/taubyte/go-sdk/errno"
)

func TestSocket(t *testing.T) {
	m := symbols.MockData{
		WebSocketURL: "/some/url",
	}.Mock()

	channel := &ChannelObject{}

	// No name
	_, err := channel.WebSocket().Url()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Size error
	channel.name = "someChannel"
	_, err = channel.WebSocket().Url()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Zero size
	m.Channel = "someChannel"
	m.Mock()

	symbols.GetWebSocketURLSize = func(channel string, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 0
		return 0
	}

	channel.name = "someChannel"
	_, err = channel.WebSocket().Url()
	if err != nil {
		t.Error(err)
		return
	}

	// Data error
	m.Mock()

	symbols.GetWebSocketURL = func(channel string, socketURLPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = channel.WebSocket().Url()
	if err == nil {
		t.Error(err)
		return
	}
}
