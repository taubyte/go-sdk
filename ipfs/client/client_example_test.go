package client_test

import (
	"fmt"
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/ipfs/client"
	"github.com/taubyte/go-sdk/ipfs/client"
)

func ExampleClient() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{}.Mock()

	ipfsClient, err := client.New()
	if err != nil {
		return
	}

	content, err := ipfsClient.Create()
	if err != nil {
		return
	}

	_, err = content.Write([]byte("Hello, world!"))
	if err != nil {
		return
	}

	cid, err := content.Push()
	if err != nil {
		return
	}

	newContent, err := ipfsClient.Open(cid)
	if err != nil {
		return
	}

	data, err := io.ReadAll(newContent)
	if err != nil {
		return
	}

	err = newContent.Close()
	if err != nil {
		return
	}

	newCid, err := newContent.Cid()
	if err != nil || cid != newCid {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
	// Output: Hello, world!
}
