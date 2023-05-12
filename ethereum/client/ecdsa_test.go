package ethereum

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestHexToECDSABytes(t *testing.T) {
	_, err := HexToECDSABytes(testPrivateKeyHex)
	assert.NilError(t, err)

	_, err = HexToECDSABytes("hello world")
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestPublicKeyFromPrivate(t *testing.T) {
	privateKey, err := HexToECDSABytes(testPrivateKeyHex)
	assert.NilError(t, err)

	_, err = PublicKeyFromPrivate(privateKey)
	assert.NilError(t, err)

	_, err = PublicKeyFromPrivate([]byte("hello world"))
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestPublicKeyFromSignedMessage(t *testing.T) {
	privateKey, err := HexToECDSABytes(testPrivateKeyHex)
	assert.NilError(t, err)

	signature, err := SignMessage([]byte("hello world"), privateKey)
	assert.NilError(t, err)

	_, err = PublicKeyFromSignedMessage([]byte("hello world"), signature)
	assert.NilError(t, err)

	_, err = PublicKeyFromSignedMessage([]byte("hello_worlds"), privateKey)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestParseSignature(t *testing.T) {
	_, err := ParseSignature("asdf")
	if err == nil {
		t.Error("expected error")
		return
	}

	_, err = ParseSignature("0x62428d75d9f9f741e233941f844a6dce056d86b3159f0091bd3cc6e65b0bd23123e9fbdf107965716d94b6ad481462cf28d33c1323c515f411903b3abca41bc71c")
	assert.NilError(t, err)
}

func TestToEthJsMessage(t *testing.T) {
	msg := ToEthJsMessage("hello world")
	if !bytes.Equal(msg, []byte("\x19Ethereum Signed Message:\n"+strconv.Itoa(len("hello world"))+"hello world")) {
		t.Error("bytes are not the same")
	}
}

func TestAddressFromPubKey(t *testing.T) {
	privateKey, err := HexToECDSABytes(testPrivateKeyHex)
	assert.NilError(t, err)

	publicKey, err := PublicKeyFromPrivate(privateKey)
	assert.NilError(t, err)

	address := AddressFromPubKey(publicKey)
	_ = address.Bytes() // for test coverage

	if address.String() != strings.ToLower(testAddress) {
		t.Errorf("expected `%s` got `%s`", testAddress, address.String())
	}

}
