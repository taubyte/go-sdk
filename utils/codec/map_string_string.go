package codec

import (
	"encoding/binary"
	"errors"
)

func (ds byteSliceDecoder) toMapStringString(result *map[string]string) error {
	stringMap := make(map[string]string, 0)
	for idx := 0; idx < len(ds); {
		if idx+2 >= len(ds) {
			break
		}
		keySize := int(binary.LittleEndian.Uint16(ds[idx : idx+2]))
		idx += 2
		if idx+keySize > len(ds) {
			break
		}
		valueIdx := idx + keySize
		key := string(ds[idx:valueIdx])
		valueSize := int(binary.LittleEndian.Uint16(ds[valueIdx : valueIdx+2]))
		value := string(ds[valueIdx+2 : valueIdx+2+valueSize])
		stringMap[key] = value
		idx = valueIdx + 2 + valueSize
	}
	*result = stringMap
	return nil
}

type mapStringStringEncoder map[string]string

func (c mapStringStringEncoder) To(i interface{}) error {
	switch i.(type) {
	case []byte:
		return errors.New("Needs to be a pointer")
	case *[]byte:
		out := i.(*[]byte)
		if *out == nil {
			*out = make([]byte, 0)
		}
		for key, value := range c {
			*out = binary.LittleEndian.AppendUint16(*out, uint16(len(key)))
			*out = append(*out, key...)
			*out = binary.LittleEndian.AppendUint16(*out, uint16(len(value)))
			*out = append(*out, value...)
		}
	default:
		return errors.New("failed to encode Slice type")
	}
	return nil
}
