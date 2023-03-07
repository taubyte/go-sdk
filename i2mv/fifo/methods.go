package fifo

import (
	"fmt"
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/i2mv/fifo"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/booleans"
)

func Open(id uint32) (Fifo, error) {
	var isCloser uint32
	if err := symbols.FifoIsCloser(id, &isCloser); err != 0 {
		return nil, fmt.Errorf("getting fifo failed with: %s", err)
	}

	return &fifo{id: id, closable: booleans.ToBool(isCloser)}, nil
}

func (f *fifo) Id() uint32 {
	return f.id
}

func (f *fifo) Write(data []byte) (n int, err error) {
	for _, buf := range data {
		if err := symbols.FifoPush(f.id, buf); err != 0 {
			return n, fmt.Errorf("pushing to fifo failed with: %s", err)
		}
		n++
	}

	return n, nil
}

func (f *fifo) Read(data []byte) (n int, err error) {
	copyData := make([]byte, 0, len(data))
	defer func() {
		copy(data, copyData)
	}()

	for range data {
		var buf byte
		if err := symbols.FifoPop(f.id, &buf); err != 0 {
			if err == errno.ErrorEOF {
				return n, io.EOF
			}

			return n, fmt.Errorf("popping from fifo failed with: %s", err)
		}

		copyData = append(copyData, buf)
		n++
	}

	return n, nil
}

func (f *fifo) Close() error {
	if f.closable {
		symbols.FifoClose(f.id)
		return nil
	}

	return fmt.Errorf("fifo `%d` is not closable", f.id)
}
