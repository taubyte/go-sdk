package i2mv_test

import (
	"fmt"
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/i2mv"
	"github.com/taubyte/go-sdk/i2mv"
)

var mvCloser io.Closer
var err error
var id uint32
var n int
var offset int64

func ExampleNew() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockNew(1)

	id, mvCloser, err = i2mv.New([]byte{1, 2, 3}, true)
	if err != nil {
		fmt.Printf("New memory view failed with: failed with: %s\n", err)
		return
	}

	fmt.Printf("%d\n", id)
	// Output: 1
}

func ExampleOpen() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockSize(5)

	mv, err := i2mv.Open(0)
	if err != nil {
		fmt.Printf("Open failed with: %s\n", err)
		return
	}

	fmt.Printf("%d\n", mv.Size())
	// Output: 5
}

func ExampleMemoryView_Read() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockSize(5)
	symbols.MockRead(1)

	mv, err := i2mv.Open(1)
	if err != nil {
		fmt.Printf("Open failed with: %s\n", err)
		return
	}

	data := make([]byte, 5)
	n, err = mv.Read(data)
	if err != nil {
		fmt.Printf("Read failed with: %s\n", err)
		return
	}

	n, err = mv.Read(data)
	if err != io.EOF {
		fmt.Printf("Expected error `%s` got `%s`", io.EOF, err)
	}

	fmt.Println("success")
	// Output: success
}

func ExampleMemoryView_Seek() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockSize(5)

	mv, err := i2mv.Open(1)
	if err != nil {
		fmt.Printf("Open failed with: %s\n", err)
		return
	}

	offset, err = mv.Seek(3, io.SeekCurrent)
	if err != nil {
		fmt.Printf("Seek failed with: %s\n", err)
		return
	}

	fmt.Printf("%d\n", offset)
	// Output 3
}

func ExampleMemoryView_Close() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockSize(5)
	symbols.MockClose(true)

	mv, err := i2mv.Open(1)
	if err != nil {
		fmt.Printf("Open failed with: %s\n", err)
		return
	}

	err = mv.Close()
	if err != nil {
		fmt.Printf("Close failed with: %s\n", err)
		return
	}

	fmt.Println("success")
	// Output success
}