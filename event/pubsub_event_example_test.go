package event_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/pubsub"
	"github.com/taubyte/go-sdk/event"
)

// Taubyte example function that gets a Pub-Sub call
func pubSubExample(e event.Event) uint32 {
	// Confirm that this is a p2p event
	pubSubEvent, err := e.PubSub()
	if err != nil {
		return 1
	}

	// The data sent to the pub-sub function
	data, err := pubSubEvent.Data()
	if err != nil {
		return 1
	}

	fmt.Println("Data:", string(data))

	// The Channel that this function was called through
	channel, err := pubSubEvent.Channel()
	if err != nil {
		return 1
	}
	fmt.Println("Channel:", channel.Name())

	return 0
}

func ExampleEvent_PubSub() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		Channel:   "someChannel",
		EventData: []byte("Hello, world!"),
	}.Mock()

	if pubSubExample(0) != 0 {
		return
	}
	// Output: Data: Hello, world!
	// Channel: someChannel
}
