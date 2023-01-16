package ethereum

import (
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

var (
	fakeMessage   = []byte("hello world")
	fakePrivKey   = []byte("fake_priv_key")
	fakeSignature = []byte{1, 2, 3}
)

func TestSign(t *testing.T) {
	ethereumSym.MockSign(fakeMessage)

	_, err := SignMessage(fakeMessage, fakePrivKey)
	if err != nil {
		t.Errorf("Signing message failed with: %s", err)
		return
	}

	_, err = SignMessage([]byte("wrong worlds"), fakePrivKey)
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
	ethereumSym.MockVerify(fakeMessage, true)

	err := VerifySignature(fakeMessage, fakePrivKey, []byte{1, 2, 3})
	if err != nil {
		t.Errorf("Verifying signature failed with: %s", err)
		return
	}

	err = VerifySignature([]byte("helloWorld"), fakePrivKey, fakeSignature)
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockVerify(fakeMessage, false)

	err = VerifySignature(fakeMessage, fakePrivKey, []byte{1, 2, 3})
	if err == nil {
		t.Error("Expected error")
	}

	err = VerifySignature(nil, nil, nil)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestVerifySignInputs(t *testing.T) {
	if err := verifySignInputs(nil, fakePrivKey, fakeSignature, false, false); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(fakeMessage, nil, []byte{1, 2, 3}, false, false); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(fakeMessage, nil, []byte{1, 2, 3}, true, false); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(fakeMessage, fakePrivKey, nil, false, true); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(fakeMessage, fakePrivKey, []byte{}, false, true); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs(fakeMessage, fakePrivKey, nil, false, false); err != nil {
		t.Errorf("verifySignInputs failed with: %s", err)
		return
	}
}
