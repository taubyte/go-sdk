package node

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/pubsub/node"
)

func TestSubscribe(t *testing.T) {
	symbols.MockData{}.Mock()

	channel := &ChannelObject{name: "someChannel"}

	err := channel.Subscribe()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
