package internal

import (
	"fmt"

	globalSym "github.com/taubyte/go-sdk-symbols/globals"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/sized"
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
	data, sizeErr, dataErr := sized.Read(
		func(size *uint32) errno.Error { return globalSym.GetGlobalValueSize(name, i.Application, i.Function, size) },
		func(data *byte) errno.Error { return globalSym.GetGlobalValue(name, i.Application, i.Function, data) },
	)
	if sizeErr != 0 {
		return nil, fmt.Errorf("Getting value from `%s` size failed with %s", name, sizeErr)
	}
	if dataErr != 0 {
		return nil, fmt.Errorf("Getting value from `%s` failed with %s", name, dataErr)
	}

	return data, nil
}

func (i *Instance) GetOrCreate(name string) ([]byte, error) {
	data, sizeErr, dataErr := sized.Read(
		func(size *uint32) errno.Error { return globalSym.GetOrCreateGlobalValueSize(name, i.Application, i.Function, size) },
		func(data *byte) errno.Error { return globalSym.GetGlobalValue(name, i.Application, i.Function, data) },
	)
	if sizeErr != 0 {
		return nil, fmt.Errorf("Getting or creating value from `%s` size failed with %s", name, sizeErr)
	}
	if dataErr != 0 {
		return nil, fmt.Errorf("Getting or creating value from `%s` failed with %s", name, dataErr)
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
