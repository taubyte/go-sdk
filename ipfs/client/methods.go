package client

import (
	"fmt"

	"github.com/ipfs/go-cid"
	ipfsClientSym "github.com/taubyte/go-sdk-symbols/ipfs/client"
	"github.com/taubyte/go-sdk/utils/codec"
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

	writer, err := codec.CidWriter(_cid)
	if err != nil {
		return nil, err
	}

	err0 := ipfsClientSym.IpfsOpenFile(uint32(c), &content.id, writer.Ptr(), writer.Len())
	if err0 != 0 {
		return nil, fmt.Errorf("opening file from cid `%s` failed with: %v", _cid.String(), err)
	}

	return content, nil
}
