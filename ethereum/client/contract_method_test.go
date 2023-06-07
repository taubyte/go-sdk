package ethereum

import (
	"bytes"
	"math/rand"
	"os"
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	"gotest.tools/assert"
)

func TestContractMethod(t *testing.T) {
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

	contract, _, err := client.DeployContract(reader, _byteCode, testChain, testBytes)
	assert.NilError(t, err)

	for method := range testContract.Methods {
		contractMethod, err := contract.Method(method)
		assert.NilError(t, err)

		if contractMethod.Name() != method {
			t.Errorf("Expected method name `%s` got `%s`", method, contractMethod.Name())
			return
		}

		_, err = contract.Method(method + testString)
		if err == nil {
			t.Errorf("Expected error")
			return
		}
	}
}

func TestInputParams(t *testing.T) {
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

	contract, _, err := client.DeployContract(reader, _byteCode, testChain, testBytes)
	assert.NilError(t, err)

	contractMethod, err := contract.Method(testPassingMethod)
	assert.NilError(t, err)

	inputBytes, err := contractMethod.handleInputs()
	assert.NilError(t, err)

	if !bytes.Equal(inputBytes, []byte{0}) {
		t.Error("Unexpected byte array")
		return
	}

	_, err = contractMethod.handleInputs(testPassingInput)
	assert.NilError(t, err)

	_, err = contractMethod.handleInputs(testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	contractMethod, err = contract.Method(testInputFailureMethod)
	assert.NilError(t, err)

	_, err = contractMethod.handleInputs(testIncompatibleVar)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestContractTransact(t *testing.T) {
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

	contract, _, err := client.DeployContract(reader, _byteCode, testChain, testBytes)
	assert.NilError(t, err)

	contractMethod, err := contract.Method(testPassingMethod)
	assert.NilError(t, err)

	ethereumSym.MockTransactContract(testClientID, testTransactionID)

	tx, err := contractMethod.Transact(testChain, testBytes, testPassingInput)
	assert.NilError(t, err)

	if tx.id != testTransactionID {
		t.Errorf("Expected transaction id `%d` got `%d`", testTransactionID, tx.id)
		return
	}

	ethereumSym.MockTransactContract(testClientID+10, testTransactionID)

	_, err = contractMethod.Transact(testChain, testBytes, testPassingInput)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = contractMethod.Transact(testChain, testBytes, testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = contractMethod.Transact(nil, testBytes, testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = contractMethod.Transact(nil, nil, testString)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestCall(t *testing.T) {
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

	contract, _, err := client.DeployContract(reader, _byteCode, testChain, testBytes)
	assert.NilError(t, err)

	contractMethod, err := contract.Method(testPassingMethod)
	assert.NilError(t, err)

	ethereumSym.MockCall(testContract, false, false)
	assert.NilError(t, err)

	_, err = contractMethod.Call(testPassingInput)
	assert.NilError(t, err)

	contractMethod, err = contract.Method(testInputFailureMethod)
	assert.NilError(t, err)

	_, err = contractMethod.Call(testPassingInput)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = contractMethod.Call(testIncompatibleVar)
	if err == nil {
		t.Error("Expected error")
		return
	}

	contractMethod, err = contract.Method(testOutputFailureMethod)
	assert.NilError(t, err)

	_, err = contractMethod.Call(testIncompatibleVar)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = contractMethod.Call(testPassingInput)
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockCall(testContract, true, false)

	contractMethod, err = contract.Method(testPassingMethod)
	assert.NilError(t, err)

	_, err = contractMethod.Call(testPassingInput)
	if err == nil {
		t.Error("expected error")
		return
	}

	contractMethod, err = contract.Method(testOutputFailureMethod)
	assert.NilError(t, err)

	_, err = contractMethod.Call(testPassingInput)
	if err == nil {
		t.Error("expected error")
		return
	}

	ethereumSym.MockCall(testContract, false, true)

	_, err = contractMethod.Call(testPassingInput)
	if err == nil {
		t.Error("expected error")
		return
	}

	contractMethod, err = contract.Method(testPassingMethod)
	assert.NilError(t, err)

	testContract.CallDataClientId = testClientID + 10
	ethereumSym.MockCall(testContract, false, false)

	_, err = contractMethod.Call(testPassingInput)
	if err == nil {
		t.Error("expected error")
		return
	}

	testContract.Methods = map[string]ethereumSym.MockContractMethod{}
	ethereumSym.MockCall(testContract, false, false)

	output, err := contractMethod.Call(testPassingInput)
	assert.NilError(t, err)

	if output != nil {
		t.Errorf("Expected nil output, got `%v`", output)
	}
}
