package utils

import (
	"bytes"
	"math/big"
	"math/rand"
	"reflect"
	"testing"

	eth "github.com/taubyte/go-sdk/ethereum/client/bytes"
	codec "github.com/taubyte/go-sdk/utils/codec"
	"gotest.tools/v3/assert"
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

	var new_fruits [][]byte
	err = codec.Convert(encoded).To(&new_fruits)
	if err != nil {
		t.Error(err)
		return
	}

	for idx, fruit := range fruits {
		if !bytes.Equal(fruit, new_fruits[idx]) {
			t.Errorf("`%v` != `%v`", new_fruits[idx], fruit)
			return
		}
	}
}

func TestEncodingSliceEthHash(t *testing.T) {
	hashBuffer1 := make([]byte, eth.HashByteLength)
	hashBuffer2 := make([]byte, eth.HashByteLength)

	_, err := rand.Read(hashBuffer1)
	assert.NilError(t, err)

	_, err = rand.Read(hashBuffer2)
	assert.NilError(t, err)

	hashes := []*eth.Hash{eth.BytesToHash(hashBuffer1), eth.BytesToHash(hashBuffer2)}
	var encoded []byte
	err = codec.Convert(hashes).To(&encoded)
	assert.NilError(t, err)

	var newHashes []*eth.Hash
	err = codec.Convert(encoded).To(&newHashes)
	assert.NilError(t, err)

	if !reflect.DeepEqual(hashes, newHashes) {
		t.Error(err)
		return
	}
}

func TestEncodingSliceSliceEthHash(t *testing.T) {
	hashBuffer1 := make([]byte, eth.HashByteLength)
	hashBuffer2 := make([]byte, eth.HashByteLength)
	hashBuffer3 := make([]byte, eth.HashByteLength)
	hashBuffer4 := make([]byte, eth.HashByteLength)

	_, err := rand.Read(hashBuffer1)
	assert.NilError(t, err)

	_, err = rand.Read(hashBuffer2)
	assert.NilError(t, err)

	_, err = rand.Read(hashBuffer3)
	assert.NilError(t, err)

	_, err = rand.Read(hashBuffer4)
	assert.NilError(t, err)

	hashes := [][]*eth.Hash{
		{eth.BytesToHash(hashBuffer1), eth.BytesToHash(hashBuffer2)},
		{eth.BytesToHash(hashBuffer3), eth.BytesToHash(hashBuffer4)},
	}

	var encoded []byte
	err = codec.Convert(hashes).To(&encoded)
	assert.NilError(t, err)

	var newHashes [][]*eth.Hash
	err = codec.Convert(encoded).To(&newHashes)
	assert.NilError(t, err)

	if !reflect.DeepEqual(hashes, newHashes) {
		t.Error(err)
		return
	}
}

func TestEncodingSliceEthAddress(t *testing.T) {
	addressBuffer1 := make([]byte, eth.AddressByteLength)
	addressBuffer2 := make([]byte, eth.AddressByteLength)

	_, err := rand.Read(addressBuffer1)
	assert.NilError(t, err)

	_, err = rand.Read(addressBuffer2)
	assert.NilError(t, err)

	addresses := []*eth.Address{eth.BytesToAddress(addressBuffer1), eth.BytesToAddress(addressBuffer2)}
	var encoded []byte
	err = codec.Convert(addresses).To(&encoded)
	assert.NilError(t, err)

	var newAddresses []*eth.Address
	err = codec.Convert(encoded).To(&newAddresses)
	assert.NilError(t, err)

	if !reflect.DeepEqual(addresses, newAddresses) {
		t.Error(err)
		return
	}
}

func TestEncodingMapStringUint64(t *testing.T) {
	testMap := map[string]uint64{
		"name": 420,
		"test": 390,
	}

	var encoded []byte
	err := codec.Convert(testMap).To(&encoded)
	assert.NilError(t, err)

	var newMap map[string]uint64
	err = codec.Convert(encoded).To(&newMap)
	assert.NilError(t, err)

	assert.DeepEqual(t, testMap, newMap)
}

func TestEncodingMapStringBool(t *testing.T) {
	testMap := map[string]bool{
		"name": false,
		"test": true,
	}

	var encoded []byte
	err := codec.Convert(testMap).To(&encoded)
	assert.NilError(t, err)

	var newMap map[string]bool
	err = codec.Convert(encoded).To(&newMap)
	assert.NilError(t, err)

	assert.DeepEqual(t, testMap, newMap)
}
