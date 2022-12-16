package event

import (
	"fmt"

	eventSym "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk/utils/codec"
)

func (e HttpEvent) Query() HttpQueries {
	return HttpQueries(e)
}

func (e HttpQueries) Get(key string) (string, error) {
	var size uint32
	err := eventSym.GetHttpEventQueryValueByNameSize(uint32(e), &size, key)
	if err != 0 {
		return "", fmt.Errorf("Getting HTTP query size for key: `%s` failed with: %s", key, err)
	}
	if size == 0 {
		return "", nil
	}

	query := make([]byte, size)
	err = eventSym.GetHttpEventQueryValueByName(uint32(e), key, &query[0], uint32(len(query)))
	if err != 0 {
		return "", fmt.Errorf("Getting HTTP query for key: `%s` failed with: %s", key, err)
	}

	return string(query), nil
}

func (e HttpQueries) List() ([]string, error) {
	var size uint32
	err := eventSym.GetHttpEventRequestQueryKeysSize(uint32(e), &size)
	if err != 0 {
		return nil, fmt.Errorf("Getting all query keys size failed with: %s", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	query := make([]byte, size)
	if err := eventSym.GetHttpEventRequestQueryKeys(uint32(e), &query[0]); err != 0 {
		return nil, fmt.Errorf("Getting all query keys failed with: %s", err)
	}

	var queries []string
	err0 := codec.Convert(query).To(&queries)
	if err0 != nil || len(queries) == 0 {
		return nil, fmt.Errorf("Converting query slice to string slice failed with: %s", err0)
	}

	return queries, nil
}
