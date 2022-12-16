package client_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/http/client"
)

func ExampleHttpRequest_Body() {
	// Mocking the calls to the vm for usage in tests and playground
	m := symbols.MockData{}.Mock()

	httpClient, err := client.New()
	if err != nil {
		return
	}

	request, err := httpClient.Request("google.com")
	if err != nil {
		return
	}

	err = request.Body().Set([]byte("Hello, world!"))
	if err != nil {
		return
	}

	fmt.Println(string(m.RequestBody))
	// Output: Hello, world!
}
