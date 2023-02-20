package event

import (
	"fmt"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/utils/codec"
)

func (e Event) Query() Queries {
	return Queries(e)
}

func (e Queries) Get(key string) (string, error) {
	var size uint32
	err := httpEventSym.GetHttpEventQueryValueByNameSize(uint32(e), &size, key)
	if err != 0 {
		return "", fmt.Errorf("Getting HTTP query size for key: `%s` failed with: %s", key, err)
	}
	if size == 0 {
		return "", nil
	}

	query := make([]byte, size)
	err = httpEventSym.GetHttpEventQueryValueByName(uint32(e), key, &query[0], uint32(len(query)))
	if err != 0 {
		return "", fmt.Errorf("Getting HTTP query for key: `%s` failed with: %s", key, err)
	}

	return string(query), nil
}

func (e Queries) List() ([]string, error) {
	var size uint32
	err := httpEventSym.GetHttpEventRequestQueryKeysSize(uint32(e), &size)
	if err != 0 {
		return nil, fmt.Errorf("Getting all query keys size failed with: %s", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	query := make([]byte, size)
	if err := httpEventSym.GetHttpEventRequestQueryKeys(uint32(e), &query[0]); err != 0 {
		return nil, fmt.Errorf("Getting all query keys failed with: %s", err)
	}

	var queries []string
	err0 := codec.Convert(query).To(&queries)
	if err0 != nil || len(queries) == 0 {
		return nil, fmt.Errorf("Converting query slice to string slice failed with: %s", err0)
	}

	return queries, nil
}
