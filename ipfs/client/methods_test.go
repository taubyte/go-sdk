package client

import (
	"testing"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/ipfs/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestMethods(t *testing.T) {
	symbols.MockData{
		ClientId: 1,
	}.Mock()

	var client Client
	fakeContent := &Content{}

	_, err := client.Create()
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = client.Open(cid.Cid{})
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = fakeContent.Push()
	if err == nil {
		t.Error("Expected error")
		return
	}

	trueClient := Client(1)
	content, err := trueClient.Create()
	if err != nil {
		t.Error(err)
		return
	}

	c := content.(*Content)
	c.client = client
	_, err = c.Push()
	if err == nil {
		t.Error("Expected error")
		return
	}

	c.client = trueClient
	cid, err := content.Push()
	if err != nil {
		t.Error(err)
		return
	}

	symbols.IpfsOpenFile = func(clientId uint32, contentIdPtr *uint32, cid *byte, cidSize uint32) (error errno.Error) {
		return 1
	}
	_, err = trueClient.Open(cid)
	if err == nil {
		t.Error("Expected error")
		return
	}
}
