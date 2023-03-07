package memoryView

import (
	"errors"
	"fmt"
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/i2mv/memoryView"
	"github.com/taubyte/go-sdk/utils/booleans"
)

func Open(id uint32) (MemoryView, error) {
	var size uint32
	var closable uint32
	if err := symbols.MemoryViewSize(id, &closable, &size); err != 0 {
		return nil, fmt.Errorf("getting memory view size failed with: %s", err)
	}

	return &memoryView{id: id, size: size, closable: booleans.ToBool(closable)}, nil
}

func (m *memoryView) Id() uint32 {
	return m.id
}

func (m *memoryView) OffSet() int64 {
	return m.offset
}

func (m *memoryView) Size() uint32 {
	return m.size
}

func (m *memoryView) Read(p []byte) (int, error) {
	var n uint32
	if len(p) == 0 {
		return 0, errors.New("cannot read to nil bytes")
	}

	if m.offset >= int64(m.size) {
		return 0, io.EOF
	}

	err := symbols.MemoryViewRead(m.id, uint32(m.offset), uint32(len(p)), &p[0], &n)
	if err != 0 {
		return 0, fmt.Errorf("reading memory view failed with: %s", err)
	}

	m.offset = m.offset + int64(n)

	return int(n), nil
}

func (m *memoryView) Seek(offSet int64, whence int) (int64, error) {
	var _whence int64
	switch whence {
	case io.SeekStart:
	case io.SeekCurrent:
		_whence = m.offset
	case io.SeekEnd:
		_whence = int64(m.size)
	default:
		return 0, fmt.Errorf("invalid whence value `%d`", whence)
	}

	if _offSet := _whence + offSet; _offSet < 0 {
		return 0, errors.New("cannot seek before start")
	} else if _offSet > int64(m.size) {
		m.offset = int64(m.size)
		return m.offset, io.EOF
	} else {
		m.offset = _offSet
		return m.offset, nil
	}
}

func (m *memoryView) Close() error {
	if m.closable {
		symbols.MemoryViewClose(m.id)
		return nil
	}

	return fmt.Errorf("memoryView `%d` not closable", m.id)
}
