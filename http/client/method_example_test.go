package client_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/http/client"
)

func ExampleHttpRequest_Method() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{}.Mock()

	httpClient, err := client.New()
	if err != nil {
		return
	}

	request, err := httpClient.Request("google.com")
	if err != nil {
		return
	}

	err = request.Method().Set("POST")
	if err != nil {
		return
	}

	method, err := request.Method().Get()
	if err != nil {
		return
	}

	fmt.Println(method)
	// Output: POST
}
