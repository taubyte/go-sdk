package f32

import "github.com/taubyte/go-sdk/globals/internal"

type Float32 interface {
	internal.BaseInterface[float32]
	internal.NumberInterface[float32]
}
