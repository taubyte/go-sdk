package u64

import "github.com/taubyte/go-sdk/globals/internal"

type Uint64 interface {
	internal.BaseInterface[uint64]
	internal.NumberInterface[uint64]
}
