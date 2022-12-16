package u64

import "github.com/taubyte/go-sdk/globals/internal"

func Get(name string, scope ...internal.Option) (Uint64, error) {
	u := new(name, scope...)

	return u, u.Get()
}

func GetOrCreate(name string, scope ...internal.Option) (Uint64, error) {
	u := new(name, scope...)

	return u, u.GetOrCreate()
}
