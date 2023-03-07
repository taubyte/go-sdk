package memoryView

import (
	"fmt"
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/i2mv/memoryView"
	"github.com/taubyte/go-sdk/utils/booleans"
)

func New(data []byte, isCloser bool) (id uint32, closer io.Closer, err error) {
	if len(data) == 0 {
		return 0, nil, fmt.Errorf("cannot create memory view from nil data")
	}

	err0 := symbols.MemoryViewNew(&data[0], uint32(len(data)), booleans.FromBool(isCloser), &id)
	if err0 != 0 {
		return 0, nil, fmt.Errorf("creating new memory view failed with: %s", err0)
	}

	return id, &memoryView{id: id}, nil
}
