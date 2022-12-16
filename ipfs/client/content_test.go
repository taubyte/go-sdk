package client

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/ipfs/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestContentError(t *testing.T) {
	symbols.MockData{
		ClientId: 1,
	}.Mock()

	c := &Content{}
	_, err := c.Write([]byte("Hello, world"))
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = c.Close()
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = c.Read(nil)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = c.Read([]byte{0})
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = c.Seek(0, 0)
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.IpfsSeekFile = func(clientId, contentId uint32, offset int64, whence int, offsetPtr *int) (error errno.Error) {
		return errno.ErrorInvalidWhence
	}

	_, err = c.Seek(0, 0)
	if err == nil {
		t.Error("Expected error")
		return
	}

	c.client = 1
	symbols.IpfsFileCid = func(clientId, contentId uint32, cidPtr *byte) (error errno.Error) {
		return 1
	}
	_, err = c.Cid()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.IpfsFileCid = func(clientId, contentId uint32, cidPtr *byte) (error errno.Error) {
		return 0
	}
	_, err = c.Cid()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestContentPush(t *testing.T) {
	symbols.MockData{
		ClientId: 1,
	}.Mock()

	symbols.IpfsSeekFile = func(clientId, contentId uint32, offset int64, whence int, offsetPtr *int) (error errno.Error) {
		return 0
	}

	c := &Content{}
	_, err := c.Push()
	if err == nil {
		t.Error(err)
		return
	}

	symbols.IpfsPushFile = func(clientId, contentId uint32, cidPtr *byte) (error errno.Error) {
		return 0
	}

	_, err = c.Push()
	if err == nil {
		t.Error(err)
		return
	}
}
