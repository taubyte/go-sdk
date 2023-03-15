package fifo_test

import (
	"fmt"
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/i2mv/fifo"
	"github.com/taubyte/go-sdk/i2mv/fifo"
)

var (
	writeCloser io.WriteCloser
	id          uint32
)

func ExampleNew() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockNew(1)
	symbols.MockPush(1)

	id, writeCloser = fifo.New(true)
	if id != 1 {
		fmt.Printf("expected id `1` got `%d`", id)
		return
	}

	n, err := writeCloser.Write([]byte("hello world"))
	if err != nil {
		fmt.Printf("writing to fifo failed with: %s\n", err)
		return
	}

	fmt.Println(n)
	// Output: 11
}

func ExampleOpen() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockOpen(2, true)

	ff, err := fifo.Open(2)
	if err != nil {
		fmt.Printf("error opening fifo with: %s\n", err)
		return
	}

	fmt.Println(ff.Id())
	// Output: 2
}

func ExampleFifo_Read() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockAll(3, true, []byte("hello world"))

	ff, err := fifo.Open(3)
	if err != nil {
		fmt.Printf("error opening fifo with: %s\n", err)
		return
	}

	data := make([]byte, 11)
	_, err = ff.Read(data)
	if err != nil {
		fmt.Printf("reading from fifo failed with: %s\n", err)
		return
	}

	fmt.Println(string(data))
	// Output: hello world
}

func ExampleFifo_Close() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockOpen(4, true)

	ff, err := fifo.Open(4)
	if err != nil {
		fmt.Printf("error opening fifo with: %s\n", err)
		return
	}

	if err = ff.Close(); err != nil {
		fmt.Printf("closing fifo failed with: %s\n", err)
		return
	}

	fmt.Println(ff.Id())
	// Output: 4
}
