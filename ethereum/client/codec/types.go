package codec

type converterType string

// converter contains methods for encoding and decoding supported Solidity types to and from bytes.
//
//	Supported Types:
//	- Ints
//	- UInts
//	- Booleans
//	- Addresses ([20]byte)
//	- Enums  (Named Ints)
//	- Bytes
//	- Arrays (Of any type listed above)
type converter struct {
	Encoder encoder
	Decoder decoder
}

// encoder returns the method used to encode the solidity supported value to bytes.
type encoder func(interface{}) ([]byte, error)

// decoder returns the method used to decode bytes to the corresponding value type.
type decoder func([]byte) (interface{}, error)
