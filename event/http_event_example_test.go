package event_test

import (
	"fmt"
	"io/ioutil"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/common"
	"github.com/taubyte/go-sdk/event"
)

// Taubyte example function that gets an http call
func httpExample(e event.Event) uint32 {
	// Confirm that this is a http event
	httpEvent, err := e.HTTP()
	if err != nil {
		return 1
	}

	body := httpEvent.Body()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return 1
	}

	fmt.Println("Body:", string(data))

	err = body.Close()
	if err != nil {
		return 1
	}

	// The host or fqdn used to reach this function
	host, err := httpEvent.Host()
	if err != nil {
		return 1
	}

	fmt.Println("Host:", host)

	// The Method that was called to reach this function
	method, err := httpEvent.Method()
	if err != nil {
		return 1
	}
	fmt.Println("Method:", method)

	// The url path not including the host used to reach this function
	path, err := httpEvent.Path()
	if err != nil {
		return 1
	}

	fmt.Println("Path:", path)

	// The user agent of the http request
	userAgent, err := httpEvent.UserAgent()
	if err != nil {
		return 1
	}

	fmt.Println("UserAgent:", userAgent)

	// The headers of the http request
	headers, err := httpEvent.Headers().List()
	if err != nil {
		return 1
	}

	fmt.Println("Headers:", headers)

	contentType, err := httpEvent.Headers().Get("Content-Type")
	if err != nil {
		return 1
	}

	fmt.Println("ContentType:", contentType)

	// A list of queries sent in the url
	queries, err := httpEvent.Query().List()
	if err != nil {
		return 1
	}

	fmt.Println("Queries:", queries)

	// The username query
	username, err := httpEvent.Query().Get("username")
	if err != nil {
		return 1
	}

	fmt.Println("Username:", username)

	// Write the return code
	err = httpEvent.Return(200)
	if err != nil {
		return 1
	}

	// Write the http response
	_, err = httpEvent.Write([]byte("Hello from the other side!"))
	if err != nil {
		return 1
	}

	return 0
}

func ExampleEvent_HTTP() {
	// Mocking the calls to the vm for usage in tests and playground
	m := httpEventSym.MockData{
		EventType: common.EventTypeHttp,
		Body:      []byte("Hello, world!"),
		Host:      "taubyte.com",
		Method:    "POST",
		Path:      "/console/v1",
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64)",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Queries: map[string]string{
			"username": "taubyte",
		},
	}.Mock()

	if httpExample(0) != 0 {
		return
	}

	fmt.Println("ReturnCode:", m.ReturnCode)
	fmt.Println("Sent:", string(m.ReturnBody))
	// Output: Body: Hello, world!
	// Host: taubyte.com
	// Method: POST
	// Path: /console/v1
	// UserAgent: Mozilla/5.0 (X11; Linux x86_64)
	// Headers: [Content-Type]
	// ContentType: application/json
	// Queries: [username]
	// Username: taubyte
	// ReturnCode: 200
	// Sent: Hello from the other side!
}
