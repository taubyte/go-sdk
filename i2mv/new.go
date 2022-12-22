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
