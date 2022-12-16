package node_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/p2p/node"
	"github.com/taubyte/go-sdk/event"
	"github.com/taubyte/go-sdk/p2p/node"
)

func ExampleCommand_Send() {
	// Mocking the calls to the vm for usage in tests and playground
	m := symbols.MockData{
		Protocol:     "/test/v1",
		Command:      "someCommand",
		SendResponse: []byte("Hello from the other side!"),
	}.Mock()

	service := node.New(m.Protocol)

	// Instantiate a command `someCommand` to protocol `/test/v1`
	command, err := service.Command(m.Command)
	if err != nil {
		return
	}

	// Send the command with data []byte("Hello, world!")
	response, err := command.Send([]byte("Hello, world!"))
	if err != nil {
		return
	}

	// A function representing the call executed by the above command
	_ = func(e event.Event) uint32 {
		p2pEvent, err := e.P2P()
		if err != nil {
			return 1
		}

		data, err := p2pEvent.Data()
		if err != nil {
			return 1
		}

		if string(data) == "Hello, world!" {
			p2pEvent.Write([]byte("Hello from the other side!"))
			return 0
		}
		return 1
	}

	fmt.Println(string(response))
	// Output: Hello from the other side!
}
