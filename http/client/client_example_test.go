package client_test

import (
	"fmt"
	"io/ioutil"

	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/http/client"
)

func ExampleNew() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{ClientId: 5}.Mock()

	httpClient, err := client.New()
	if err != nil {
		return
	}

	fmt.Printf("%d\n", httpClient)
	// Output: 5
}

func ExampleHttpClient_Request() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{ClientId: 5, RequestId: 3}.Mock()

	httpClient, err := client.New()
	if err != nil {
		return
	}

	request, err := httpClient.Request(
		"www.google.com",
		client.Headers(map[string][]string{"Content-Type": {"application/json"}}),
		client.Method("GET"),
		client.Body([]byte("Hello, world!")),
	)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}

	fmt.Printf("%d\n", request)
	// Output: {3 5}
}

func ExampleHttpClient() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{ClientId: 5, RequestId: 3, ResponseBody: []byte(`{"hello": "world"}`)}.Mock()

	httpClient, err := client.New()
	if err != nil {
		return
	}

	request, err := httpClient.Request("google.com")
	if err != nil {
		return
	}

	if request.Headers().Set("Content-Type", "application/json") != nil {
		return
	}

	response, err := request.Do()
	if err != nil {
		return
	}

	body := response.Body()
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	fmt.Println(string(data))
	// Output: {"hello": "world"}
}
