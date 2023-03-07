package storage

import (
	"errors"
	"fmt"
	"io"

	"github.com/ipfs/go-cid"
	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

// Create creates and returns the new content.
func Create() (ReadWriteContent, error) {
	var contentId uint32
	err := storageSym.StorageNewContent(&contentId)
	if err != 0 {
		return nil, fmt.Errorf("Failed creating new content with %v", err)
	}

	return &Content{contentId}, nil
}

// Open creates a new content using the cid given as the file.
// Returns a new content.
func Open(_cid cid.Cid) (ReadOnlyContent, error) {
	content := &Content{}

	writer, err := codec.CidWriter(_cid)
	if err != nil {
		return nil, err
	}

	err0 := storageSym.StorageOpenCid(&content.Id, writer.Ptr(), writer.Len())
	if err0 != 0 {
		return nil, fmt.Errorf("Failed opening cid %s with %v", _cid.String(), err0)
	}

	return content, nil
}

// Write writes the passed in data into the file.
// Returns how much was written and an error.
func (c *Content) Write(p []byte) (int, error) {
	if p == nil || len(p) == 0 {
		return 0, errors.New("Invalid buffer to write into the file")
	}

	var n uint32
	err := storageSym.ContentWriteFile(c.Id, &p[0], uint32(len(p)), &n)
	if err != 0 {
		return 0, fmt.Errorf("Failed content write with %v", err)
	}

	return int(n), nil
}

// Close closes the file associated with the content.
// Returns an error.
func (c *Content) Close() error {
	err := storageSym.ContentCloseFile(c.Id)
	if err != 0 {
		return fmt.Errorf("Failed closing content with %v", err)
	}

	return nil
}

// Cid returns the cid of the file and an error.
func (c *Content) Cid() (cid.Cid, error) {
	_cid := codec.CidReader()

	err0 := storageSym.ContentFileCid(c.Id, _cid.Ptr())
	if err0 != 0 {
		return cid.Cid{}, fmt.Errorf("Failed getting cid with %v", err0)
	}

	return _cid.Parse()
}

// Read reads up to len p in the file.
// Returns how much was read and an error.
func (c *Content) Read(p []byte) (int, error) {
	if p == nil || len(p) == 0 {
		return 0, errors.New("Invalid buffer to read into")
	}

	var counter uint32
	err := storageSym.ContentReadFile(c.Id, &p[0], uint32(len(p)), &counter)
	if err != 0 {
		if err == errno.ErrorEOF {
			return int(counter), io.EOF
		} else {
			return 0, fmt.Errorf("Failed reading content with %v", err)
		}
	}
	return int(counter), nil
}

// Push adds the file into the network.
// Updates the cid of the file.
// Returns cid and an error
func (c *Content) Push() (cid.Cid, error) {
	_, err0 := c.Seek(0, 0)
	if err0 != nil {
		return cid.Cid{}, fmt.Errorf("Failed seeking beginning of content with: %v", err0)
	}

	_cid := codec.CidReader()

	err := storageSym.ContentPushFile(c.Id, _cid.Ptr())
	if err != 0 {
		return cid.Cid{}, fmt.Errorf("Failed reading content with %v", err)
	}

	// Closing content so it cannot be modified
	c.Close()

	return _cid.Parse()
}

// Seek moves to a position inside the file.
// Offset is how much to move the current position
// Whence has three options: 0 = SeekStart, 1 = SeekCurrent, or 2 = SeekEnd
// Combines both offset and whence to find a new offset inside the file
// Returns the new offset and an error
func (c *Content) Seek(offset int64, whence int) (int64, error) {
	var n int
	err := storageSym.ContentSeekFile(c.Id, offset, whence, &n)
	if err == errno.ErrorInvalidWhence {
		return 0, errors.New("Invalid Valid. Valid whence are 0 = SeekStart, 1 = SeekCurrent, or 2 = SeekEnd")
	}

	if err != 0 {
		return 0, fmt.Errorf("Failed seek with %v", err)
	}

	return int64(n), nil
}
