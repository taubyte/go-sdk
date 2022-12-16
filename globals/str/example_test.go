package str_test

import (
	"fmt"

	globalSym "github.com/taubyte/go-sdk-symbols/globals"
	"github.com/taubyte/go-sdk/globals/scope"
	"github.com/taubyte/go-sdk/globals/str"
)

func ExampleString() {
	// Mocking the calls to the vm for usage in tests and playground
	globalSym.MockData{
		Data: map[string][]uint8{
			"/string/name":             []byte("Hello, world!"),
			"/application/string/name": []byte("Hello, world! (scoped)"),
		},
	}.Mock()

	{
		name, err := str.GetOrCreate("name")
		if err != nil {
			return
		}

		fmt.Println("Global scope:", name.Value())

		err = name.Set("Hello, Mars!")
		if err != nil {
			return
		}

		name, err = str.Get("name")
		if err != nil {
			return
		}

		fmt.Println("Global scope set:", name.Value())
	}

	{
		name, err := str.Get("name", scope.Application)
		if err != nil {
			return
		}

		fmt.Println("Application scope:", name.Value())
	}

	// Output: Global scope: Hello, world!
	// Global scope set: Hello, Mars!
	// Application scope: Hello, world! (scoped)
}
