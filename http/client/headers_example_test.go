package client_test

import (
	"fmt"
	"strings"

	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/http/client"
	"github.com/taubyte/go-sdk/utils/slices"
)

func ExampleHttpRequestHeaders_Set() {
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

	testHeaders := map[string][]string{
		"content-type": {"json", "yaml"},
		"fruits":       {"apple", "orange"},
	}
	for key, values := range testHeaders {
		err = request.Headers().Set(key, values...)
		if err != nil {
			return
		}
	}

	for key, values := range testHeaders {
		for _, v := range values {
			if !slices.Contains(m.Headers[key], v) {
				fmt.Printf("Expected %s to contain %s", strings.Join(testHeaders[key], ", \n"), v)
				return
			}
		}
	}

	err = request.Headers().Set("fruits", "banana")
	if err != nil {
		return
	}

	fruits := m.Headers["fruits"]
	if len(fruits) != 1 || fruits[0] != "banana" {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleHttpRequestHeaders_Add() {
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

	err = request.Headers().Add("fruits", "apple")
	if err != nil {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleHttpRequestHeaders_Get() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{}.Mock()

	httpClient, err := client.New()
	if err != nil {
		return
	}

	request, err := httpClient.Request("google.com", client.Headers(map[string][]string{
		"fruits": {"banana", "orange"},
	}))
	if err != nil {
		return
	}

	headers, err := request.Headers().Get("fruits")
	if err != nil {
		return
	}

	if !slices.Contains(headers, "banana") || !slices.Contains(headers, "orange") {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleHttpRequestHeaders_List() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{}.Mock()

	httpClient, err := client.New()
	if err != nil {
		return
	}

	request, err := httpClient.Request("google.com", client.Headers(map[string][]string{
		"fruits":       {"banana"},
		"Content-Type": {"application/json"},
	}))
	if err != nil {
		return
	}

	headerKeys, err := request.Headers().List()
	if err != nil {
		return
	}

	if !slices.Contains(headerKeys, "fruits") || !slices.Contains(headerKeys, "Content-Type") {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleHttpRequestHeaders_GetAll() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{}.Mock()

	httpClient, err := client.New()
	if err != nil {
		return
	}

	request, err := httpClient.Request("google.com", client.Headers(map[string][]string{
		"fruits":       {"banana", "orange"},
		"Content-Type": {"application/json"},
	}))
	if err != nil {
		return
	}

	headers, err := request.Headers().GetAll()
	if err != nil {
		return
	}

	fruits := headers["fruits"]
	if !slices.Contains(fruits, "banana") || !slices.Contains(fruits, "orange") {
		return
	}

	if !slices.Contains(headers["Content-Type"], "application/json") {
		return
	}

	fmt.Println("Success")
	// Output: Success
}
