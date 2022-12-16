package node

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/p2p/node"
	"github.com/taubyte/go-sdk/errno"
)

func TestListen(t *testing.T) {
	m := symbols.MockData{ListenProtocol: "/test/v1", Command: "someCommand"}.Mock()

	service := &Service{}
	_, err := service.Listen()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.ListenToProtocolSize = func(protocol string, responseSize *uint32) (error errno.Error) {
		*responseSize = 1
		return 1
	}

	service.protocol = m.ListenProtocol
	_, err = service.Listen()
	if err == nil {
		t.Error("Expected error")
		return
	}

	m.Mock()
	symbols.ListenToProtocol = func(protocol string, response *byte, responseSize uint32) (error errno.Error) {
		return 1
	}

	_, err = service.Listen()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
