package ethereum

import (
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

func TestHexToECDSABytes(t *testing.T) {
	testHexString := "test_hex_string"
	testPrivateKey := []byte("test_private_key")
	ethereumSym.MockHexToECDSA(testHexString, testPrivateKey, uint32(len(testPrivateKey)))

	_, err := HexToECDSABytes(testHexString)
	if err != nil {
		t.Errorf("HexToECDSABytes failed with: %s", err)
		return
	}

	_, err = HexToECDSABytes("hello world")
	if err == nil {
		t.Error("expected error")
		return
	}

	ethereumSym.MockHexToECDSA(testHexString, testPrivateKey, 3)

	_, err = HexToECDSABytes(testHexString)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestPublicKeyFromPrivate(t *testing.T) {
	testPrivateKey := []byte("test_private_key")
	ethereumSym.MockPublicKeyFromPrivate(testPrivateKey)

	_, err := PublicKeyFromPrivate(testPrivateKey)
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
	testMessage := []byte("hello_world")
	ethereumSym.MockPublicKeyFromSignedMessage(testMessage)

	_, err := PublicKeyFromSignedMessage(testMessage, testMessage)
	if err != nil {
		t.Errorf("getting public key from private failed with: %s", err)
		return
	}

	_, err = PublicKeyFromSignedMessage([]byte("hello_worlds"), testMessage)
	if err == nil {
		t.Error("expected error")
		return
	}
}
