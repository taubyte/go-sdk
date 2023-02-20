package event_test

import (
	"fmt"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/common"
	"github.com/taubyte/go-sdk/event"
)

func ExampleEvent_Return() {
	// Mocking the calls to the vm for usage in tests and playground
	m := httpEventSym.MockData{
		EventType: common.EventTypeHttp,
	}.Mock()

	var e event.Event

	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	err = httpEvent.Return(404)
	if err != nil {
		return
	}

	fmt.Println("Code:", m.ReturnCode)
	// Output: Code: 404
}

func ExampleEvent_Write() {
	// Mocking the calls to the vm for usage in tests and playground
	m := httpEventSym.MockData{
		EventType: common.EventTypeHttp,
	}.Mock()

	var e event.Event

	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	toWrite := []byte("Hello, world!")

	n, err := httpEvent.Write(toWrite)
	if err != nil {
		return
	}
	if len(toWrite) != n {
		return
	}

	fmt.Println("ReturnBody:", string(m.ReturnBody))
	// Output: ReturnBody: Hello, world!
}

func ExampleEvent_Host() {
	// Mocking the calls to the vm for usage in tests and playground
	httpEventSym.MockData{
		EventType: common.EventTypeHttp,
		Host:      "hal.computers.com",
	}.Mock()

	var e event.Event

	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	host, err := httpEvent.Host()
	if err != nil {
		return
	}

	fmt.Println("Host:", host)
	// Output: Host: hal.computers.com
}

func ExampleEvent_Method() {
	// Mocking the calls to the vm for usage in tests and playground
	httpEventSym.MockData{
		EventType: common.EventTypeHttp,
		Method:    "POST",
	}.Mock()

	var e event.Event

	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	method, err := httpEvent.Method()
	if err != nil {
		return
	}

	fmt.Println("Method:", method)
	// Output: Method: POST
}

func ExampleEvent_Path() {
	// Mocking the calls to the vm for usage in tests and playground
	httpEventSym.MockData{
		EventType: common.EventTypeHttp,
		Path:      "/test/v1",
	}.Mock()

	var e event.Event

	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	path, err := httpEvent.Path()
	if err != nil {
		return
	}

	fmt.Println("Path:", path)
	// Output: Path: /test/v1
}

func ExampleEvent_UserAgent() {
	// Mocking the calls to the vm for usage in tests and playground
	httpEventSym.MockData{
		EventType: common.EventTypeHttp,
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/517.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/5447.36",
	}.Mock()

	var e event.Event

	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	path, err := httpEvent.UserAgent()
	if err != nil {
		return
	}

	fmt.Println("UserAgent:", path)
	// Output: UserAgent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/517.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/5447.36
}
