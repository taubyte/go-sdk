package event_test

import (
	"fmt"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/p2p/event"
	"github.com/taubyte/go-sdk/event"
)

// Taubyte example function that gets a p2p call
func p2pExample(e event.Event) uint32 {
	// Confirm that this is a p2p event
	p2pEvent, err := e.P2P()
	if err != nil {
		return 1
	}

	// The data that was sent over p2p to this function
	data, err := p2pEvent.Data()
	if err != nil {
		return 1
	}

	fmt.Println("Data:", string(data))

	// The protocol that was called to reach this function
	protocol, err := p2pEvent.Protocol()
	if err != nil {
		return 1
	}

	fmt.Println("Protocol:", protocol)

	// The Command that was called to reach this function
	command, err := p2pEvent.Command()
	if err != nil {
		return 1
	}
	fmt.Println("Command:", command)

	// The cid of the node that sent the p2p request
	from, err := p2pEvent.From()
	if err != nil {
		return 1
	}

	fmt.Println("From:", from.String())

	// The cid of the receiving node, or the node that this function is running on
	to, err := p2pEvent.To()
	if err != nil {
		return 1
	}

	fmt.Println("To:", to.String())

	// Write the response
	err = p2pEvent.Write([]byte("Hello from the other side!"))
	if err != nil {
		return 1
	}

	return 0
}

func ExampleEvent_P2P() {
	// Mocking the calls to the vm for usage in tests and playground
	fromCid, _ := cid.Parse("QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3")
	toCid, _ := cid.Parse("QmTLqapULGiA5ZFWk2ckEjKq2B6kGVuFUF1frSAeeGjuGt")
	m := symbols.MockData{
		From:     fromCid,
		To:       toCid,
		Command:  "ping",
		Protocol: "/test/v1",
		Data:     []byte("Hello, world!"),
	}.Mock()

	if p2pExample(0) != 0 {
		return
	}

	fmt.Println("Sent:", string(m.DataToSend))
	// Output: Data: Hello, world!
	// Protocol: /test/v1
	// Command: ping
	// From: QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3
	// To: QmTLqapULGiA5ZFWk2ckEjKq2B6kGVuFUF1frSAeeGjuGt
	// Sent: Hello from the other side!
}
