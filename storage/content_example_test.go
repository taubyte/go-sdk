package storage_test

import (
	"fmt"
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/storage"
)

func ExampleContent() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.ContentMockData{Id: 2}.Mock()

	content, err := storage.Create()
	if err != nil {
		return
	}

	_, err = content.Write([]byte("Hello World!"))
	if err != nil {
		return
	}

	_, err = content.Seek(0, 0)
	if err != nil {
		return
	}

	cid, err := content.Push()
	if err != nil {
		return
	}

	newContent, err := storage.Open(cid)
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
		return
	}

	fmt.Println(string(data))
	// Output: Hello World!
}
