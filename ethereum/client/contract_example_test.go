package ethereum_test

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

var (
	contract *ethereum.Contract
)

func ExampleClient_DeployContract() {
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

	// Mocking byte code
	byteCodeData := make([]byte, 1024)
	rand.Read(byteCodeData)

	chainId := big.NewInt(5)

	contract, tx, err = client.DeployContract(bytes.NewReader(abiRawData), bytes.NewReader(byteCodeData), chainId, "private key")
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}

func ExampleClient_NewBoundContract() {
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

	contract, err = client.NewBoundContract(bytes.NewReader(abiRawData), "address")
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}

func ExampleContract_Methods() {
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

	fmt.Println(len(contract.Methods()))
	// Output: 1
}

func ExampleContract_Method() {
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

	fmt.Println(fakeMethod.Name())
	// Output: fakeMethod
}
