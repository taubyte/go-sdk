package pubsub_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/pubsub"
	"github.com/taubyte/go-sdk/pubsub"
)

func ExampleChannelObject_Publish() {
	// Mocking the calls to the vm for usage in tests and playground
	m := symbols.MockData{
		Channel: "someChannel",
	}.Mock()

	channel, err := pubsub.Channel("someChannel")
	if err != nil {
		return
	}

	err = channel.Publish([]byte("Hello, world!"))
	if err != nil {
		return
	}

	fmt.Println("Published:", string(m.PublishedData))
	// Output: Published: Hello, world!
}
