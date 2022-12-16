package ethereum

import (
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

func TestSign(t *testing.T) {
	ethereumSym.MockSign("hello world")

	_, err := SignMessage("hello world", "fake_priv_key")
	if err != nil {
		t.Errorf("Signing message failed with: %s", err)
		return
	}

	_, err = SignMessage("helloWorld", "fake_priv_key")
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = SignMessage("", "")
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestSignatureVerify(t *testing.T) {
	ethereumSym.MockVerify("hello world", true)

	err := VerifySignature("hello world", "fake_priv_key", []byte{1, 2, 3})
	if err != nil {
		t.Errorf("Verifying signature failed with: %s", err)
		return
	}

	err = VerifySignature("helloWorld", "fake_priv_key", []byte{1, 2, 3})
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockVerify("hello world", false)

	err = VerifySignature("hello world", "fake_priv_key", []byte{1, 2, 3})
	if err == nil {
		t.Error("Expected error")
	}

	err = VerifySignature("", "", nil)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestVerifySignInputs(t *testing.T) {
	if err := verifySignInputs("", "fake_priv_key", []byte{1, 2, 3}, false); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs("hello world", "", []byte{1, 2, 3}, false); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs("hello world", "fake_priv_key", nil, true); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs("hello world", "fake_priv_key", []byte{}, true); err == nil {
		t.Error("Expected error")
		return
	}

	if err := verifySignInputs("hello world", "fake_priv_key", nil, false); err != nil {
		t.Errorf("verifySignInputs failed with: %s", err)
		return
	}
}
