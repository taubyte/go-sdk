package u64

import (
	"encoding/binary"

	"github.com/taubyte/go-sdk/globals/internal"
)

func new(name string, scope ...internal.Option) *internal.Number[uint64] {
	return &internal.Number[uint64]{
		Base: &internal.Base[uint64]{
			Name:  name,
			Key:   "uint64",
			Scope: scope,

			ToBase: func(b []byte) (uint64, error) {
				return binary.BigEndian.Uint64(b), nil
			},

			ToBytes: func(v uint64) ([]byte, error) {
				b := make([]byte, 8)
				binary.BigEndian.PutUint64(b, v)

				return b, nil
			},
		},
	}
}
