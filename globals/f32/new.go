package f32

import (
	"encoding/binary"
	"math"

	"github.com/taubyte/go-sdk/globals/internal"
)

func new(name string, scope ...internal.Option) *internal.Number[float32] {
	return &internal.Number[float32]{
		Base: &internal.Base[float32]{
			Name:  name,
			Key:   "float32",
			Scope: scope,

			// https://stackoverflow.com/questions/22491876/convert-byte-slice-uint8-to-float32-in-golang
			ToBase: func(bytes []byte) (float32, error) {
				bits := binary.BigEndian.Uint32(bytes)
				float := math.Float32frombits(bits)
				return float, nil
			},

			ToBytes: func(v float32) ([]byte, error) {
				bits := math.Float32bits(v)
				bytes := make([]byte, 8)
				binary.BigEndian.PutUint32(bytes, bits)
				return bytes, nil
			},
		},
	}
}
