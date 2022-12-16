package codec

import (
	"errors"
)

func (ds byteSliceDecoder) toStringSlice(result *[]string) error {
	if *result == nil {
		*result = make([]string, 0)
	}

	var lastIndex int
	for idx := 0; idx < len(ds); idx++ {
		if ds[idx] == 0 {
			*result = append(*result, string(ds[lastIndex:idx]))
			lastIndex = idx + 1
		}
	}

	return nil
}

type stringSliceEncoder []string

func (c stringSliceEncoder) To(i interface{}) error {
	switch i.(type) {
	case []byte:
		return errors.New("Needs to be a pointer")
	case *[]byte:
		out := i.(*[]byte)
		if *out == nil {
			*out = make([]byte, 0)
		}
		for _, s := range c {
			*out = append(*out, []byte(s)...)
			*out = append(*out, 0)
		}
	default:
		return errors.New("failed to encode Slice type")
	}

	return nil
}
