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
	"gotest.tools/assert"
)

func TestGetContract(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	client, err := newMockClient()
	assert.NilError(t, err)

	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, false)
	assert.NilError(t, err)

	contract, err := client.getContract(testAddress, testContractID, testContractMethodSize, testContractEventsSize)
	assert.NilError(t, err)

	if contract.id != testContractID {
		t.Errorf("Expected contract id `%d` got `%d`", testContractID, contract.id)
		return
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, true)
	assert.NilError(t, err)

	_, err = client.getContract(testAddress, testContractID, testContractMethodSize, testContractEventsSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractID, true, false)
	assert.NilError(t, err)

	_, err = client.getContract(testAddress, testContractID, testContractMethodSize, testContractEventsSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.MethodDataClientId = testClientID + 10
	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, false)
	assert.NilError(t, err)

	_, err = client.getContract(testAddress, testContractID, testContractMethodSize, testContractEventsSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.Methods[testPassingMethod] = ethereumSym.MockContractMethod{
		Inputs:  []interface{}{testPassingInput},
		Outputs: []interface{}{},
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, false)
	assert.NilError(t, err)

	_, err = client.getContract(testAddress, testContractID, testContractMethodSize, testContractEventsSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.Methods[testPassingMethod] = ethereumSym.MockContractMethod{
		Inputs:  []interface{}{},
		Outputs: []interface{}{},
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, false)
	assert.NilError(t, err)

	_, err = client.getContract(testAddress, testContractID, testContractMethodSize, testContractEventsSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.MethodSizeClientId = testClientID + 10
	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, false)
	assert.NilError(t, err)

	_, err = client.getContract(testAddress, testContractID, testContractMethodSize, testContractEventsSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.ContractDataClientId = testClientID + 10
	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, false)
	assert.NilError(t, err)

	_, err = client.getContract(testAddress, testContractID, testContractMethodSize, testContractEventsSize)
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.EthNewContract = func(clientId, contractId uint32, methodsPtr, eventsPtr *byte) (error errno.Error) {
		d := unsafe.Slice(methodsPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}
	_, err = client.getContract(testAddress, testContractID, 11, 12)
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.EthNewContract = func(clientId, contractId uint32, methodsPtr, eventsPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = client.getContract(testAddress, testContractID, testContractMethodSize, testContractEventsSize)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestBoundContract(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	client, err := newMockClient()
	assert.NilError(t, err)

	reader, err := os.Open("contract_test.go")
	assert.NilError(t, err)

	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, false)
	assert.NilError(t, err)

	_, err = client.NewBoundContract(reader, testAddress)
	assert.NilError(t, err)

	reader.Close()
	_, err = client.NewBoundContract(reader, testAddress)
	if err == nil {
		t.Error("Expected read error")
		return
	}

	testContract.ContractSizeClientId = testClientID + 10
	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, false)
	assert.NilError(t, err)

	reader, err = os.Open("contract_test.go")
	assert.NilError(t, err)

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
	assert.NilError(t, err)

	client, err := newMockClient()
	assert.NilError(t, err)

	reader, err := os.Open("contract_test.go")
	assert.NilError(t, err)

	err = ethereumSym.MockDeployContract(testContract, testAddress, testTransactionID, testContractID, false, false)
	assert.NilError(t, err)

	byteCode := make([]byte, 1024)
	rand.Read(byteCode)
	_byteCode := bytes.NewReader(byteCode)

	contract, tx, err := client.DeployContract(reader, _byteCode, testChain, testBytes)
	assert.NilError(t, err)

	if contract.id != testContractID {
		t.Errorf("Expected contract id `%d` got `%d`", testContractID, contract.id)
		return
	}

	if contract.Address() != testAddress {
		t.Errorf("Expected testAddress `%s` got `%s`", contract.Address(), testAddress)
		return
	}

	if tx.id != testTransactionID {
		t.Errorf("Expected transaction id `%d` got `%d`", testTransactionID, tx.id)
	}

	for _, method := range contract.Methods() {
		method.Name()
	}

	err = ethereumSym.MockNewBoundContract(testContract, testContractID, false, true)
	assert.NilError(t, err)

	reader, err = os.Open("contract_test.go")
	assert.NilError(t, err)

	_byteCode.Seek(0, io.SeekStart)

	_, _, err = client.DeployContract(reader, _byteCode, testChain, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	testContract.ContractSizeClientId = testClientID + 10
	err = ethereumSym.MockDeployContract(testContract, testAddress, testTransactionID, testContractID, false, false)
	assert.NilError(t, err)

	reader, err = os.Open("contract_test.go")
	assert.NilError(t, err)

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

	testContract.ContractDataClientId = testClientID + 10
	err = ethereumSym.MockDeployContract(testContract, testAddress, testTransactionID, testContractID, false, false)
	assert.NilError(t, err)

	reader, err = os.Open("contract_test.go")
	assert.NilError(t, err)

	reader2, err := os.Open("contract_test.go")
	assert.NilError(t, err)

	_, _, err = client.DeployContract(reader, reader2, testChain, testBytes)
	if err == nil {
		t.Error("Expected error")
		return
	}

	reader, err = os.Open("contract_test.go")
	assert.NilError(t, err)

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

	err = ethereumSym.MockCurrentChainId(testClientID+10, testClientID, testChain)
	assert.NilError(t, err)

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
