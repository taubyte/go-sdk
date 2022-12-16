package utils

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	codec "github.com/taubyte/go-sdk/utils/codec"
)

func TestEncodingString(t *testing.T) {
	fruits := []string{"apple", "banana", "orange"}
	var encoded []byte
	err := codec.Convert(fruits).To(&encoded)
	if err != nil {
		t.Error(err)
		return
	}

	var new_fruits []string
	err = codec.Convert(encoded).To(&new_fruits)
	if err != nil {
		t.Error(err)
		return
	}

	for idx, fruit := range fruits {
		if fruit != new_fruits[idx] {
			t.Errorf("`%s` != `%s`", new_fruits[idx], fruit)
			return
		}
	}
}

func TestEncodingInt(t *testing.T) {
	nums := []int32{789779, 5667, 342347}
	var encoded []byte
	err := codec.Convert(nums).To(&encoded)
	if err != nil {
		t.Error(err)
		return
	}

	var new_nums []int32
	err = codec.Convert(encoded).To(&new_nums)
	if err != nil {
		t.Error(err)
		return
	}

	for idx, num := range nums {
		if num != new_nums[idx] {
			t.Errorf("`%d` != `%d`", new_nums[idx], num)
			return
		}
	}
}

func TestEncodingUInt32(t *testing.T) {
	nums := []uint32{789779, 5667, 342347}
	var encoded []byte
	err := codec.Convert(nums).To(&encoded)
	if err != nil {
		t.Error(err)
		return
	}

	var new_nums []uint32
	err = codec.Convert(encoded).To(&new_nums)
	if err != nil {
		t.Error(err)
		return
	}

	for idx, num := range nums {
		if num != new_nums[idx] {
			t.Errorf("`%d` != `%d`", new_nums[idx], num)
			return
		}
	}
}

func TestEncodingByteSliceSlice(t *testing.T) {
	fruits := [][]byte{[]byte("apple"), big.NewInt(123).Bytes(), {0, 2, 3, 65, 4, 65, 87, 52}}
	var encoded []byte
	err := codec.Convert(fruits).To(&encoded)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(encoded)
	var new_fruits [][]byte
	err = codec.Convert(encoded).To(&new_fruits)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(new_fruits)
	for idx, fruit := range fruits {
		if bytes.Compare(fruit, new_fruits[idx]) != 0 {
			t.Errorf("`%v` != `%v`", new_fruits[idx], fruit)
			return
		}
	}
}
