package rand_test

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/crypto/rand"
	"github.com/taubyte/go-sdk/crypto/rand"
)

func TestReader(t *testing.T) {
	buffer := make([]byte, 32)

	symbols.MockReader(32, 32)
	reader := rand.NewReader()

	_, err := reader.Read(buffer)
	if err != nil {
		t.Error(err)
		return
	}

	symbols.MockReader(12, 12)
	_, err = reader.Read(buffer)
	if err == nil {
		t.Error("Expected error")
	}

}
