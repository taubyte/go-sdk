package ethereum_test

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

func ExampleTransaction_Nonce() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		TransactionU64:   7,
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	nonce, err := tx.Nonce()
	if err != nil {
		return
	}

	fmt.Println(nonce)
	// Output: 7
}

func ExampleTransaction_GasPrice() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		TransactionBytes: big.NewInt(7).Bytes(),
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	gasPrice, err := tx.GasPrice()
	if err != nil {
		return
	}

	fmt.Println(gasPrice)
	// Output: 7
}

func ExampleTransaction_GasTipCap() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		TransactionBytes: big.NewInt(7).Bytes(),
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	gasTipCap, err := tx.GasTipCap()
	if err != nil {
		return
	}

	fmt.Println(gasTipCap)
	// Output: 7
}

func ExampleTransaction_GasFeeCap() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		TransactionBytes: big.NewInt(7).Bytes(),
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	gasFeeCap, err := tx.GasFeeCap()
	if err != nil {
		return
	}

	fmt.Println(gasFeeCap)
	// Output: 7
}

func ExampleTransaction_Gas() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		TransactionU64:   7,
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	gas, err := tx.Gas()
	if err != nil {
		return
	}

	fmt.Println(gas)
	// Output: 7
}

func ExampleTransaction_Value() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		TransactionBytes: big.NewInt(7).Bytes(),
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	value, err := tx.Value()
	if err != nil {
		return
	}

	fmt.Println(value)
	// Output: 7
}

func ExampleTransaction_Data() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		TransactionBytes: []byte("Hello World"),
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	data, err := tx.Data()
	if err != nil {
		return
	}

	fmt.Println(string(data))
	// Output: Hello World
}

func ExampleTransaction_RawSignatures() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		VSig:             big.NewInt(7),
		RSig:             big.NewInt(8),
		SSig:             big.NewInt(9),
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	rawSignatures, err := tx.RawSignatures()
	if err != nil {
		return
	}

	fmt.Println(rawSignatures.VSig, rawSignatures.RSig, rawSignatures.SSig)
	// Output: 7 8 9
}

func ExampleTransaction_Chain() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		TransactionBytes: big.NewInt(7).Bytes(),
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	chain, err := tx.Chain()
	if err != nil {
		return
	}

	fmt.Println(chain)
	// Output: 7
}

func ExampleTransaction_Hash() {
	// Mocking transaction hash
	txHash := make([]byte, 32)
	_, err := rand.Read(txHash)
	if err != nil {
		return
	}

	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:           4,
		BlockByNumber:    5,
		BlockTransaction: 6,
		TransactionBytes: txHash,
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	hash, err := tx.Hash()
	if err != nil || !bytes.Equal(hash, txHash) {
		return
	}

	fmt.Println("success")
	// Output: success
}

func ExampleTransaction_Send() {
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

	tx, err := block.Transaction(txHash)
	if err != nil {
		return
	}

	err = tx.Send()
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}
