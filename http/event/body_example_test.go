package event_test

import (
	"fmt"
	"io/ioutil"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/common"
	"github.com/taubyte/go-sdk/event"
)

func ExampleEvent_Body() {
	// Mocking the calls to the vm for usage in tests and playground
	httpEventSym.MockData{
		EventType: common.EventTypeHttp,
		Body:      []byte("Hello, world!"),
	}.Mock()

	var e event.Event
	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	body := httpEvent.Body()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	err = body.Close()
	if err != nil {
		return
	}

	fmt.Println(string(data))
	// Output: Hello, world!
}
