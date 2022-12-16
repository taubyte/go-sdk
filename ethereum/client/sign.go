package ethereum

import (
	"fmt"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

// SignMessage returns an ECDSA signed message
func SignMessage(message, privKey string) ([]byte, error) {
	err := verifySignInputs(message, privKey, nil, false)
	if err != nil {
		return nil, fmt.Errorf("Invalid inputs: %s", err)
	}

	signature := make([]byte, EcdsaSignatureLength)
	err0 := ethereumSym.EthSignMessage(message, privKey, &signature[0])
	if err0 != 0 {
		return nil, fmt.Errorf("Signing message failed with: %s", err0)
	}

	return signature, nil
}

// VerifyMessage checks the signed ECDSA message with the original message to verify
// if the message was signed by the given private key
func VerifySignature(message, privKey string, signature []byte) error {
	err := verifySignInputs(message, privKey, signature, true)
	if err != nil {
		return fmt.Errorf("Invalid inputs: %s", err)
	}

	var isVerified uint32
	err0 := ethereumSym.EthVerifySignature(message, &signature[0], privKey, &isVerified)
	if err0 != 0 {
		return fmt.Errorf("Verifying signature failed with: %s", err0)
	}

	if isVerified == 0 {
		return fmt.Errorf("Signature not signed by this private key")
	}

	return nil
}

func verifySignInputs(message, privKey string, signature []byte, checkSig bool) error {
	if len(message) == 0 {
		return fmt.Errorf("Message is empty")
	}

	if len(privKey) == 0 {
		return fmt.Errorf("Private Key is empty")
	}

	if checkSig == true && (len(signature) == 0 || signature == nil) {
		return fmt.Errorf("Signature is nil")
	}

	return nil
}
