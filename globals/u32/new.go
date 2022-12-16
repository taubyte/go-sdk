package u32

import (
	"encoding/binary"

	"github.com/taubyte/go-sdk/globals/internal"
)

func new(name string, scope ...internal.Option) *internal.Number[uint32] {
	return &internal.Number[uint32]{
		Base: &internal.Base[uint32]{
			Name:  name,
			Key:   "uint32",
			Scope: scope,

			ToBase: func(b []byte) (uint32, error) {
				return binary.BigEndian.Uint32(b), nil
			},

			ToBytes: func(v uint32) ([]byte, error) {
				a := make([]byte, 4)
				binary.BigEndian.PutUint32(a, v)

				return a, nil
			},
		},
	}
}
