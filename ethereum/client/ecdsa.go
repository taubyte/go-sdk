package ethereum

import (
	"encoding/hex"
	"fmt"
	"strconv"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/ecdsa"
	"golang.org/x/crypto/sha3"
)

// HexToECDSABytes returns the ECDSA []byte format of the given hex string representation of a Private Key.
func HexToECDSABytes(hexString string) ([]byte, error) {
	privKeyBytes := make([]byte, EcdsaPrivKeySize)
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

// ParseSignature parses a hex string signature to bytes
func ParseSignature(signature string) ([]byte, error) {
	signHex := []byte(signature)

	if len(signHex) >= 2 && (signHex[0] == '0' && (signHex[1] == 'x' || signHex[1] == 'X')) {
		signHex = signHex[2:]
	}

	sign := make([]byte, hex.DecodedLen(len(signHex)))
	_, err := hex.Decode(sign, signHex)
	if err != nil {
		return nil, fmt.Errorf("decoding signature `%s` failed with: %s", signature, err)
	}

	if sign[64] >= 27 {
		sign[64] -= 27
	}

	return sign, nil
}

// ToEthJsMessage returns the signed message in the JS Ethereum message format as bytes.
func ToEthJsMessage(message string) []byte {
	return []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
}

// AddressFromPubKey returns the Address for the given public key.
func AddressFromPubKey(pubKey []byte) Address {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKey[1:]) // remove EC prefix 04
	buf := hash.Sum(nil)[12:]
	var address Address

	copy(address[:], buf)
	return address
}

// String returns the 0x prefixed hex string representation of the address.
func (a Address) String() string {
	return "0x" + hex.EncodeToString(a[:])
}

// Bytes returns the address as []byte representation.
func (a Address) Bytes() []byte {
	return a[:]
}
