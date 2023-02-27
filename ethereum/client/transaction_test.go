package ethereum

import (
	"bytes"
	"math/big"
	"math/rand"
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	"gotest.tools/assert"
)

func TestTransactionNonce(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	expectedU16 := rand.Uint64()
	ethereumSym.MockU64method(testClientID+10, expectedU16)

	nonce, err := tx.Nonce()
	if err == nil {
		t.Errorf("expected error")
		return
	}

	ethereumSym.MockU64method(testClientID, expectedU16)
	nonce, err = tx.Nonce()
	assert.NilError(t, err)

	if nonce != expectedU16 {
		t.Errorf("Expected nonce `%d` got `%d`", expectedU16, nonce)
		return
	}

	nonce, err = tx.Nonce()
	if err != nil {
		t.Errorf("Getting nonce failed with: %s", err)
		return
	}
	if nonce != expectedU16 {
		t.Errorf("Expected nonce `%d` got `%d`", expectedU16, nonce)
		return
	}

}

func TestBytesMethod(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	bigInt := big.NewInt(rand.Int63())
	ethereumSym.MockBytesMethod(testClientID, testClientID+10, bigInt.Bytes())

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	_, err = tx.callBytesMethod("")
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID+10, big.NewInt(0).Bytes())

	data, err := tx.callBytesMethod("")
	assert.NilError(t, err)
	if data != nil {
		t.Errorf("expected nil data")
		return
	}

	ethereumSym.MockBytesMethod(testClientID+10, testClientID, big.NewInt(0).Bytes())

	_, err = tx.callBytesMethod("")
	if err == nil {
		t.Error("expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID, bigInt.Bytes())
	data, err = tx.callBytesMethod("")
	assert.NilError(t, err)

	if bytes.Compare(data, bigInt.Bytes()) != 0 {
		t.Error("Sent and received data are not the same")
		return
	}
}

func TestGasPrice(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	testGasPrice := big.NewInt(rand.Int63())

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	ethereumSym.MockBytesMethod(testClientID+10, testClientID, testGasPrice.Bytes())

	gasPrice, err := tx.GasPrice()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID, testGasPrice.Bytes())

	gasPrice, err = tx.GasPrice()
	assert.NilError(t, err)

	if gasPrice.Cmp(testGasPrice) != 0 {
		t.Errorf("Expected gas price `%d` got `%d`", testGasPrice, gasPrice)
		return
	}

	gasPrice, err = tx.GasPrice()
	assert.NilError(t, err)

	if gasPrice.Cmp(testGasPrice) != 0 {
		t.Errorf("Expected gas price `%d` got `%d`", testGasPrice, gasPrice)
		return
	}

}

func TestGasTipCap(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	testGasTipCap := big.NewInt(rand.Int63())

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	ethereumSym.MockBytesMethod(testClientID+10, testClientID, testGasTipCap.Bytes())

	gasTipCap, err := tx.GasTipCap()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID, testGasTipCap.Bytes())

	gasTipCap, err = tx.GasTipCap()
	assert.NilError(t, err)

	if gasTipCap.Cmp(testGasTipCap) != 0 {
		t.Errorf("Expected gas tip cap `%d` got `%d`", testGasTipCap, gasTipCap)
		return
	}

	gasTipCap, err = tx.GasTipCap()
	assert.NilError(t, err)

	if gasTipCap.Cmp(testGasTipCap) != 0 {
		t.Errorf("Expected gas tip cap `%d` got `%d`", testGasTipCap, gasTipCap)
		return
	}

}

func TestGasFeeCap(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)
	testGasFeeCap := big.NewInt(rand.Int63())

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	ethereumSym.MockBytesMethod(testClientID+10, testClientID, testGasFeeCap.Bytes())

	gasFeeCap, err := tx.GasFeeCap()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID, testGasFeeCap.Bytes())

	gasFeeCap, err = tx.GasFeeCap()
	assert.NilError(t, err)

	if gasFeeCap.Cmp(testGasFeeCap) != 0 {
		t.Errorf("Expected gas tip cap `%d` got `%d`", testGasFeeCap, gasFeeCap)
		return
	}

	gasFeeCap, err = tx.GasFeeCap()
	assert.NilError(t, err)

	if gasFeeCap.Cmp(testGasFeeCap) != 0 {
		t.Errorf("Expected gas tip cap `%d` got `%d`", testGasFeeCap, gasFeeCap)
		return
	}

}

func TestTransactionGas(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	ethereumSym.MockU64method(testClientID+10, 5)

	gas, err := tx.Gas()
	if err == nil {
		t.Errorf("expected error")
		return
	}

	expectedU64 := rand.Uint64()
	ethereumSym.MockU64method(testClientID, expectedU64)
	gas, err = tx.Gas()
	assert.NilError(t, err)
	if gas != expectedU64 {
		t.Errorf("Expected gas `%d` got `%d`", expectedU64, gas)
		return
	}

	gas, err = tx.Gas()
	assert.NilError(t, err)
	if gas != expectedU64 {
		t.Errorf("Expected gas `%d` got `%d`", expectedU64, gas)
		return
	}

}

func TestValue(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	testValue := big.NewInt(rand.Int63())

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	ethereumSym.MockBytesMethod(testClientID+10, testClientID, testValue.Bytes())

	value, err := tx.Value()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID, testValue.Bytes())

	value, err = tx.Value()
	assert.NilError(t, err)

	if value.Cmp(testValue) != 0 {
		t.Errorf("Expected value `%d` got `%d`", testValue, value)
		return
	}

	value, err = tx.Value()
	assert.NilError(t, err)

	if value.Cmp(testValue) != 0 {
		t.Errorf("Expected value `%d` got `%d`", testValue, value)
		return
	}

}

func TestData(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	testBytes := make([]byte, 1)
	rand.Read(testBytes)

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	ethereumSym.MockBytesMethod(testClientID+10, testClientID, testBytes)

	data, err := tx.Data()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID, testBytes)

	data, err = tx.Data()
	assert.NilError(t, err)

	if bytes.Compare(data, testBytes) != 0 {
		t.Error("sent data nad received data are not the same")
		return
	}

	data, err = tx.Data()
	assert.NilError(t, err)

	if bytes.Compare(data, testBytes) != 0 {
		t.Error("sent data nad received data are not the same")
		return
	}
}

func TestAddress(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	ethereumSym.MockBytesMethod(testClientID+10, testClientID+10, testAddressBytes)

	address, err := tx.ToAddress()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID, testAddressBytes)

	address, err = tx.ToAddress()
	assert.NilError(t, err)

	if bytes.Compare(address, testAddressBytes) != 0 {
		t.Error("sent address and received address are not the same")
		return
	}

	address, err = tx.ToAddress()
	assert.NilError(t, err)

	if bytes.Compare(address, testAddressBytes) != 0 {
		t.Error("sent address and received address are not the same")
		return
	}
}

func TestChain(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	ethereumSym.MockBytesMethod(testClientID+10, testClientID+10, testChain.Bytes())

	chain, err := tx.Chain()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID, testChain.Bytes())

	chain, err = tx.Chain()
	assert.NilError(t, err)

	if chain.Cmp(testChain) != 0 {
		t.Errorf("Expected chain `%d` got `%d`", testChain, chain)
		return
	}

	chain, err = tx.Chain()
	assert.NilError(t, err)

	if chain.Cmp(testChain) != 0 {
		t.Errorf("Expected chain `%d` got `%d`", testChain, chain)
		return
	}
}

func TestHash(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	testBytes := make([]byte, 1)
	rand.Read(testBytes)

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	tx.hash = nil

	ethereumSym.MockBytesMethod(testClientID+10, testClientID+10, testBytes)
	hash, err := tx.Hash()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientID, testClientID, testBytes)

	hash, err = tx.Hash()
	assert.NilError(t, err)

	if bytes.Compare(hash, testBytes) != 0 {
		t.Error("sent hash and received hash are not the same")
		return
	}

	hash, err = tx.Hash()
	assert.NilError(t, err)

	if bytes.Compare(hash, testBytes) != 0 {
		t.Error("sent hash and received data are not the same")
		return
	}
}

func TestSend(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	ethereumSym.MockSendTransaction(testClientID)

	err = tx.Send()
	assert.NilError(t, err)

	ethereumSym.MockSendTransaction(testClientID + 10)

	err = tx.Send()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestRawSignatures(t *testing.T) {
	tx, _, _, err := newMockTransaction()
	assert.NilError(t, err)

	rSig := big.NewInt(rand.Int63())
	vSig := big.NewInt(rand.Int63())
	sSig := big.NewInt(rand.Int63())

	ethereumSym.MockRawSignatures(testClientID, testClientID+10, vSig, rSig, sSig)

	_, err = tx.RawSignatures()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockRawSignatures(testClientID+10, testClientID, vSig, rSig, sSig)

	_, err = tx.RawSignatures()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockRawSignatures(testClientID, testClientID, vSig, rSig, sSig)

	rawSigs, err := tx.RawSignatures()
	assert.NilError(t, err)

	if vSig.Cmp(rawSigs.VSig) != 0 {
		t.Error("vsigs are not the same")
		return
	}

	if rSig.Cmp(rawSigs.RSig) != 0 {
		t.Error("rsigs are not the same")
		return
	}

	if sSig.Cmp(rawSigs.SSig) != 0 {
		t.Error("ssigs are not the same")
		return
	}

	rawSigs, err = tx.RawSignatures()
	assert.NilError(t, err)

	if vSig.Cmp(rawSigs.VSig) != 0 {
		t.Error("vsigs are not the same")
		return
	}

	if rSig.Cmp(rawSigs.RSig) != 0 {
		t.Error("rsigs are not the same")
		return
	}

	if sSig.Cmp(rawSigs.SSig) != 0 {
		t.Error("ssigs are not the same")
		return
	}
}
