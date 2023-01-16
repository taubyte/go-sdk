package ethereum_test

import (
	"fmt"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

func ExampleHexToECDSABytes() {
	// Mocking the calls to the vm for usage in tests and playground
	privateKeyHex := "private_key"
	privateKeyBytes := []byte(privateKeyHex)
	ethereumSym.MockHexToECDSA(privateKeyHex, privateKeyBytes, uint32(len(privateKeyBytes)))

	privateKey, err := ethereum.HexToECDSABytes(privateKeyHex)
	if err != nil {
		return
	}

	fmt.Println(string(privateKey))
	// Output: private_key
}

var publicKey []byte
var err error

func ExampleTestPublicKeyFromPrivate() {
	// Mocking the calls to the vm for usage in tests and playground
	privateKeyBytes := []byte("private_key")
	ethereumSym.MockPublicKeyFromPrivate(privateKeyBytes)

	publicKey, err = ethereum.PublicKeyFromPrivate(privateKeyBytes)
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}

func ExampleTestPublicKeyFromSignedMessage() {
	// Mocking the calls to the vm for usage in tests and playground
	signature := []byte("signed_message")
	message := []byte("message")
	ethereumSym.MockPublicKeyFromSignedMessage(message)

	publicKey, err = ethereum.PublicKeyFromSignedMessage(message, signature)
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}
