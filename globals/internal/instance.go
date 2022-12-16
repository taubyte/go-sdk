package internal

import (
	"fmt"

	globalSym "github.com/taubyte/go-sdk-symbols/globals"
)

func instance(scope ...Option) *Instance {
	instance := &Instance{}
	if len(scope) > 0 {
		for _, opt := range scope {
			opt(instance)
		}
	}

	return instance
}

func (i *Instance) Get(name string) ([]byte, error) {
	var size uint32
	err := globalSym.GetGlobalValueSize(name, i.Application, i.Function, &size)
	if err != 0 {
		return nil, fmt.Errorf("Getting value from `%s` size failed with %s", name, err)
	}
	if size == 0 {
		return []byte{}, nil
	}

	data := make([]byte, size)
	err = globalSym.GetGlobalValue(name, i.Application, i.Function, &data[0])
	if err != 0 {
		return nil, fmt.Errorf("Getting value from `%s` failed with %s", name, err)
	}

	return data, nil
}

func (i *Instance) GetOrCreate(name string) ([]byte, error) {
	var size uint32
	err := globalSym.GetOrCreateGlobalValueSize(name, i.Application, i.Function, &size)
	if err != 0 {
		return nil, fmt.Errorf("Getting or creating value from `%s` size failed with %s", name, err)
	}
	if size == 0 {
		return []byte{}, nil
	}

	data := make([]byte, size)
	err = globalSym.GetGlobalValue(name, i.Application, i.Function, &data[0])
	if err != 0 {
		return nil, fmt.Errorf("Getting or creating value from `%s` failed with %s", name, err)
	}

	return data, nil
}

func (i *Instance) Put(name string, value []byte) error {
	err := globalSym.PutGlobalValue(name, i.Application, i.Function, value)
	if err != 0 {
		return fmt.Errorf("Putting value in `%s` failed with %s", name, err)
	}

	return nil
}
