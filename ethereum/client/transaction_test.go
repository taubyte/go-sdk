package ethereum

import (
	"bytes"
	"math/big"
	"math/rand"
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

func TestTransactionNonce(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	expectedU16 := rand.Uint64()
	ethereumSym.MockU64method(testClientId+10, expectedU16)

	nonce, err := tx.Nonce()
	if err == nil {
		t.Errorf("expected error")
		return
	}

	ethereumSym.MockU64method(testClientId, expectedU16)
	nonce, err = tx.Nonce()
	if err != nil {
		t.Errorf("Getting nonce failed with: %s", err)
		return
	}
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
	if err != nil {
		t.Error(err)
		return
	}

	bigInt := big.NewInt(rand.Int63())
	ethereumSym.MockBytesMethod(testClientId, testClientId+10, bigInt.Bytes())

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	_, err = tx.callBytesMethod("")
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId+10, big.NewInt(0).Bytes())

	data, err := tx.callBytesMethod("")
	if err != nil {
		t.Errorf("call bytes method failed with: %s", err)
		return
	}
	if data != nil {
		t.Errorf("expected nil data")
		return
	}

	ethereumSym.MockBytesMethod(testClientId+10, testClientId, big.NewInt(0).Bytes())

	_, err = tx.callBytesMethod("")
	if err == nil {
		t.Error("expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId, bigInt.Bytes())
	data, err = tx.callBytesMethod("")
	if err != nil {
		t.Errorf("mock bytes method failed with: %s", err)
		return
	}

	if bytes.Compare(data, bigInt.Bytes()) != 0 {
		t.Error("Sent and received data are not the same")
		return
	}
}

func TestGasPrice(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	testGasPrice := big.NewInt(rand.Int63())

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockBytesMethod(testClientId+10, testClientId, testGasPrice.Bytes())

	gasPrice, err := tx.GasPrice()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId, testGasPrice.Bytes())

	gasPrice, err = tx.GasPrice()
	if err != nil {
		t.Errorf("Getting mock gas price failed with: %s", err)
		return
	}

	if gasPrice.Cmp(testGasPrice) != 0 {
		t.Errorf("Expected gas price `%d` got `%d`", testGasPrice, gasPrice)
		return
	}

	gasPrice, err = tx.GasPrice()
	if err != nil {
		t.Errorf("Getting mock gas price failed with: %s", err)
		return
	}

	if gasPrice.Cmp(testGasPrice) != 0 {
		t.Errorf("Expected gas price `%d` got `%d`", testGasPrice, gasPrice)
		return
	}

}

func TestGasTipCap(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	testGasTipCap := big.NewInt(rand.Int63())

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockBytesMethod(testClientId+10, testClientId, testGasTipCap.Bytes())

	gasTipCap, err := tx.GasTipCap()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId, testGasTipCap.Bytes())

	gasTipCap, err = tx.GasTipCap()
	if err != nil {
		t.Errorf("Getting mock gas tip cap failed with: %s", err)
		return
	}

	if gasTipCap.Cmp(testGasTipCap) != 0 {
		t.Errorf("Expected gas tip cap `%d` got `%d`", testGasTipCap, gasTipCap)
		return
	}

	gasTipCap, err = tx.GasTipCap()
	if err != nil {
		t.Errorf("Getting mock gas tip cap failed with: %s", err)
		return
	}

	if gasTipCap.Cmp(testGasTipCap) != 0 {
		t.Errorf("Expected gas tip cap `%d` got `%d`", testGasTipCap, gasTipCap)
		return
	}

}

func TestGasFeeCap(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	testGasFeeCap := big.NewInt(rand.Int63())

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockBytesMethod(testClientId+10, testClientId, testGasFeeCap.Bytes())

	gasFeeCap, err := tx.GasFeeCap()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId, testGasFeeCap.Bytes())

	gasFeeCap, err = tx.GasFeeCap()
	if err != nil {
		t.Errorf("Getting mock gas fee cap failed with: %s", err)
		return
	}

	if gasFeeCap.Cmp(testGasFeeCap) != 0 {
		t.Errorf("Expected gas tip cap `%d` got `%d`", testGasFeeCap, gasFeeCap)
		return
	}

	gasFeeCap, err = tx.GasFeeCap()
	if err != nil {
		t.Errorf("Getting mock gas fee cap failed with: %s", err)
		return
	}

	if gasFeeCap.Cmp(testGasFeeCap) != 0 {
		t.Errorf("Expected gas tip cap `%d` got `%d`", testGasFeeCap, gasFeeCap)
		return
	}

}

func TestTransactionGas(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockU64method(testClientId+10, 5)

	gas, err := tx.Gas()
	if err == nil {
		t.Errorf("expected error")
		return
	}

	expectedU64 := rand.Uint64()
	ethereumSym.MockU64method(testClientId, expectedU64)
	gas, err = tx.Gas()
	if err != nil {
		t.Errorf("Getting gas failed with: %s", err)
		return
	}
	if gas != expectedU64 {
		t.Errorf("Expected gas `%d` got `%d`", expectedU64, gas)
		return
	}

	gas, err = tx.Gas()
	if err != nil {
		t.Errorf("Getting gas failed with: %s", err)
		return
	}
	if gas != expectedU64 {
		t.Errorf("Expected gas `%d` got `%d`", expectedU64, gas)
		return
	}

}

func TestValue(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	testValue := big.NewInt(rand.Int63())

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockBytesMethod(testClientId+10, testClientId, testValue.Bytes())

	value, err := tx.Value()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId, testValue.Bytes())

	value, err = tx.Value()
	if err != nil {
		t.Errorf("Getting value failed with: %s", err)
		return
	}

	if value.Cmp(testValue) != 0 {
		t.Errorf("Expected value `%d` got `%d`", testValue, value)
		return
	}

	value, err = tx.Value()
	if err != nil {
		t.Errorf("Getting value failed with: %s", err)
		return
	}

	if value.Cmp(testValue) != 0 {
		t.Errorf("Expected value `%d` got `%d`", testValue, value)
		return
	}

}

func TestData(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	testBytes := make([]byte, 1)
	rand.Read(testBytes)

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockBytesMethod(testClientId+10, testClientId, testBytes)

	data, err := tx.Data()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId, testBytes)

	data, err = tx.Data()
	if err != nil {
		t.Errorf("Getting data failed with: %s", err)
		return
	}

	if bytes.Compare(data, testBytes) != 0 {
		t.Error("sent data nad received data are not the same")
		return
	}

	data, err = tx.Data()
	if err != nil {
		t.Errorf("Getting value failed with: %s", err)
		return
	}

	if bytes.Compare(data, testBytes) != 0 {
		t.Error("sent data nad received data are not the same")
		return
	}
}

func TestAddress(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockBytesMethod(testClientId+10, testClientId+10, testAddressBytes)

	address, err := tx.ToAddress()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId, testAddressBytes)

	address, err = tx.ToAddress()
	if err != nil {
		t.Errorf("Getting address failed with: %s", err)
		return
	}

	if bytes.Compare(address, testAddressBytes) != 0 {
		t.Error("sent address and received address are not the same")
		return
	}

	address, err = tx.ToAddress()
	if err != nil {
		t.Errorf("Getting value failed with: %s", err)
		return
	}

	if bytes.Compare(address, testAddressBytes) != 0 {
		t.Error("sent address and received address are not the same")
		return
	}
}

func TestChain(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockBytesMethod(testClientId+10, testClientId+10, testChain.Bytes())

	chain, err := tx.Chain()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId, testChain.Bytes())

	chain, err = tx.Chain()
	if err != nil {
		t.Errorf("Getting chain failed with: %s", err)
		return
	}

	if chain.Cmp(testChain) != 0 {
		t.Errorf("Expected chain `%d` got `%d`", testChain, chain)
		return
	}

	chain, err = tx.Chain()
	if err != nil {
		t.Errorf("Getting chain failed with: %s", err)
		return
	}

	if chain.Cmp(testChain) != 0 {
		t.Errorf("Expected chain `%d` got `%d`", testChain, chain)
		return
	}
}

func TestHash(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	testBytes := make([]byte, 1)
	rand.Read(testBytes)

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	tx.hash = nil

	ethereumSym.MockBytesMethod(testClientId+10, testClientId+10, testBytes)
	hash, err := tx.Hash()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockBytesMethod(testClientId, testClientId, testBytes)

	hash, err = tx.Hash()
	if err != nil {
		t.Errorf("Getting hash failed with: %s", err)
		return
	}

	if bytes.Compare(hash, testBytes) != 0 {
		t.Error("sent hash and received hash are not the same")
		return
	}

	hash, err = tx.Hash()
	if err != nil {
		t.Errorf("Getting hash failed with: %s", err)
		return
	}

	if bytes.Compare(hash, testBytes) != 0 {
		t.Error("sent hash and received data are not the same")
		return
	}
}

func TestSend(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockSendTransaction(testClientId)

	err = tx.Send()
	if err != nil {
		t.Errorf("Sending mock transaction failed with: %s", err)
		return
	}

	ethereumSym.MockSendTransaction(testClientId + 10)

	err = tx.Send()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestRawSignatures(t *testing.T) {
	tx, _, _, err := newMockTransaction()
	if err != nil {
		t.Errorf("Getting mock transaction failed with: %s", err)
		return
	}

	rSig := big.NewInt(rand.Int63())
	vSig := big.NewInt(rand.Int63())
	sSig := big.NewInt(rand.Int63())

	ethereumSym.MockRawSignatures(testClientId, testClientId+10, vSig, rSig, sSig)

	_, err = tx.RawSignatures()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockRawSignatures(testClientId+10, testClientId, vSig, rSig, sSig)

	_, err = tx.RawSignatures()
	if err == nil {
		t.Error("Expected error")
		return
	}

	ethereumSym.MockRawSignatures(testClientId, testClientId, vSig, rSig, sSig)

	rawSigs, err := tx.RawSignatures()
	if err != nil {
		t.Errorf("Getting mock raw signatures failed with: %s", err)
		return
	}

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
	if err != nil {
		t.Errorf("Getting mock raw signatures failed with: %s", err)
		return
	}

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
