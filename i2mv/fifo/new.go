package fifo

import (
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/i2mv/fifo"
	"github.com/taubyte/go-sdk/utils/booleans"
)

// New creates a new fifo.
// Fifo is always closable by the creator, but only closable by accessor if closable is set to true.
func New(closable bool) (uint32, io.WriteCloser) {
	id := symbols.FifoNew(booleans.FromBool(closable))
	return id, &fifo{id: id, closable: true}
}
