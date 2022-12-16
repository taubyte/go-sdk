package str

import "github.com/taubyte/go-sdk/globals/internal"

func Get(name string, scope ...internal.Option) (String, error) {
	s := new(name, scope...)
	return s, s.Get()
}

func GetOrCreate(name string, scope ...internal.Option) (String, error) {
	s := new(name, scope...)
	return s, s.GetOrCreate()
}
