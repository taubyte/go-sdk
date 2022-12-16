package storage

import (
	"testing"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/errno"
)

var (
	testCid   = "QmQUXYRNqU2U351aE8mrPFAyqXtKupF9bDspXKLsdkTLGn"
	testData  = "Yoda stores data then gives you a 'Care, I Do'"
	testWrite = 0
)

func TestContentError(t *testing.T) {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.ContentMockData{Id: 12}.Mock()

	symbols.StorageNewContent = func(contentIdPtr *uint32) (error errno.Error) {
		return 1
	}

	_, err := Create()
	if err == nil {
		t.Error("Expected Error")
	}

	_, err = Open(cid.Cid{})
	if err == nil {
		t.Error("Expected Error")
	}

	c := Content{}

	_, err = c.Push()
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = c.Write([]byte{})
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = c.Write([]byte("Hello, world"))
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

	symbols.ContentSeekFile = func(contentId uint32, offset int64, whence int, offsetPtr *int) (error errno.Error) {
		return errno.ErrorInvalidWhence
	}

	_, err = c.Seek(0, 0)
	if err == nil {
		t.Error("Expected error")
		return
	}

	c.Id = 5

	symbols.ContentFileCid = func(contentId uint32, cidPtr *byte) (error errno.Error) {
		return 5
	}

	_, err = c.Cid()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.ContentFileCid = func(contentId uint32, cidPtr *byte) (error errno.Error) {
		return 0
	}

	_, err = c.Cid()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.StorageOpenCid = func(contentIdPtr *uint32, cid *byte, cidSize uint32) (error errno.Error) {
		return 1
	}

	testCid, err := cid.Parse("QmQUXYRNqU2U351aE8mrPFAyqXtKupF9bDspXKLsdkTLGn")
	if err != nil {
		return
	}

	_, err = Open(testCid)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestContentPush(t *testing.T) {
	symbols.ContentMockData{Id: 2}.Mock()

	symbols.ContentSeekFile = func(contentId uint32, offset int64, whence int, offsetPtr *int) (error errno.Error) {
		return 0
	}

	c := &Content{}
	_, err := c.Push()
	if err == nil {
		t.Error(err)
		return
	}

	symbols.ContentPushFile = func(contentId uint32, cidPtr *byte) (error errno.Error) {
		return 0
	}

	_, err = c.Push()
	if err == nil {
		t.Error(err)
		return
	}
}
