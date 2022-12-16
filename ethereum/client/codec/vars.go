package codec

import (
	"encoding/binary"
	"fmt"
	"math/big"
)

// converterMap defines the basic types which are supported by the Ethereum Client.
// Note that besides []byte type, array types are not listed.
// Array types will be handled internally by the Encoder and Decoder methods.
var converterMap = map[converterType]converter{
	"*big.Int": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(*big.Int)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a *big.Int", val)
			}

			return value.Bytes(), nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			value := new(big.Int)
			return value.SetBytes(val), nil
		},
	},
	"uint8": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(uint8)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a uint8", val)
			}

			return []byte{value}, nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			if len(val) != 1 {
				return nil, fmt.Errorf("Unexpected value `%s` expected uint8 value of byte length 1", val)
			}
			return val[0], nil
		},
	},
	"uint16": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(uint16)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a uint16", val)
			}

			buf := make([]byte, 2)
			binary.LittleEndian.PutUint16(buf, value)

			return buf, nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			if len(val) != 2 {
				return nil, fmt.Errorf("Unexpected value `%s` expected uint16 value of byte length 2", val)
			}

			return binary.LittleEndian.Uint16(val), nil
		},
	},
	"uint32": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(uint32)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a uint32", val)
			}

			buf := make([]byte, 4)
			binary.LittleEndian.PutUint32(buf, value)

			return buf, nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			if len(val) != 4 {
				return nil, fmt.Errorf("Unexpected value `%s` expected uint32 value of byte length 4", val)
			}

			return binary.LittleEndian.Uint32(val), nil
		},
	},
	"uint64": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(uint64)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a uint64", val)
			}

			buf := make([]byte, 8)
			binary.LittleEndian.PutUint64(buf, value)

			return buf, nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			if len(val) != 8 {
				return nil, fmt.Errorf("Unexpected value `%s` expected uint64 value of byte length 5", val)
			}

			return binary.LittleEndian.Uint64(val), nil
		},
	},
	"int8": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(int8)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a int8", val)
			}

			var neg uint8
			if value < 0 {
				value = value * -1
				neg = 1
			}

			return []byte{neg, uint8(value)}, nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			if len(val) != 2 {
				return nil, fmt.Errorf("Unexpected value `%s` expected int8 value of byte length 2", val)
			}

			value := int8(val[1])
			if val[0] == 1 {
				value = value * -1
			}

			return value, nil
		},
	},
	"int16": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(int16)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a int16", val)
			}

			var neg uint8
			if value < 0 {
				value = value * -1
				neg = 1
			}

			b := make([]byte, 2)
			binary.LittleEndian.PutUint16(b, uint16(value))

			return append([]byte{neg}, b...), nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			if len(val) != 3 {
				return nil, fmt.Errorf("Unexpected value `%s` expected int16 value of byte length 3", val)
			}

			value := int16(binary.LittleEndian.Uint16(val[1:]))
			if val[0] == 1 {
				value = value * -1
			}

			return value, nil
		},
	},
	"int32": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(int32)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a int32", val)
			}

			var neg uint8
			if value < 0 {
				value = value * -1
				neg = 1
			}

			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, uint32(value))

			return append([]byte{neg}, b...), nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			if len(val) != 5 {
				return nil, fmt.Errorf("Unexpected value `%s` expected int32 value of byte length 5", val)
			}

			value := int32(binary.LittleEndian.Uint32(val[1:]))
			if val[0] == 1 {
				value = value * -1
			}

			return value, nil
		},
	},
	"int64": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(int64)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a int64", val)
			}

			var neg uint8
			if value < 0 {
				value = value * -1
				neg = 1
			}

			b := make([]byte, 4)
			binary.LittleEndian.PutUint64(b, uint64(value))

			return append([]byte{neg}, b...), nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			if len(val) != 9 {
				return nil, fmt.Errorf("Unexpected value `%s` expected int32 value of byte length 9", val)
			}

			value := int64(binary.LittleEndian.Uint64(val[1:]))
			if val[0] == 1 {
				value = value * -1
			}

			return value, nil
		},
	},
	"[]byte": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.([]byte)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a uint32", val)
			}

			return value, nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			return val, nil
		},
	},
	"common.Address": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.([]byte)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not an address", val)
			}

			return value, nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			return val, nil
		},
	},
	"[]uint8": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.([]uint8)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a []uint8", val)
			}

			return value, nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			return val, nil
		},
	},
	"bool": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(bool)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a bool", val)
			}

			if value == false {
				return []byte{0}, nil
			}

			return []byte{1}, nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			if len(val) != 1 {
				return nil, fmt.Errorf("Unexpected value `%s` expected bool value of byte length 1", val)
			}

			var value bool
			if val[0] == 1 {
				value = true
			}

			return value, nil
		},
	},
	"string": {
		Encoder: func(val interface{}) ([]byte, error) {
			value, ok := val.(string)
			if ok == false {
				return nil, fmt.Errorf("Value %s is not a string", val)
			}

			return []byte(value), nil
		},
		Decoder: func(val []byte) (interface{}, error) {
			return string(val), nil
		},
	},
}
