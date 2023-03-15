package node

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/pubsub/node"
)

func TestChannel(t *testing.T) {
	symbols.MockData{Channel: "someChannel"}.Mock()

	_, err := Channel("")
	if err == nil {
		t.Error("Expected error")
		return
	}
}
