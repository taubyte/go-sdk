package i2mv

import (
	"fmt"
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/i2mv"
)

func New(data []byte, readCloser bool) (id uint32, closer io.Closer, err error) {
	if data == nil || len(data) == 0 {
		return 0, nil, fmt.Errorf("cannot create memory view from nil data")
	}

	var closable uint32
	if readCloser {
		closable = 1
	}

	err0 := symbols.MemoryViewNew(&data[0], uint32(len(data)), closable, &id)
	if err0 != 0 {
		return 0, nil, fmt.Errorf("creating new memory view failed with: %s", err)
	}

	return id, &memoryView{id: id}, nil
}

func Open(id uint32) (MemoryView, error) {
	var size uint32
	if err := symbols.MemoryViewSize(id, &size); err != 0 {
		return nil, fmt.Errorf("getting memory view size failed with: %s", err)
	}
	return &memoryView{id: id, size: size}, nil
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
	if p == nil || len(p) == 0 {
		return 0, fmt.Errorf("Cannot read to nil bytes")
	}

	err := symbols.MemoryViewRead(m.id, uint32(m.offset), uint32(len(p)), &p[0], &n)
	if err != 0 {
		return 0, fmt.Errorf("reading memory view failed with: %s", err)
	}

	m.offset = m.offset + int64(n)

	return int(n), nil
}

func (m *memoryView) Seek(offSet int64, whence int) (int64, error) {
	current := m.offset
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
		return 0, fmt.Errorf("cannot seek before start")
	} else if _offSet > int64(m.size) {
		m.offset = int64(m.size)
		return m.offset - current, io.EOF
	} else {
		m.offset = _offSet
		return m.offset - current, nil
	}
}

func (m *memoryView) Close() error {
	if err := symbols.MemoryViewClose(m.id); err != 0 {
		return fmt.Errorf("closing memory view failed with: %s", err)
	}

	return nil
}
