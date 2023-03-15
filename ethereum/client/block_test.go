package ethereum

import (
	"math/big"
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	"gotest.tools/assert"
)

func TestTransaction(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	_, block, err := newMockBlock()
	assert.NilError(t, err)

	ethereumSym.MockBlockTransaction(testClientID, testTransactionID)

	_, err = block.Transaction(nil)
	if err == nil {
		t.Error("Expected error")
		return
	}

	tx, err := block.Transaction(testTransactionHash)
	assert.NilError(t, err)
	if tx.id != testTransactionID {
		t.Errorf("Expected transaction id `%d` got `%d`", testTransactionID, tx.id)
		return
	}

	ethereumSym.MockBlockTransaction(testClientID+10, testTransactionID)

	_, err = block.Transaction(testTransactionHash)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestTransactions(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	_, block, err := newMockBlock()
	assert.NilError(t, err)

	tx, err := block.Transactions()
	if err != nil && tx != nil {
		t.Error("nil tx and error")
		return
	}

	err = ethereumSym.MockBlockTransactions(testClientID, testClientID, testTransactions, false)
	assert.NilError(t, err)

	txs, err := block.Transactions()
	assert.NilError(t, err)
	if len(txs) != len(testTransactions) {
		t.Errorf("Expected `%d` transactions got `%d`", len(testTransactions), len(txs))
		return
	}

	err = ethereumSym.MockBlockTransactions(testClientID, testClientID, []uint32{1}, true)
	assert.NilError(t, err)

	_, err = block.Transactions()
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockBlockTransactions(testClientID, testClientID+10, testTransactions, false)
	assert.NilError(t, err)

	_, err = block.Transactions()
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockBlockTransactions(testClientID+10, testClientID, testTransactions, false)
	assert.NilError(t, err)

	_, err = block.Transactions()
	if err == nil {
		t.Error("Expected error")
		return
	}

}

func TestBlockNumber(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	err = ethereumSym.MockBlockNumber(testClientID, testClientID, testBlockNumber)
	assert.NilError(t, err)

	_, block, err := newMockBlock()
	assert.NilError(t, err)

	blockNum, err := block.Number()
	assert.NilError(t, err)

	if blockNum.Cmp(testBlockNumber) != 0 {
		t.Errorf("Expected block number `%d` got `%d`", testBlockNumber, blockNum)
		return
	}

	err = ethereumSym.MockBlockNumber(testClientID, testClientID+10, testBlockNumber)
	assert.NilError(t, err)

	_, err = block.Number()
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockBlockNumber(testClientID, testClientID, big.NewInt(0))
	assert.NilError(t, err)

	_, err = block.Number()
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockBlockNumber(testClientID+10, testClientID, testBlockNumber)
	assert.NilError(t, err)

	_, err = block.Number()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
