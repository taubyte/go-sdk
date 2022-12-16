package internal

import "golang.org/x/exp/constraints"

type Instance struct {
	Application uint32
	Function    uint32
}

type Option func(i *Instance)

type Base[T any] struct {
	value T

	Name string
	Key  string

	ToBase  func([]byte) (T, error)
	ToBytes func(T) ([]byte, error)
	Scope   []Option
}

type BaseInterface[T any] interface {
	Value() T
	Set(v T) error
}

type NumberInterface[T constraints.Unsigned | constraints.Float] interface {
	BaseInterface[T]
}

type Number[T constraints.Unsigned | constraints.Float] struct {
	*Base[T]
}

type StringInterface interface {
	BaseInterface[string]
}

type String struct {
	*Base[string]
}
