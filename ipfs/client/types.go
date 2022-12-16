package client

import (
	"io"

	"github.com/ipfs/go-cid"
)

type Client uint32

type Content struct {
	id     uint32
	client Client
}

type ReadWriteContent interface {
	io.ReadWriteSeeker
	io.Closer
	Push() (cid.Cid, error)
}

type ReadOnlyContent interface {
	io.ReadSeekCloser
	Cid() (cid.Cid, error)
}

const CidBufferSize = 64
