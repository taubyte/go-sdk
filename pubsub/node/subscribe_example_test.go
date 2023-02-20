package node_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/pubsub"
	pubsub "github.com/taubyte/go-sdk/pubsub/node"
)

func ExampleChannelObject_Subscribe() {
	// Mocking the calls to the vm for usage in tests and playground
	m := symbols.MockData{
		Channel: "someChannel",
	}.Mock()

	channel, err := pubsub.Channel("someChannel")
	if err != nil {
		return
	}

	err = channel.Subscribe()
	if err != nil {
		return
	}

	fmt.Println("Subs:", m.Subscriptions)
	// Output: Subs: [someChannel]
}
