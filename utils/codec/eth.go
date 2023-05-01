package codec

import (
	"errors"

	gethCommon "github.com/ethereum/go-ethereum/common"
	eth "github.com/taubyte/go-sdk/ethereum/client/bytes"
)

func (ds byteSliceDecoder) toSliceGEthAddress(result *[]gethCommon.Address) error {
	if *result == nil {
		*result = make([]gethCommon.Address, 0)
	}

	for idx := 0; idx < len(ds); {
		*result = append(*result, gethCommon.BytesToAddress(ds[idx:idx+gethCommon.AddressLength]))
		idx += gethCommon.AddressLength
	}

	return nil
}

type sliceGEthAddressEncoder []gethCommon.Address

func (e sliceGEthAddressEncoder) To(i interface{}) error {
	switch out := i.(type) {
	case []byte:
		return errors.New("needs to be a pointer")
	case *[]byte:
		if *out == nil {
			*out = make([]byte, 0)
		}

		for _, s := range e {
			*out = append(*out, s.Bytes()...)
		}
	}

	return nil
}

func (ds byteSliceDecoder) toSliceGEthHash(result *[]gethCommon.Hash) error {
	if *result == nil {
		*result = make([]gethCommon.Hash, 0)
	}

	for idx := 0; idx < len(ds); {
		*result = append(*result, gethCommon.BytesToHash(ds[idx:idx+gethCommon.HashLength]))
		idx += gethCommon.HashLength
	}

	return nil
}

type sliceGEthHashEncoder []gethCommon.Hash

func (s sliceGEthHashEncoder) To(i interface{}) error {
	switch out := i.(type) {
	case []byte:
		return errors.New("needs to be a pointer")
	case *[]byte:
		if *out == nil {
			*out = make([]byte, 0)
		}

		for _, s := range s {
			*out = append(*out, s.Bytes()...)
		}
	}

	return nil
}

func (ds byteSliceDecoder) toSliceSliceGEthHash(result *[][]gethCommon.Hash) error {
	resBytesSliceSlice := make([][]byte, 0)
	if err := ds.toByteSliceSlice(&resBytesSliceSlice); err != nil {
		return err
	}

	*result = make([][]gethCommon.Hash, len(resBytesSliceSlice))

	for idx, buf := range resBytesSliceSlice {

		if err := byteSliceDecoder(buf).toSliceGEthHash(&(*result)[idx]); err != nil {
			return err
		}
	}

	return nil
}

type sliceSliceGEthHashEncoder [][]gethCommon.Hash

func (s sliceSliceGEthHashEncoder) To(i interface{}) error {
	switch out := i.(type) {
	case []byte:
		return errors.New("needs to be a pointer")
	case *[]byte:
		if *out == nil {
			*out = make([]byte, 0)
		}

		bufSlice := make([][]byte, len(s))
		for idx, s := range s {
			buf := make([]byte, 0)
			if err := sliceGEthHashEncoder(s).To(&buf); err != nil {
				return err
			}

			bufSlice[idx] = buf
		}

		return byteSliceSliceEncoder(bufSlice).To(i)
	}

	return nil
}

func (ds byteSliceDecoder) toSliceEthAddress(result *[]*eth.Address) error {
	if *result == nil {
		*result = make([]*eth.Address, 0)
	}

	for idx := 0; idx < len(ds); {

		*result = append(*result, eth.BytesToAddress(ds[idx:idx+eth.AddressByteLength]))
		idx += eth.AddressByteLength
	}

	return nil
}

type sliceEthAddressEncoder []*eth.Address

func (e sliceEthAddressEncoder) To(i interface{}) error {
	switch out := i.(type) {
	case []byte:
		return errors.New("needs to be a pointer")
	case *[]byte:
		if *out == nil {
			*out = make([]byte, 0)
		}

		for _, s := range e {
			*out = append(*out, s.Bytes()...)
		}
	}

	return nil
}

func (ds byteSliceDecoder) toSliceEthHash(result *[]*eth.Hash) error {
	if *result == nil {
		*result = make([]*eth.Hash, 0)
	}

	for idx := 0; idx < len(ds); {
		*result = append(*result, eth.BytesToHash(ds[idx:idx+eth.HashByteLength]))
		idx += eth.HashByteLength
	}

	return nil
}

type sliceEthHashEncoder []*eth.Hash

func (s sliceEthHashEncoder) To(i interface{}) error {
	switch out := i.(type) {
	case []byte:
		return errors.New("needs to be a pointer")
	case *[]byte:
		if *out == nil {
			*out = make([]byte, 0)
		}

		for _, s := range s {
			*out = append(*out, s.Bytes()...)
		}
	}

	return nil
}

func (ds byteSliceDecoder) toSliceSliceEthHash(result *[][]*eth.Hash) error {
	resBytesSliceSlice := make([][]byte, 0)
	if err := ds.toByteSliceSlice(&resBytesSliceSlice); err != nil {
		return err
	}

	*result = make([][]*eth.Hash, len(resBytesSliceSlice))

	for idx, buf := range resBytesSliceSlice {

		if err := byteSliceDecoder(buf).toSliceEthHash(&(*result)[idx]); err != nil {
			return err
		}
	}

	return nil
}

type sliceSliceEthHashEncoder [][]*eth.Hash

func (s sliceSliceEthHashEncoder) To(i interface{}) error {
	switch out := i.(type) {
	case []byte:
		return errors.New("needs to be a pointer")
	case *[]byte:
		if *out == nil {
			*out = make([]byte, 0)
		}

		bufSlice := make([][]byte, len(s))
		for idx, s := range s {
			buf := make([]byte, 0)
			if err := sliceEthHashEncoder(s).To(&buf); err != nil {
				return err
			}

			bufSlice[idx] = buf
		}

		return byteSliceSliceEncoder(bufSlice).To(i)
	}

	return nil
}
