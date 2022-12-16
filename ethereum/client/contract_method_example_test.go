package ethereum_test

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

func ExampleContractMethod_Call() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:          4,
		ContractAddress: "address",
		Contract: map[string]ethereumSym.MockContractMethod{
			"fakeMethod": {
				Inputs:  []interface{}{big.NewInt(5)},
				Outputs: []interface{}{big.NewInt(6)},
			},
		},
		ContractTransactionId: 2,
		ContractId:            3,
	}
	mockData.Mock()

	client, err := ethereum.New("https://testRPC.url")
	if err != nil {
		return
	}

	// Mocking abi data
	abiRawData := make([]byte, 1024)
	rand.Read(abiRawData)

	contract, err := client.NewBoundContract(bytes.NewReader(abiRawData), "address")
	if err != nil {
		return
	}

	fakeMethod, err := contract.Method("fakeMethod")
	if err != nil {
		return
	}

	outputs, err := fakeMethod.Call(big.NewInt(5))
	if err != nil || len(outputs) != 1 {
		return
	}

	fmt.Println(outputs[0])
	// Output: 6
}

func ExampleContractMethod_Transact() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:          4,
		ContractAddress: "address",
		Contract: map[string]ethereumSym.MockContractMethod{
			"fakeMethod": {
				Inputs:  []interface{}{big.NewInt(5)},
				Outputs: []interface{}{big.NewInt(6)},
			},
		},
		ContractTransactionId: 2,
		ContractId:            3,
	}
	mockData.Mock()

	// Creates new client from given RPC url, this is not a real rpc url.
	client, err := ethereum.New("https://testRPC.url")
	if err != nil {
		return
	}

	// Mocking abi data
	abiRawData := make([]byte, 1024)
	rand.Read(abiRawData)

	contract, err := client.NewBoundContract(bytes.NewReader(abiRawData), "address")
	if err != nil {
		return
	}

	fakeMethod, err := contract.Method("fakeMethod")
	if err != nil {
		return
	}

	chainId := big.NewInt(10)

	tx, err = fakeMethod.Transact(chainId, "privateKey", big.NewInt(5))
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}
