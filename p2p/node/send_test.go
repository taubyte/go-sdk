package node

import (
	"testing"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/p2p/node"
	"github.com/taubyte/go-sdk/errno"
	"gotest.tools/v3/assert"
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

func TestSendTo(t *testing.T) {
	cid1, err := cid.Parse("bafzaajaiaejcatsa2r73dij2iewq47p2c6runxmvht2evx6agmojnk3pjflfcn52")
	if err != nil {
		panic(err)
	}

	m := symbols.MockData{CommandId: 1,
		Protocol:     "/test/v1",
		Command:      "someCommand",
		SendResponse: nil,
		SendTo:       cid1,
	}.Mock()

	// Invalid cid
	{
		c := Command(m.CommandId)
		_, err = c.SendTo([]byte("Hello, world!"), cid.Cid{})
		assert.Assert(t, err != nil)
	}

	// SendCommandTo error
	{
		c := Command(0)
		_, err = c.SendTo([]byte("Hello, world!"), cid1)
		assert.Assert(t, err != nil)
	}

	// Nil response
	{
		c := Command(m.CommandId)
		response, err := c.SendTo([]byte("Hello, world!"), cid1)
		assert.NilError(t, err)
		assert.Assert(t, response == nil)
	}

	// ReadResponse error
	{
		m.SendResponse = []byte("Hello from the other side!")
		m.Mock()

		symbols.ReadCommandResponse = func(id uint32, data *byte, dataSize uint32) (error errno.Error) {
			return 1
		}

		c := Command(m.CommandId)
		_, err = c.SendTo([]byte("Hello, world!"), cid1)
		assert.Assert(t, err != nil)
	}
}
