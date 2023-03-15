package client_test

import (
	"fmt"
	"io/ioutil"

	symbols "github.com/taubyte/go-sdk-symbols/http/client"
	"github.com/taubyte/go-sdk/http/client"
)

func ExampleHttpRequest_Do() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{ResponseBody: []byte("Hello, world!"), ResponseHeaders: map[string][]string{"Content-Type": {"application/json"}}}.Mock()

	httpClient, err := client.New()
	if err != nil {
		return
	}

	request, err := httpClient.Request("google.com")
	if err != nil {
		return
	}

	response, err := request.Do()
	if err != nil {
		return
	}

	body := response.Body()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	fmt.Println(string(data))

	err = body.Close()
	if err != nil {
		return
	}

	headers := response.Headers()

	contentType, err := headers.Get("Content-Type")
	if err != nil {
		return
	}
	fmt.Println(contentType[0])

	headerKeys, err := headers.List()
	if err != nil {
		return
	}
	fmt.Println(len(headerKeys))

	headerMap, err := headers.GetAll()
	if err != nil {
		return
	}
	fmt.Println(len(headerMap))

	// Output: Hello, world!
	// application/json
	// 1
	// 1
}
