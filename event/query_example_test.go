package event_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/common"
	"github.com/taubyte/go-sdk/event"
)

func ExampleHttpEvent_Query() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		EventType: common.EventTypeHttp,
		Queries: map[string]string{
			"username": "taubyte",
		},
	}.Mock()

	var e event.Event
	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	query := httpEvent.Query()

	username, err := query.Get("username")
	if err != nil {
		return
	}
	if username != "taubyte" {
		return
	}

	queryList, err := query.List()
	if err != nil {
		return
	}
	fmt.Println("Queries:", queryList)

	// Output: Queries: [username]
}
