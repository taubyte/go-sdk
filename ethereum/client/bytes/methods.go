package bytes

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// AddressToHex returns the string value of an address. Addresses must be a 20 byte array.
func AddressToHex(value []byte) (string, error) {
	if len(value) != 20 {
		return "", fmt.Errorf("invalid address, addresses are composed of 20 bytes, given value is `%d`", len(value))
	}

	return string(checkSumHex(value)), nil
}

// AddressFromHex will return the [20]byte value of a address hex string
// Any values exceeding [20]byte will be trimmed.
func AddressFromHex(s string) (*Address, error) {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}

	buf, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return BytesToAddress(buf), nil
}

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func hexBytes(value []byte) []byte {
	var buf [20*2 + 2]byte
	copy(buf[:2], "0x")
	hex.Encode(buf[2:], value[:])
	return buf[:]
}

func checkSumHex(value []byte) []byte {
	buf := hexBytes(value)

	sha := sha3.NewLegacyKeccak256()
	sha.Write(buf[2:])
	hash := sha.Sum(nil)
	for i := 2; i < len(buf); i++ {
		hashByte := hash[(i-2)/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if buf[i] > '9' && hashByte > 7 {
			buf[i] -= 32
		}
	}
	return buf[:]
}

// String returns the 0x prefixed hex string representation of the address.
func (a *Address) String() string {
	return "0x" + hex.EncodeToString(a[:])
}

// Bytes returns the address as []byte representation.
func (a *Address) Bytes() []byte {
	return a[:]
}

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressByteLength:]
	}
	copy(a[AddressByteLength-len(b):], b)
}

func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashByteLength:]
	}
	copy(h[HashByteLength-len(b):], b)
}

func (h *Hash) Bytes() []byte {
	return h[:]
}

func BytesToAddress(val []byte) *Address {
	var a Address
	a.SetBytes(val)
	return &a
}

func BytesToHash(val []byte) *Hash {
	var h Hash
	h.SetBytes(val)
	return &h
}
