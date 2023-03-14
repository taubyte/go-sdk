package node

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/p2p/node"
)

func TestCommand(t *testing.T) {
	m := symbols.MockData{Protocol: "/test/v1", Command: "someCommand"}.Mock()

	service := &Service{}
	_, err := service.Command(m.Command)
	if err == nil {
		t.Error("Expected error")
		return
	}
}
