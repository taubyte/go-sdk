package ethereum_test

import (
	"fmt"
	"math/big"
	"math/rand"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

var tx *ethereum.Transaction

func ExampleBlock_Transaction() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
	}
	mockData.Mock()

	client, err := ethereum.New("https://testRPC.url")
	if err != nil {
		return
	}

	block, err := client.BlockByNumber(big.NewInt(20))
	if err != nil {
		return
	}

	// Mocking transaction hash
	txHash := make([]byte, 32)
	_, err = rand.Read(txHash)
	if err != nil {
		return
	}

	tx, err = block.Transaction(txHash)
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}

func ExampleBlock_Transactions() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:            4,
		BlockByNumber:     5,
		BlockTransactions: []uint32{1, 2, 3},
	}
	mockData.Mock()

	client, err := ethereum.New("https://testRPC.url")
	if err != nil {
		return
	}

	block, err := client.BlockByNumber(big.NewInt(20))
	if err != nil {
		return
	}

	txs, err := block.Transactions()
	if err != nil {
		return
	}

	fmt.Println(len(txs))
	// Output: 3
}

func ExampleBlock_Number() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:        4,
		BlockByNumber: 5,
		BlockNumber:   big.NewInt(20),
	}
	mockData.Mock()

	client, err := ethereum.New("https://testRPC.url")
	if err != nil {
		return
	}

	block, err := client.BlockByNumber(big.NewInt(20))
	if err != nil {
		return
	}

	blockNumber, err := block.Number()
	if err != nil {
		return
	}

	fmt.Printf("%d\n", blockNumber)
	// Output: 20
}
