package node_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/p2p/node"
	"github.com/taubyte/go-sdk/p2p/node"
)

func ExampleService_Command() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		Protocol:  "/test/v1",
		Command:   "someCommand",
		CommandId: 5,
	}.Mock()

	service := node.New("/test/v1")

	// Instantiate a command `someCommand` to protocol `/test/v1`
	command, err := service.Command("someCommand")
	if err != nil {
		return
	}

	fmt.Println(command)
	// Output: 5
}

func ExampleService_Protocol() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		Protocol: "/test/v1",
	}.Mock()

	service := node.New("/test/v1")

	fmt.Println(service.Protocol())
	// Output: /test/v1
}

func ExampleService_Listen() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		ListenHash:     "QmZjBpQzR",
		ListenProtocol: "/test/v1",
	}.Mock()

	// Instantiate a service with protocol `/test/v1`
	service := node.New("/test/v1")

	listenProtocol, err := service.Listen()
	if err != nil {
		return
	}

	fmt.Println(listenProtocol)
	// Output: /QmZjBpQzR/test/v1
}
