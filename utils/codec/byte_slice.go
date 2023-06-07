package codec

import (
	"encoding/binary"
	"errors"
)

func (ds byteSliceDecoder) toByteSliceSlice(result *[][]byte) error {
	if *result == nil {
		*result = make([][]byte, 0)
	}

	for idx := 0; idx < len(ds); {
		if idx+2 >= len(ds) {
			break
		}
		size := int(binary.LittleEndian.Uint16(ds[idx : idx+2]))
		idx += 2
		if idx+size > len(ds) {
			break
		}
		*result = append(*result, ds[idx:idx+size])
		idx += size
	}
	return nil
}

type byteSliceSliceEncoder [][]byte

func (c byteSliceSliceEncoder) To(i interface{}) error {
	switch out := i.(type) {
	case []byte:
		return pointerError()
	case *[]byte:
		if *out == nil {
			*out = make([]byte, 0)
		}

		for _, s := range c {
			*out = binary.LittleEndian.AppendUint16(*out, uint16(len(s)))
			*out = append(*out, s...)
		}
	default:
		return errors.New("failed to encode Slice type")
	}
	return nil
}
