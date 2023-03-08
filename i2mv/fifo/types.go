package fifo

import "io"

// TODO: Move to Interfaces once it is open sourced
type Fifo interface {
	io.ReadCloser
	Id() uint32
}

var _ Fifo = &fifo{}

type fifo struct {
	id       uint32
	closable bool
}
