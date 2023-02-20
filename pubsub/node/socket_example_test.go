package node_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/pubsub/node"
	pubsub "github.com/taubyte/go-sdk/pubsub/node"
)

func ExampleChannelObject_WebSocket() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		Channel:      "someChannel",
		WebSocketURL: "ws-QmQq71tkq1yKmYobFUhPWF2MejK5CrvpY4h7HQDDT8b8Zb/someChannel",
	}.Mock()

	channel, err := pubsub.Channel("someChannel")
	if err != nil {
		return
	}

	url, err := channel.WebSocket().Url()
	if err != nil {
		return
	}

	fmt.Println("URL:", url.Path)
	// Output: URL: ws-QmQq71tkq1yKmYobFUhPWF2MejK5CrvpY4h7HQDDT8b8Zb/someChannel
}
