package i2mv

import "io"

// TODO: Move to Interfaces once it is open sourced
type MemoryView interface {
	io.ReadSeekCloser
	Size() uint32
	OffSet() int64
	Id() uint32
}

var _ MemoryView = &memoryView{}

type memoryView struct {
	offset int64
	id     uint32
	size   uint32
}
