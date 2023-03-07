package memoryView

import "io"

type MemoryView interface {
	io.ReadSeekCloser
	Size() uint32
	OffSet() int64
	Id() uint32
}

var _ MemoryView = &memoryView{}

type memoryView struct {
	offset   int64
	id       uint32
	size     uint32
	closable bool
}
