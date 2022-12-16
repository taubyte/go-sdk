package client

import (
	"fmt"

	ipfsClientSym "github.com/taubyte/go-sdk-symbols/ipfs/client"
)

func New() (Client, error) {
	var clientId uint32
	err := ipfsClientSym.NewIPFSClient(&clientId)
	if err != 0 {
		return 0, fmt.Errorf("Creating a new ipfs client failed with: %s", err)
	}

	return Client(clientId), nil
}
