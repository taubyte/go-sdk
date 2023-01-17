package ethereum

import (
	"bytes"
	"io"
	"math/rand"
	"os"
	"testing"
	"unsafe"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	"github.com/taubyte/go-sdk/errno"
)

func TestGetContract(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	client, err := newMockClient()
	if err != nil {
		t.Error(err)
		return
	}
	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	contract, err := client.getContract(testAddress, testContractId, testContractMethodSize)
	if err != nil {
		t.Errorf("Getting contract failed with: %s", err)
		return
	}

	if contract.id != testContractId {
		t.Errorf("Expected contract id `%d` got `%d`", testContractId, contract.id)
		return
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, true)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.getContract(testAddress, testContractId, testContractMethodSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractId, true, false)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.getContract(testAddress, testContractId, testContractMethodSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.MethodDataClientId = testClientId + 10
	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.getContract(testAddress, testContractId, testContractMethodSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.Methods[testPassingMethod] = ethereumSym.MockContractMethod{
		Inputs:  []interface{}{testPassingInput},
		Outputs: []interface{}{},
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.getContract(testAddress, testContractId, testContractMethodSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.Methods[testPassingMethod] = ethereumSym.MockContractMethod{
		Inputs:  []interface{}{},
		Outputs: []interface{}{},
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.getContract(testAddress, testContractId, testContractMethodSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.MethodSizeClientId = testClientId + 10
	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.getContract(testAddress, testContractId, testContractMethodSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.ContractDataClientId = testClientId + 10
	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.getContract(testAddress, testContractId, testContractMethodSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.EthNewContract = func(clientId, contractId uint32, methodsPtr *byte) (error errno.Error) {
		d := unsafe.Slice(methodsPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}

	_, err = client.getContract(testAddress, testContractId, 11)
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.EthNewContract = func(clientId, contractId uint32, methodsPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = client.getContract(testAddress, testContractId, testContractMethodSize)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestBoundContract(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	client, err := newMockClient()
	if err != nil {
		t.Error(err)
		return
	}

	reader, err := os.Open("contract_test.go")
	if err != nil {
		t.Error(err)
		return
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.NewBoundContract(reader, testAddress)
	if err != nil {
		t.Errorf("New bound contract failed with: %s", err)
		return
	}

	reader.Close()
	_, err = client.NewBoundContract(reader, testAddress)
	if err == nil {
		t.Error("Expected read error")
		return
	}

	testContract.ContractSizeClientId = testClientId + 10
	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	reader, err = os.Open("contract_test.go")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.NewBoundContract(reader, testAddress)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = client.NewBoundContract(nil, testAddress)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = client.NewBoundContract(reader, "")
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestDeployContract(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	client, err := newMockClient()
	if err != nil {
		t.Error(err)
		return
	}

	reader, err := os.Open("contract_test.go")
	if err != nil {
		t.Error(err)
		return
	}

	err = ethereumSym.MockDeployContract(testContract, testAddress, testTransactionId, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	byteCode := make([]byte, 1024)
	rand.Read(byteCode)
	_byteCode := bytes.NewReader(byteCode)

	contract, tx, err := client.DeployContract(reader, _byteCode, testChain, testBytes)
	if err != nil {
		t.Errorf("Deploying contract failed with: %s", err)
		return
	}
	if contract.id != testContractId {
		t.Errorf("Expected contract id `%d` got `%d`", testContractId, contract.id)
		return
	}

	if contract.Address() != testAddress {
		t.Errorf("Expected testAddress `%s` got `%s`", contract.Address(), testAddress)
		return
	}

	if tx.id != testTransactionId {
		t.Errorf("Expected transaction id `%d` got `%d`", testTransactionId, tx.id)
	}

	for _, method := range contract.Methods() {
		method.Name()
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractId, false, true)
	if err != nil {
		t.Error(err)
		return
	}

	reader, err = os.Open("contract_test.go")
	if err != nil {
		t.Error(err)
		return
	}

	_byteCode.Seek(0, io.SeekStart)

	_, _, err = client.DeployContract(reader, _byteCode, testChain, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.ContractSizeClientId = testClientId + 10
	err = ethereumSym.MockDeployContract(testContract, testAddress, testTransactionId, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	reader, err = os.Open("contract_test.go")
	if err != nil {
		t.Error(err)
		return
	}

	_byteCode.Seek(0, io.SeekStart)

	_, _, err = client.DeployContract(reader, _byteCode, testChain, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	reader.Close()

	_byteCode.Seek(0, io.SeekStart)

	_, _, err = client.DeployContract(reader, _byteCode, testChain, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.ContractDataClientId = testClientId + 10
	err = ethereumSym.MockDeployContract(testContract, testAddress, testTransactionId, testContractId, false, false)
	if err != nil {
		t.Error(err)
		return
	}

	reader, err = os.Open("contract_test.go")
	if err != nil {
		t.Error(err)
		return
	}

	reader2, err := os.Open("contract_test.go")
	if err != nil {
		t.Error(err)
		return
	}

	_, _, err = client.DeployContract(reader, reader2, testChain, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	reader, err = os.Open("contract_test.go")
	if err != nil {
		t.Error(err)
		return
	}

	reader2.Close()

	_, _, err = client.DeployContract(reader, reader2, testChain, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, _, err = client.DeployContract(nil, _byteCode, testChain, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_byteCode.Seek(0, io.SeekStart)

	_, _, err = client.DeployContract(reader, _byteCode, nil, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockCurrentChainId(testClientId+10, testClientId, testChain)
	if err != nil {
		t.Error(err)
		return
	}

	_byteCode.Seek(0, io.SeekStart)

	_, _, err = client.DeployContract(reader, _byteCode, nil, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_byteCode.Seek(0, io.SeekStart)

	_, _, err = client.DeployContract(reader, _byteCode, nil, nil)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, _, err = client.DeployContract(reader, nil, nil, nil)
	if err == nil {
		t.Error("Expected error")
		return
	}
}
