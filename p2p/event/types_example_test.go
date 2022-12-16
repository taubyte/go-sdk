package event_test

import (
	"fmt"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/p2p/event"

	"github.com/taubyte/go-sdk/p2p/event"
)

func ExampleEvent_Command() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		Command: "someCommand",
	}.Mock()

	var p2pEvent event.Event // Event for a taubyte p2p function called by command, `someCommand`

	command, err := p2pEvent.Command()
	if err != nil {
		return
	}

	fmt.Println(command)
	// Output: someCommand
}

func ExampleEvent_Data() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		Data: []byte("Hello, world!"),
	}.Mock()

	var p2pEvent event.Event // Event for a taubyte p2p function called with data []byte("Hello, world!")

	data, err := p2pEvent.Data()
	if err != nil {
		return
	}

	fmt.Println(string(data))
	// Output: Hello, world!
}

func ExampleEvent_From() {
	// Mocking the calls to the vm for usage in tests and playground
	_cid, err := cid.Parse("QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3")
	if err != nil {
		return
	}
	symbols.MockData{
		From: _cid,
	}.Mock()

	var p2pEvent event.Event // Event for a taubyte p2p function sent by a node with id: `QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3`

	cid, err := p2pEvent.From()
	if err != nil {
		return
	}

	fmt.Println(cid.String())
	// Output: QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3
}

func ExampleEvent_To() {
	// Mocking the calls to the vm for usage in tests and playground
	_cid, err := cid.Parse("QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3")
	if err != nil {
		return
	}
	symbols.MockData{
		To: _cid,
	}.Mock()

	var p2pEvent event.Event // Event for a taubyte p2p function sent by a node with id: `QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3`

	cid, err := p2pEvent.To()
	if err != nil {
		return
	}

	fmt.Println(cid.String())
	// Output: QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3
}

func ExampleEvent_Protocol() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		Protocol: "/test/v1",
	}.Mock()

	var p2pEvent event.Event // Event for a taubyte p2p function sent on protocol `/test/v1`

	protocol, err := p2pEvent.Protocol()
	if err != nil {
		return
	}

	fmt.Println(protocol)
	// Output: /test/v1
}

func ExampleEvent_Write() {
	// Mocking the calls to the vm for usage in tests and playground
	m := symbols.MockData{}.Mock()

	var p2pEvent event.Event // Event for a taubyte p2p function sent on protocol `/test/v1`

	err := p2pEvent.Write([]byte("Hello, world!"))
	if err != nil {
		return
	}

	fmt.Println(string(m.DataToSend))
	// Output: Hello, world!
}
