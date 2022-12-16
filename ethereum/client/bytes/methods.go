package bytes

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// AddressToHex returns the string value of an address. Addresses must be a 20 byte array.
func AddressToHex(value []byte) (string, error) {
	if len(value) != 20 {
		return "", fmt.Errorf("Invalid address, addresses are composed of 20 bytes, given value is `%d`", len(value))
	}

	return string(checkSumHex(value)), nil
}

// AddressFromHex will return the [20]byte value of a address hex string
// Any values exceeding [20]byte will be trimmed.
func AddressFromHex(s string) ([]byte, error) {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}

	return hex2Bytes(s)
}

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func hex2Bytes(str string) ([]byte, error) {
	return hex.DecodeString(str)
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
