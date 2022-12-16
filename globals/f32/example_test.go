package f32_test

import (
	"fmt"

	globalSym "github.com/taubyte/go-sdk-symbols/globals"
	"github.com/taubyte/go-sdk/globals/f32"
	"github.com/taubyte/go-sdk/globals/scope"
)

func ExampleFloat32() {
	// Mocking the calls to the vm for usage in tests and playground
	globalSym.MockData{
		Data: map[string][]uint8{
			"/float32/count":             {0x41, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			"/application/float32/count": {0x43, 0xad, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
	}.Mock()

	{
		count, err := f32.GetOrCreate("count")
		if err != nil {
			return
		}

		fmt.Println("Global scope:", count.Value())

		err = count.Set(64)
		if err != nil {
			return
		}

		count, err = f32.Get("count")
		if err != nil {
			return
		}

		fmt.Println("Global scope set:", count.Value())
	}

	{
		count, err := f32.Get("count", scope.Application)
		if err != nil {
			return
		}

		fmt.Println("Application scope:", count.Value())
	}

	// Output: Global scope: 12
	// Global scope set: 64
	// Application scope: 347
}
