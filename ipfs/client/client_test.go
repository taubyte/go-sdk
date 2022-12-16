package client

import (
	"testing"

	ipfsClientSym "github.com/taubyte/go-sdk-symbols/ipfs/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestNew(t *testing.T) {
	ipfsClientSym.NewIPFSClient = func(clientId *uint32) (error errno.Error) {
		return errno.ErrorCap
	}
	_, err := New()
	if err == nil {
		t.Error("Expected error")
		return
	}

	testId := uint32(6)
	ipfsClientSym.NewIPFSClient = func(clientId *uint32) (error errno.Error) {
		*clientId = testId
		return 0
	}

	client, err := New()
	if err != nil {
		t.Error(err)
		return
	}
	if uint32(client) != testId {
		t.Errorf("Got client `%d`, expected `%d`", uint32(client), testId)
		return
	}
}
