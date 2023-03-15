package ethereum

import (
	"testing"

	"gotest.tools/assert"
)

func TestSign(t *testing.T) {
	privateKey, err := HexToECDSABytes(testPrivateKeyHex)
	assert.NilError(t, err)

	_, err = SignMessage([]byte("hello world"), privateKey)
	assert.NilError(t, err)

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
	assert.NilError(t, err)

	publicKey, err := PublicKeyFromPrivate(privateKey)
	assert.NilError(t, err)

	signature, err := SignMessage([]byte("hello world"), privateKey)
	assert.NilError(t, err)

	err = VerifySignature([]byte("hello world"), publicKey, signature)
	assert.NilError(t, err)

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

	err := verifySignInputs(testBytes, testBytes, nil, false, false)
	assert.NilError(t, err)
}
