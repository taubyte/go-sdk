package codec

import (
	"errors"
	"fmt"

	geth "github.com/ethereum/go-ethereum/common"
	eth "github.com/taubyte/go-sdk/ethereum/client/bytes"
)

type Convertable interface {
	To(i interface{}) error
}

type byteSliceDecoder []byte

func (c byteSliceDecoder) To(i interface{}) error {
	switch val := i.(type) {
	case []string:
		return pointerError()
	case *[]string:
		return c.toStringSlice(val)
	case []int32:
		return pointerError()
	case *[]int32:
		return c.toInt32Slice(val)
	case []uint32:
		return pointerError()
	case *[]uint32:
		return c.toUInt32Slice(val)
	case [][]byte:
		return pointerError()
	case *[][]byte:
		return c.toByteSliceSlice(val)
	case []eth.Address:
		return pointerError()
	case *[]*eth.Address:
		return c.toSliceEthAddress(val)
	case []*eth.Hash:
		return pointerError()
	case *[]*eth.Hash:
		return c.toSliceEthHash(val)
	case [][]*eth.Hash:
		return pointerError()
	case *[][]*eth.Hash:
		return c.toSliceSliceEthHash(val)
	case []geth.Address:
		return pointerError()
	case *[]geth.Address:
		return c.toSliceGEthAddress(val)
	case []geth.Hash:
		return pointerError()
	case *[]geth.Hash:
		return c.toSliceGEthHash(val)
	case [][]geth.Hash:
		return pointerError()
	case *[][]geth.Hash:
		return c.toSliceSliceGEthHash(val)
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
	switch val := i.(type) {
	case [][]byte:
		return byteSliceSliceEncoder(val)
	case []byte:
		return byteSliceDecoder(val)
	case []string:
		return stringSliceEncoder(val)
	case []int32:
		return int32SliceEncoder(val)
	case []uint32:
		return uint32SliceEncoder(val)
	case []geth.Address:
		return sliceGEthAddressEncoder(val)
	case []geth.Hash:
		return sliceGEthHashEncoder(val)
	case [][]geth.Hash:
		return sliceSliceGEthHashEncoder(val)
	case []*eth.Address:
		return sliceEthAddressEncoder(val)
	case []*eth.Hash:
		return sliceEthHashEncoder(val)
	case [][]*eth.Hash:
		return sliceSliceEthHashEncoder(val)
	default:
		return errorConvertable{err: fmt.Errorf("Convert: incompatible type %#v", i)}
	}
}

func pointerError() error {
	return fmt.Errorf("needs to be a pointer")
}
