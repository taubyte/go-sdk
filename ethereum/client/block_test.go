package ethereum

import (
	"math/big"
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

func TestTransaction(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	_, block, err := newMockBlock()
	if err != nil {
		t.Errorf("Getting mocked block failed with: %s", err)
		return
	}

	ethereumSym.MockBlockTransaction(testClientId, testTransactionId)

	_, err = block.Transaction(nil)
	if err == nil {
		t.Error("Expected error")
		return
	}

	tx, err := block.Transaction(testTransactionHash)
	if err != nil {
		t.Errorf("Getting transaction failed with: %s", err)
		return
	}
	if tx.id != testTransactionId {
		t.Errorf("Expected transaction id `%d` got `%d`", testTransactionId, tx.id)
		return
	}

	ethereumSym.MockBlockTransaction(testClientId+10, testTransactionId)

	_, err = block.Transaction(testTransactionHash)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestTransactions(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	_, block, err := newMockBlock()
	if err != nil {
		t.Errorf("Getting mocked block failed with: %s", err)
		return
	}

	tx, err := block.Transactions()
	if err != nil && tx != nil {
		t.Error("nil tx and error")
		return
	}

	err = ethereumSym.MockBlockTransactions(testClientId, testClientId, testTransactions, false)
	if err != nil {
		t.Error(err)
		return
	}

	txs, err := block.Transactions()
	if err != nil {
		t.Errorf("Getting mocked transactions failed with: %s", err)
		return
	}
	if len(txs) != len(testTransactions) {
		t.Errorf("Expected `%d` transactions got `%d`", len(testTransactions), len(txs))
		return
	}

	err = ethereumSym.MockBlockTransactions(testClientId, testClientId, []uint32{1}, true)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = block.Transactions()
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockBlockTransactions(testClientId, testClientId+10, testTransactions, false)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = block.Transactions()
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockBlockTransactions(testClientId+10, testClientId, testTransactions, false)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = block.Transactions()
	if err == nil {
		t.Error("Expected error")
		return
	}

}

func TestBlockNumber(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	err = ethereumSym.MockBlockNumber(testClientId, testClientId, testBlockNumber)
	if err != nil {
		t.Error(err)
		return
	}

	_, block, err := newMockBlock()
	if err != nil {
		t.Errorf("Getting mocked block failed with: %s", err)
		return
	}

	blockNum, err := block.Number()
	if err != nil {
		t.Errorf("Getting mocked block number failed with: %s", err)
		return
	}

	if blockNum.Cmp(testBlockNumber) != 0 {
		t.Errorf("Expected block number `%d` got `%d`", testBlockNumber, blockNum)
		return
	}

	err = ethereumSym.MockBlockNumber(testClientId, testClientId+10, testBlockNumber)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = block.Number()
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockBlockNumber(testClientId, testClientId, big.NewInt(0))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = block.Number()
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockBlockNumber(testClientId+10, testClientId, testBlockNumber)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = block.Number()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestNonce(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	ethereumSym.MockBlockNonce(testClientId, testNonce)
	err = ethereumSym.MockBlockNumber(testClientId, testClientId, testBlockNumber)
	if err != nil {
		t.Error(err)
		return
	}

	_, block, err := newMockBlock()
	if err != nil {
		t.Errorf("Getting mocked block failed with: %s", err)
		return
	}

	nonceVal, err := block.NonceFromPrivateKey(testString)
	if err != nil {
		t.Errorf("Getting mocked nonce failed with: %s", err)
		return
	}
	if nonceVal != testNonce {
		t.Errorf("Expected nonce `%d` got `%d`", testNonce, nonceVal)
		return
	}

	ethereumSym.MockBlockNonce(testClientId+10, testNonce)

	_, err = block.NonceFromPrivateKey(testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	err = ethereumSym.MockBlockNumber(testClientId+10, testClientId, testBlockNumber)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = block.NonceFromPrivateKey(testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	_, err = block.NonceFromPrivateKey("")
	if err == nil {
		t.Error("Expected error")
		return
	}
}
