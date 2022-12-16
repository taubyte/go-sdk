package node

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/p2p/node"
	"github.com/taubyte/go-sdk/errno"
)

func TestSend(t *testing.T) {
	m := symbols.MockData{CommandId: 1, Protocol: "/test/v1", Command: "someCommand", SendResponse: nil}.Mock()

	c := Command(0)
	_, err := c.Send([]byte("Hello, world!"))
	if err == nil {
		t.Error(err)
		return
	}

	c = Command(m.CommandId)
	response, err := c.Send([]byte("Hello, world!"))
	if response != nil || err != nil {
		t.Error("Expected all nil")
		return
	}

	m.SendResponse = []byte("Hello from the other side!")
	m.Mock()

	symbols.ReadCommandResponse = func(id uint32, data *byte, dataSize uint32) (error errno.Error) {
		return 1
	}

	c = Command(m.CommandId)
	_, err = c.Send([]byte("Hello, world!"))
	if err == nil {
		t.Error("Expected error")
		return
	}

}
