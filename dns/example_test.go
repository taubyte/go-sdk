package dns_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/dns"
	"github.com/taubyte/go-sdk/dns"
	"github.com/taubyte/go-sdk/event"
)

func dnsFunction(e event.Event) uint32 {
	resolver, err := dns.NewResolver()
	if err != nil {
		return 1
	}

	values, err := resolver.LookupAddress("8.8.8.8")
	if err != nil {
		return 1
	}
	fmt.Println("Address:", values)

	value, err := resolver.LookupCNAME("taubyte.com.")
	if err != nil {
		return 1
	}
	fmt.Println("CNAME:", value)

	mxValues, err := resolver.LookupMX("google.com")
	if err != nil {
		return 1
	}
	fmt.Println("MX:", mxValues[0].Host, mxValues[0].Pref)

	values, err = resolver.LookupTXT("google.com")
	if err != nil {
		return 1
	}
	fmt.Println("TxT:", values[0])

	err = resolver.Reroute("8.8.8.8:53", "udp")
	if err != nil {
		return 1
	}

	err = resolver.Reset()
	if err != nil {
		return 1
	}

	return 0
}

func ExampleResolver() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{Id: 5,
		TxtData: map[string][]byte{
			"google.com": symbols.Convert([]string{
				"MS=15412F62916164C0B20BB",
				"verification=J9cpOJM0nikft0jAgjmsQ",
			}),
		},
		AddressData: map[string][]byte{
			"8.8.8.8": symbols.Convert([]string{
				"dns.google",
			}),
		},
		CnameData: map[string][]byte{
			"taubyte.com.": symbols.Convert([]string{
				"nodes.taubyte.com.",
			}),
		},
		MxData: map[string][]byte{
			"google.com": []byte("smtp.google.com.\x00/10\x00"),
		},
	}.Mock()

	e := dnsFunction(0)
	if e > 0 {
		fmt.Println(e)
		return
	}

	fmt.Println("Success")
	// Output:
	// Address: [dns.google]
	// CNAME: nodes.taubyte.com.
	// MX: smtp.google.com. 10
	// TxT: MS=15412F62916164C0B20BB
	// Success
}
