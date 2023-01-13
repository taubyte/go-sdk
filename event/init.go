package event

import (
	"crypto/rand"

	vmRand "github.com/taubyte/go-sdk/crypto/rand"
)

func init() {
	// crypto/rand library would otherwise not find a random source.
	rand.Reader = vmRand.NewReader()
}
