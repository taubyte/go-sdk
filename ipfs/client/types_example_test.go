package client_test

import (
	"fmt"
	"os"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/ipfs/client"
	"github.com/taubyte/go-sdk/ipfs/client"
)

func ExampleContent_Cid() {
	_cid, _ := cid.Parse("bafkreibrl5n5w5wqpdcdxcwaazheualemevr7ttxzbutiw74stdvrfhn2m")
	symbols.MockData{
		Files: map[cid.Cid]*os.File{
			_cid: nil,
		},
	}.Mock()

	client, err := client.New()
	if err != nil {
		return
	}

	content, err := client.Open(_cid)
	if err != nil {
		return
	}

	cid, err := content.Cid()
	if err != nil {
		return
	}

	fmt.Println("Cid:", cid.String())
	// Output: Cid: bafkreibrl5n5w5wqpdcdxcwaazheualemevr7ttxzbutiw74stdvrfhn2m
}
