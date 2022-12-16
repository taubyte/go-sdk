package storage_test

import (
	"fmt"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/storage"
)

func ExampleNew() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockNew(32, "exampleStorage")

	db, err := storage.New("exampleStorage")
	if err != nil {
		return
	}

	fmt.Println(db)
	// Output: 32
}

func ExampleGet() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockGet(map[string]uint32{"exampleStorage": 16}, 16)

	db, err := storage.Get("exampleStorage")
	if err != nil {
		return
	}

	fmt.Println(db)
	// Output: 16

}

func ExampleCreate() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockCreate(10)

	content, err := storage.Create()
	if err != nil {
		return
	}

	fmt.Println(content)
	// Output: &{10}
}

func ExampleOpen() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockOpen(11, "QmQUXYRNqU2U351aE8mrPFAyqXtKupF9bDspXKLsdkTLGn")

	testCid, err := cid.Parse("QmQUXYRNqU2U351aE8mrPFAyqXtKupF9bDspXKLsdkTLGn")
	if err != nil {
		return
	}

	content, err := storage.Open(testCid)
	if err != nil {
		return
	}

	fmt.Println(content)
	// Output: &{11}
}
