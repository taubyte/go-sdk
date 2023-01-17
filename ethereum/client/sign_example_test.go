package ethereum_test

import (
	"fmt"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

func ExampleSign() {
	exampleMessage := []byte("hello world")
	examplePrivateKey := []byte("fake_private_key")
	// Mocking the calls to the vm for usage in tests and playground
	ethereumSym.MockSign(exampleMessage)
	ethereumSym.MockVerify(exampleMessage, true)

	signature, err := ethereum.SignMessage(exampleMessage, examplePrivateKey)
	if err != nil {
		return
	}

	err = ethereum.VerifySignature(exampleMessage, examplePrivateKey, signature)
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}
