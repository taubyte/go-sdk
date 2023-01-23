package ethereum

import (
	"testing"
)

func TestSign(t *testing.T) {
	privateKey, err := HexToECDSABytes(testPrivateKeyHex)
	if err != nil {
		t.Errorf("HexToECDSABytes failed with: %s", err)
		return
	}

	_, err = SignMessage([]byte("hello world"), privateKey)
	if err != nil {
		t.Errorf("Signing message failed with: %s", err)
		return
	}

	_, err = SignMessage([]byte("hello world"), []byte("hello world"))
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = SignMessage(nil, nil)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestSignatureVerify(t *testing.T) {
	privateKey, err := HexToECDSABytes(testPrivateKeyHex)
	if err != nil {
		t.Errorf("HexToECDSABytes failed with: %s", err)
		return
	}

	publicKey, err := PublicKeyFromPrivate(privateKey)
	if err != nil {
		t.Errorf("getting public key failed with: %s", err)
	}

	signature, err := SignMessage([]byte("hello world"), privateKey)
	if err != nil {
		t.Errorf("Signing message failed with: %s", err)
		return
	}

	err = VerifySignature([]byte("hello world"), publicKey, signature)
	if err != nil {
		t.Errorf("Verifying signature failed with: %s", err)
		return
	}

	err = VerifySignature([]byte("helloworld"), publicKey, []byte{1, 2, 3})
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = VerifySignature([]byte("helloworld"), publicKey, signature)
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = VerifySignature(nil, nil, nil)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestVerifySignInputs(t *testing.T) {
	if err := verifySignInputs(nil, testBytes, testBytes, false, false); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(testBytes, nil, testBytes, false, false); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(testBytes, nil, testBytes, true, false); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(testBytes, testBytes, nil, false, true); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(testBytes, testBytes, []byte{}, false, true); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(testBytes, testBytes, nil, false, false); err != nil {
		t.Errorf("verifySignInputs failed with: %s", err)
		return
	}
}
