package dns_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/dns"
	"github.com/taubyte/go-sdk/dns"
)

func ExampleResolver_LookupTXT() {
	// Mocking the calls to the vm for usage in tests and playground
	testData := map[string][]byte{
		"google.com": symbols.Convert([]string{
			"MS=15412F62916164C0B20BB",
			"verification=J9cpOJM0nikft0jAgjmsQ",
		}),
	}
	symbols.MockData{TxtData: testData}.Mock()

	r, err := dns.NewResolver()
	if err != nil {
		return
	}

	values, err := r.LookupTXT("google.com")
	if err != nil {
		return
	}

	fmt.Println(values[0])
	// Output: MS=15412F62916164C0B20BB
}

func ExampleResolver_LookupAddress() {
	// Mocking the calls to the vm for usage in tests and playground
	testData := map[string][]byte{
		"8.8.8.8": symbols.Convert([]string{
			"dns.google",
		}),
	}
	symbols.MockData{AddressData: testData}.Mock()

	r, err := dns.NewResolver()
	if err != nil {
		return
	}

	values, err := r.LookupAddress("8.8.8.8")
	if err != nil {
		return
	}

	fmt.Println(values[0])
	// Output: dns.google
}

func ExampleResolver_LookupCNAME() {
	// Mocking the calls to the vm for usage in tests and playground
	testData := map[string][]byte{
		"taubyte.com.": symbols.Convert([]string{
			"nodes.taubyte.com.",
		}),
	}
	symbols.MockData{CnameData: testData}.Mock()

	r, err := dns.NewResolver()
	if err != nil {
		return
	}

	values, err := r.LookupCNAME("taubyte.com.")
	if err != nil {
		return
	}

	fmt.Println(values)
	// Output: nodes.taubyte.com.
}

func ExampleResolver_LookupMX() {
	// Mocking the calls to the vm for usage in tests and playground
	testData := map[string][]byte{
		"google.com": []byte("smtp.google.com.\x00/10\x00"),
	}
	symbols.MockData{MxData: testData}.Mock()

	r, err := dns.NewResolver()
	if err != nil {
		return
	}

	values, err := r.LookupMX("google.com")
	if err != nil {
		return
	}

	fmt.Println(values[0].Host, values[0].Pref)
	// Output: smtp.google.com. 10
}
