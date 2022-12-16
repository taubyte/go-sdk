package f64_test

import (
	"fmt"

	globalSym "github.com/taubyte/go-sdk-symbols/globals"
	"github.com/taubyte/go-sdk/globals/f64"
	"github.com/taubyte/go-sdk/globals/scope"
)

func ExampleFloat64() {
	// Mocking the calls to the vm for usage in tests and playground
	globalSym.MockData{
		Data: map[string][]uint8{
			"/float64/count":             {0x40, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			"/application/float64/count": {0x40, 0x80, 0xa8, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
	}.Mock()

	{
		count, err := f64.GetOrCreate("count")
		if err != nil {
			return
		}

		fmt.Println("Global scope:", count.Value())

		err = count.Set(23)
		if err != nil {
			return
		}

		count, err = f64.Get("count")
		if err != nil {
			return
		}

		fmt.Println("Global scope set:", count.Value())
	}

	{
		count, err := f64.Get("count", scope.Application)
		if err != nil {
			return
		}

		fmt.Println("Application scope:", count.Value())
	}

	// Output: Global scope: 16
	// Global scope set: 23
	// Application scope: 533
}
