package ints

import (
	"math/big"
)

func NewBigInt(buf []byte) *big.Int {
	num := new(big.Int)
	if len(buf) != 0 {
		return num.SetBytes(buf)
	}

	return num.SetUint64(0)
}
