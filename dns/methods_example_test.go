package dns_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/dns"
	"github.com/taubyte/go-sdk/dns"
)

func ExampleResolver_Reroute() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockNew(0)
	symbols.MockReroute(0, "8.8.8.8:53", "udp")

	r, err := dns.NewResolver()
	if err != nil {
		return
	}

	err = r.Reroute("8.8.8.8:53", "udp")
	if err != nil {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleResolver_Reset() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{}.Mock()

	r, err := dns.NewResolver()
	if err != nil {
		return
	}

	err = r.Reset()
	if err != nil {
		return
	}

	fmt.Println("Success")
	// Output: Success
}
