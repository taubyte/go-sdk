package str

import (
	"github.com/taubyte/go-sdk/globals/internal"
)

func new(name string, scope ...internal.Option) *internal.String {
	return &internal.String{
		Base: &internal.Base[string]{
			Name: name,
			Key:  "string",

			ToBase: func(b []byte) (string, error) {
				return string(b), nil
			},

			ToBytes: func(s string) ([]byte, error) {
				return []byte(s), nil
			},

			Scope: scope,
		},
	}
}
