package internal

import (
	"path"
)

func (b *Base[T]) Set(v T) error {
	b.value = v

	return b.Put()
}

func (b *Base[T]) Value() T {
	return b.value
}

func (b *Base[T]) getKey() string {
	return path.Join("/"+b.Key, b.Name)
}

func (b *Base[T]) parse(getMethod func(name string) ([]byte, error)) error {
	v, err := getMethod(b.getKey())
	if b != nil {
		if len(v) == 0 {
			b.value = *new(T)
			return err
		}

		b.value, err = b.ToBase(v)
		if err != nil {
			return err
		}
	}

	return err
}

func (b *Base[T]) Get() error {
	return b.parse(instance(b.Scope...).Get)
}

func (b *Base[T]) GetOrCreate() error {
	return b.parse(instance(b.Scope...).GetOrCreate)
}

func (b *Base[T]) Put() error {
	v, err := b.ToBytes(b.value)
	if err != nil {
		return err
	}

	err = instance(b.Scope...).Put(b.getKey(), v)
	if err != nil {
		return err
	}

	return b.Get()
}
