package ethereum_test

import (
	"fmt"

	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

var privateKey []byte
var publicKey []byte
var err error

func ExampleHexToECDSABytes() {
	privateKey, err = ethereum.HexToECDSABytes("91e4a13e5a30ad353cdf5ea7bb909dfdf967122e3b43e331ad947b68a3899b2c")
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}

func ExampleTestPublicKeyFromPrivate() {
	privateKey, err := ethereum.HexToECDSABytes("91e4a13e5a30ad353cdf5ea7bb909dfdf967122e3b43e331ad947b68a3899b2c")
	if err != nil {
		return
	}

	publicKey, err = ethereum.PublicKeyFromPrivate(privateKey)
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}

func ExampleTestPublicKeyFromSignedMessage() {
	privateKey, err := ethereum.HexToECDSABytes("91e4a13e5a30ad353cdf5ea7bb909dfdf967122e3b43e331ad947b68a3899b2c")
	if err != nil {
		return
	}

	signature, err := ethereum.SignMessage([]byte("hello world"), privateKey)
	if err != nil {
		return
	}

	publicKey, err = ethereum.PublicKeyFromSignedMessage([]byte("hello world"), signature)
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}
