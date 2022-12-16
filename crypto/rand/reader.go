package rand

import (
	"fmt"
	"io"

	symbols "github.com/taubyte/go-sdk-symbols/crypto/rand"
)

func NewReader() io.Reader {
	return reader{}
}

func (r reader) Read(p []byte) (int, error) {
	var read uint64
	err0 := symbols.CryptoRead(&p[0], uint32(len(p)), &read)
	if err0 != 0 {
		return 0, fmt.Errorf("Reading failed with: %s", err0)
	}

	return int(read), nil
}
