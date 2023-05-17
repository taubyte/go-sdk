package codec

import (
	"encoding/binary"
	"errors"

	"github.com/taubyte/go-sdk/utils/booleans"
)

func (ds byteSliceDecoder) toMapStringUInt64(result *map[string]uint64) error {
	_result := make(map[string]uint64)
	intBufferSize := 8
	for idx := 0; idx < len(ds); {
		if idx+2 >= len(ds) {
			break
		}
		size := int(binary.LittleEndian.Uint16(ds[idx : idx+2]))
		idx += 2
		if idx+size+intBufferSize > len(ds) {
			break
		}

		key := string(ds[idx : idx+size])
		val := binary.LittleEndian.Uint64(ds[idx+size : idx+size+intBufferSize])
		_result[key] = val
		idx += (intBufferSize + size)
	}

	*result = _result
	return nil
}

type mapStringUint64Encoder map[string]uint64

func (m mapStringUint64Encoder) To(i interface{}) error {
	switch out := i.(type) {
	case []byte:
		return errors.New("needs to be a pointer")
	case *[]byte:
		if *out == nil {
			*out = make([]byte, 0)
		}

		for key, val := range m {
			*out = binary.LittleEndian.AppendUint16(*out, uint16(len([]byte(key))))
			*out = append(*out, []byte(key)...)
			*out = binary.LittleEndian.AppendUint64(*out, val)
		}
	default:
		return errors.New("failed to encode map string uint64")
	}

	return nil
}

func (ds byteSliceDecoder) toMapStringBool(result *map[string]bool) error {
	_result := make(map[string]bool)
	for idx := 0; idx < len(ds); {
		if idx+2 >= len(ds) {
			break
		}
		size := int(binary.LittleEndian.Uint16(ds[idx : idx+2]))
		idx += 2
		if idx+size+1 > len(ds) {
			break
		}

		key := string(ds[idx : idx+size])

		val := booleans.ToBool(uint32(ds[idx+size]))
		_result[key] = val
		idx += (1 + size)
	}

	*result = _result
	return nil
}

type mapStringBoolEncoder map[string]bool

func (m mapStringBoolEncoder) To(i interface{}) error {
	switch out := i.(type) {
	case []byte:
		return errors.New("needs to be a pointer")
	case *[]byte:
		if *out == nil {
			*out = make([]byte, 0)
		}

		for key, val := range m {
			*out = binary.LittleEndian.AppendUint16(*out, uint16(len([]byte(key))))
			*out = append(*out, []byte(key)...)
			*out = append(*out, byte(booleans.FromBool(val)))
		}
	default:
		return errors.New("failed to encode map string uint64")
	}

	return nil
}
