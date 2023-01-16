package ethereum

import (
	"fmt"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

// HexToECDSABytes returns the ECDSA []byte format of the given hex string representation of a Private Key.
func HexToECDSABytes(hexString string) ([]byte, error) {
	var size uint32
	if err0 := ethereumSym.EthHexToECDSASize(hexString, &size); err0 != 0 {
		return nil, fmt.Errorf("converting hex string `%s` to ECDSA formatted bytes failed with: %s", hexString, err0)
	}

	privKeyBytes := make([]byte, size)
	if err0 := ethereumSym.EthHexToECDSA(hexString, &privKeyBytes[0]); err0 != 0 {
		return nil, fmt.Errorf("converted hex string `%s` to ECDSA formatted bytes, but unable to read bytes with: %s", hexString, err0)
	}

	return privKeyBytes, nil
}

// PublicKeyFromPrivate returns the uncompressed ECDSA public key from the given privateKey,
func PublicKeyFromPrivate(privateKey []byte) ([]byte, error) {
	publicKey := make([]byte, EcdsaPubKeySize)
	if err0 := ethereumSym.EthPubFromPriv(&privateKey[0], uint32(len(privateKey)), &publicKey[0]); err0 != 0 {
		return nil, fmt.Errorf("getting public key from private key failed with: %s", err0)
	}

	return publicKey, nil
}

// PublicKeyFromSignedMessage returns the ECDSA public key used to sign the given message to the given signature.
func PublicKeyFromSignedMessage(message []byte, signature []byte) ([]byte, error) {
	publicKey := make([]byte, EcdsaPubKeySize)
	if err0 := ethereumSym.EthPubKeyFromSignedMessage(&message[0], uint32(len(message)), &signature[0], uint32(len(signature)), &publicKey[0]); err0 != 0 {
		return nil, fmt.Errorf("getting public key from message and signature failed with: %s", err0)
	}

	return publicKey, nil
}
