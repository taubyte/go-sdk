package client

import (
	"errors"
	"fmt"

	"github.com/ipfs/go-cid"
	ipfsClientSym "github.com/taubyte/go-sdk-symbols/ipfs/client"
)

// Creates creates and returns the new content.
func (c Client) Create() (ReadWriteContent, error) {
	newContent := &Content{client: c}

	err := ipfsClientSym.IpfsNewContent(uint32(c), &newContent.id)
	if err != 0 {
		return nil, fmt.Errorf("Failed creating new content with %v", err)
	}

	return newContent, nil
}

// Open creates a new content using the cid given as the file.
// Returns a new content.
func (c Client) Open(_cid cid.Cid) (ReadOnlyContent, error) {
	content := &Content{
		client: c,
	}

	cidBytes := _cid.Bytes()
	if cidBytes == nil || len(cidBytes) == 0 {
		return nil, errors.New("Invalid cid")
	}

	err := ipfsClientSym.IpfsOpenFile(uint32(c), &content.id, &cidBytes[0], uint32(len(cidBytes)))
	if err != 0 {
		return nil, fmt.Errorf("Failed opening cid %s with: %v", _cid.String(), err)
	}

	return content, nil
}
