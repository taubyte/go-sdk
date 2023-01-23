package ethereum

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestHexToECDSABytes(t *testing.T) {
	_, err := HexToECDSABytes(testPrivateKeyHex)
	if err != nil {
		t.Errorf("HexToECDSABytes failed with: %s", err)
		return
	}

	_, err = HexToECDSABytes("hello world")
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestPublicKeyFromPrivate(t *testing.T) {
	privateKey, err := HexToECDSABytes(testPrivateKeyHex)
	if err != nil {
		t.Errorf("HexToECDSABytes failed with: %s", err)
		return
	}

	_, err = PublicKeyFromPrivate(privateKey)
	if err != nil {
		t.Errorf("getting public key from private failed with: %s", err)
		return
	}

	_, err = PublicKeyFromPrivate([]byte("hello world"))
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestPublicKeyFromSignedMessage(t *testing.T) {
	privateKey, err := HexToECDSABytes(testPrivateKeyHex)
	if err != nil {
		t.Errorf("HexToECDSABytes failed with: %s", err)
		return
	}

	signature, err := SignMessage([]byte("hello world"), privateKey)
	if err != nil {
		t.Errorf("signing message failed with: %s", err)
	}
	_, err = PublicKeyFromSignedMessage([]byte("hello world"), signature)
	if err != nil {
		t.Errorf("getting public key from private failed with: %s", err)
		return
	}

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
	if err != nil {
		t.Errorf("parsing signature failed with: %s", err)
		return
	}
}

func TestToEthJsMessage(t *testing.T) {
	msg := ToEthJsMessage("hello world")
	if bytes.Compare(msg, []byte("\x19Ethereum Signed Message:\n"+strconv.Itoa(len("hello world"))+"hello world")) != 0 {
		t.Error("bytes are not the same")
	}
}

func TestAddressFromPubKey(t *testing.T) {
	privateKey, err := HexToECDSABytes(testPrivateKeyHex)
	if err != nil {
		t.Errorf("HexToECDSABytes failed with: %s", err)
		return
	}

	publicKey, err := PublicKeyFromPrivate(privateKey)
	if err != nil {
		t.Errorf("getting public key from private failed with: %s", err)
		return
	}

	address := AddressFromPubKey(publicKey)
	_ = address.Bytes() // for test coverage

	if address.String() != strings.ToLower(testAddress) {
		t.Errorf("expected `%s` got `%s`", testAddress, address.String())
	}

}
