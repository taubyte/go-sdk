package ethereum

import (
	"fmt"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

// SignMessage returns an ECDSA signed message
func SignMessage(message, privKey []byte) ([]byte, error) {
	err := verifySignInputs(message, privKey, nil, true, false)
	if err != nil {
		return nil, fmt.Errorf("Invalid inputs: %s", err)
	}

	signature := make([]byte, EcdsaSignatureLength)
	err0 := ethereumSym.EthSignMessage(&message[0], uint32(len(message)), &privKey[0], uint32(len(privKey)), &signature[0])
	if err0 != 0 {
		return nil, fmt.Errorf("Signing message failed with: %s", err0)
	}

	return signature, nil
}

// VerifySignature checks the signed ECDSA message with the original message to verify if the message was signed by the given public key
func VerifySignature(message, publicKey, signature []byte) error {
	err := verifySignInputs(message, publicKey, signature, false, true)
	if err != nil {
		return fmt.Errorf("invalid inputs: %s", err)
	}

	var isVerified uint32
	err0 := ethereumSym.EthVerifySignature(&message[0], uint32(len(message)), &publicKey[0], uint32(len(publicKey)), &signature[0], &isVerified)
	if err0 != 0 {
		return fmt.Errorf("verifying signature failed with: %s", err0)
	}

	if isVerified == 0 {
		return fmt.Errorf("signature not signed by this private key")
	}

	return nil
}

func verifySignInputs(message, key, signature []byte, keyPrivate, checkSig bool) error {
	if len(message) == 0 {
		return fmt.Errorf("message is empty")
	}

	if len(key) == 0 {
		keyType := "public"
		if keyPrivate {
			keyType = "private"
		}
		return fmt.Errorf("%s Key is empty", keyType)
	}

	if checkSig == true && (len(signature) == 0 || signature == nil) {
		return fmt.Errorf("signature is nil")
	}

	return nil
}
