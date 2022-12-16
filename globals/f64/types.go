package f64

import "github.com/taubyte/go-sdk/globals/internal"

type Float64 interface {
	internal.BaseInterface[float64]
	internal.NumberInterface[float64]
}
