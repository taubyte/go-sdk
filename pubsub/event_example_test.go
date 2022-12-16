package pubsub_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/pubsub"
	"github.com/taubyte/go-sdk/event"
)

func ExamplePubSubEvent_Data() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		EventData: []byte("Hello, world!"),
	}.Mock()

	// An event that would be received by a taubyte function
	var e event.Event

	pubSubEvent, err := e.PubSub()
	if err != nil {
		return
	}

	data, err := pubSubEvent.Data()
	if err != nil {
		return
	}

	fmt.Println("Data:", string(data))

	// Output: Data: Hello, world!
}

func ExamplePubSubEvent_Channel() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		Channel: "someChannel",
	}.Mock()

	// An event that would be received by a taubyte function
	var e event.Event

	pubSubEvent, err := e.PubSub()
	if err != nil {
		return
	}

	channel, err := pubSubEvent.Channel()
	if err != nil {
		return
	}

	fmt.Println("Channel:", channel.Name())

	// Output: Channel: someChannel
}
