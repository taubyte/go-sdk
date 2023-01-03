package event_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/common"
	"github.com/taubyte/go-sdk/event"
)

func ExampleHttpEvent_Redirect() {
	// Mocking the calls to the vm for usage in tests and playground
	m := symbols.MockData{
		EventType: common.EventTypeHttp,
	}.Mock()

	var e event.Event
	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	err = httpEvent.Redirect("https://google.com").Temporary()
	if err != nil {
		return
	}

	fmt.Println(m.RedirectedTo)
	// Output: https://google.com
}
