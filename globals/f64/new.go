package f64

import (
	"encoding/binary"
	"math"

	"github.com/taubyte/go-sdk/globals/internal"
)

func new(name string, scope ...internal.Option) *internal.Number[float64] {
	return &internal.Number[float64]{
		Base: &internal.Base[float64]{
			Name:  name,
			Key:   "float64",
			Scope: scope,

			// https://stackoverflow.com/questions/22491876/convert-byte-slice-uint8-to-float64-in-golang
			ToBase: func(bytes []byte) (float64, error) {
				bits := binary.BigEndian.Uint64(bytes)
				float := math.Float64frombits(bits)
				return float, nil
			},

			ToBytes: func(v float64) ([]byte, error) {
				bits := math.Float64bits(v)
				bytes := make([]byte, 8)
				binary.BigEndian.PutUint64(bytes, bits)
				return bytes, nil
			},
		},
	}
}
