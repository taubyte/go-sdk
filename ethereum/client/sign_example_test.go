package ethereum_test

import (
	"fmt"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

func ExampleSign() {
	// Mocking the calls to the vm for usage in tests and playground
	ethereumSym.MockSign("hello world")
	ethereumSym.MockVerify("hello world", true)

	signature, err := ethereum.SignMessage("hello world", "fake_private_key")
	if err != nil {
		return
	}

	err = ethereum.VerifySignature("hello world", "fake_private_key", signature)
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}
