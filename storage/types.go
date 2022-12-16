package storage

import (
	"io"

	"github.com/ipfs/go-cid"
)

type Storage uint32

type Content struct {
	Id uint32
}

const CidBufferSize = 64

type StorageFile struct {
	storage Storage
	fd      uint32
}

type File struct {
	storage Storage
	name    string
	version uint32
}

type DefaultFile interface {
	Add(data []byte, overWrite bool) (int, error)
	GetFile() (file *StorageFile, err error)
	Delete() error
	DeleteAllVersions() error
	CurrentVersion() (int, error)
	Versions() ([]string, error)
}

type VersionedFile struct {
	DefaultFile
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
