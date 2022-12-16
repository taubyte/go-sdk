package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func (ds byteSliceDecoder) toInt32Slice(result *[]int32) error {
	if result == nil {
		return errors.New("result is nil")
	}

	*result = make([]int32, 0)
	for i := 0; i < len(ds); i += 4 {
		var s int32
		binary.Read(bytes.NewReader(ds[i:i+4]), binary.LittleEndian, &s)
		*result = append(*result, s)
	}

	return nil
}

type int32SliceEncoder []int32

func (c int32SliceEncoder) To(i interface{}) error {
	switch i.(type) {
	case []byte:
		return errors.New("Needs to be a pointer")
	case *[]byte:
		var buf bytes.Buffer
		out := i.(*[]byte)
		if *out == nil {
			*out = make([]byte, 0)
		}
		for _, s := range c {
			binary.Write(&buf, binary.LittleEndian, s)
		}
		*out = buf.Bytes()
	default:
		return errors.New("Failed to Encode:Unknown type")
	}
	return nil
}

func (ds byteSliceDecoder) toUInt32Slice(result *[]uint32) error {
	if result == nil {
		return errors.New("result is nil")
	}

	*result = make([]uint32, 0)
	for i := 0; i < len(ds); i += 4 {
		var s uint32
		binary.Read(bytes.NewReader(ds[i:i+4]), binary.LittleEndian, &s)
		*result = append(*result, s)
	}

	return nil
}

type uint32SliceEncoder []uint32

func (c uint32SliceEncoder) To(i interface{}) error {
	switch i.(type) {
	case []byte:
		return errors.New("Needs to be a pointer")
	case *[]byte:
		var buf bytes.Buffer
		out := i.(*[]byte)
		if *out == nil {
			*out = make([]byte, 0)
		}
		for _, s := range c {
			binary.Write(&buf, binary.LittleEndian, s)
		}
		*out = buf.Bytes()
	default:
		return errors.New("Failed to Encode:Unknown type")
	}
	return nil
}
