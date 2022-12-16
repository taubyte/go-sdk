package codec

import (
	"errors"
	"fmt"
)

type Convertable interface {
	To(i interface{}) error
}

type byteSliceDecoder []byte

func (c byteSliceDecoder) To(i interface{}) error {
	switch i.(type) {
	case []string:
		return errors.New("Needs to be a pointer")
	case *[]string:
		return c.toStringSlice(i.(*[]string))
	case []int32:
		return errors.New("Needs to be a pointer")
	case *[]int32:
		return c.toInt32Slice(i.(*[]int32))
	case []uint32:
		return errors.New("Needs to be a pointer")
	case *[]uint32:
		return c.toUInt32Slice(i.(*[]uint32))
	case [][]byte:
		return errors.New("Needs to be a pointer")
	case *[][]byte:
		return c.toByteSliceSlice(i.(*[][]byte))
	default:
		return errors.New("Convert: Unknown")
	}
}

type errorConvertable struct {
	err error
}

func (e errorConvertable) To(i interface{}) error {
	return e.err
}

func Convert(i interface{}) Convertable {
	switch i.(type) {
	case [][]byte:
		return byteSliceSliceEncoder(i.([][]byte))
	case []byte:
		return byteSliceDecoder(i.([]byte))
	case []string:
		return stringSliceEncoder(i.([]string))
	case []int32:
		return int32SliceEncoder(i.([]int32))
	case []uint32:
		return uint32SliceEncoder(i.([]uint32))
	default:
		return errorConvertable{err: fmt.Errorf("Convert: incompatible type %#v", i)}
	}
}
