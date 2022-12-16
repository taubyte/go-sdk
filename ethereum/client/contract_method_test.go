package ethereum

import (
	"bytes"
	"math/rand"
	"os"
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

func TestContractMethod(t *testing.T) {
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

	contract, _, err := client.DeployContract(reader, _byteCode, testChain, testString)
	if err != nil {
		t.Errorf("Deploying contract failed with: %s", err)
		return
	}

	for method := range testContract.Methods {
		contractMethod, err := contract.Method(method)
		if err != nil {
			t.Errorf("Getting contract method `%s` failed with: %s", method, err)
			return
		}

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

	contract, _, err := client.DeployContract(reader, _byteCode, testChain, testString)
	if err != nil {
		t.Errorf("Deploying contract failed with: %s", err)
		return
	}

	contractMethod, err := contract.Method(testPassingMethod)
	if err != nil {
		t.Errorf("Getting contract method `%s` failed with: %s", testPassingMethod, err)
		return
	}

	inputBytes, err := contractMethod.handleInputs()
	if err != nil {
		t.Errorf("Getting input bytes for 0 inputs failed with: %s", err)
		return
	}

	if bytes.Compare(inputBytes, []byte{0}) != 0 {
		t.Error("Unexpected byte array")
		return
	}

	_, err = contractMethod.handleInputs(testPassingInput)
	if err != nil {
		t.Errorf("Getting input bytes contract method failed with: %s", err)
		return
	}

	_, err = contractMethod.handleInputs(testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	contractMethod, err = contract.Method(testInputFailureMethod)
	if err != nil {
		t.Errorf("Getting contract method `%s` failed with: %s", testInputFailureMethod, err)
		return
	}

	_, err = contractMethod.handleInputs(testIncompatibleVar)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestContractTransact(t *testing.T) {
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

	contract, _, err := client.DeployContract(reader, _byteCode, testChain, testString)
	if err != nil {
		t.Errorf("Deploying contract failed with: %s", err)
		return
	}

	contractMethod, err := contract.Method(testPassingMethod)
	if err != nil {
		t.Errorf("Getting contract method `%s` failed with: %s", testPassingMethod, err)
		return
	}

	ethereumSym.MockTransactContract(testClientId, testTransactionId)

	tx, err := contractMethod.Transact(testChain, testString, testPassingInput)
	if err != nil {
		t.Errorf("Calling contract method transaction failed with: %s", err)
		return
	}

	if tx.id != testTransactionId {
		t.Errorf("Expected transaction id `%d` got `%d`", testTransactionId, tx.id)
		return
	}

	ethereumSym.MockTransactContract(testClientId+10, testTransactionId)

	_, err = contractMethod.Transact(testChain, testString, testPassingInput)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = contractMethod.Transact(testChain, testString, testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = contractMethod.Transact(nil, testString, testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = contractMethod.Transact(nil, "", testString)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestCall(t *testing.T) {
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

	contract, _, err := client.DeployContract(reader, _byteCode, testChain, testString)
	if err != nil {
		t.Errorf("Deploying contract failed with: %s", err)
		return
	}

	contractMethod, err := contract.Method(testPassingMethod)
	if err != nil {
		t.Errorf("Getting contract method `%s` failed with: %s", testPassingMethod, err)
		return
	}

	ethereumSym.MockCall(testContract, false, false)
	if err != nil {
		t.Errorf("Initializing mock call failed with: %s", err)
		return
	}

	_, err = contractMethod.Call(testPassingInput)
	if err != nil {
		t.Errorf("Calling contract method `%s` failed with: %s", contractMethod.Name(), err)
		return
	}

	contractMethod, err = contract.Method(testInputFailureMethod)
	if err != nil {
		t.Errorf("Getting contract method `%s` failed with: %s", testInputFailureMethod, err)
		return
	}

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
	if err != nil {
		t.Errorf("Getting contract method `%s` failed with: %s", testOutputFailureMethod, err)
		return
	}

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
	if err != nil {
		t.Errorf("Getting contract method `%s` failed with: %s", testPassingMethod, err)
		return
	}

	_, err = contractMethod.Call(testPassingInput)
	if err == nil {
		t.Error("expected error")
		return
	}

	contractMethod, err = contract.Method(testOutputFailureMethod)
	if err != nil {
		t.Errorf("Getting contract method `%s` failed with: %s", testOutputFailureMethod, err)
		return
	}

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
	if err != nil {
		t.Errorf("Getting contract method `%s` failed with: %s", testOutputFailureMethod, err)
		return
	}

	testContract.CallDataClientId = testClientId + 10
	ethereumSym.MockCall(testContract, false, false)

	_, err = contractMethod.Call(testPassingInput)
	if err == nil {
		t.Error("expected error")
		return
	}

	testContract.Methods = map[string]ethereumSym.MockContractMethod{}
	ethereumSym.MockCall(testContract, false, false)

	output, err := contractMethod.Call(testPassingInput)
	if err != nil {
		t.Errorf("Call failed with: %s", err)
		return
	}
	if output != nil {
		t.Errorf("Expected nil output, got `%v`", output)
	}
}
