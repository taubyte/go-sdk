package ethereum

import (
	"math/big"
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	"gotest.tools/assert"
)

func TestClient(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	client, err := newMockClient()
	assert.NilError(t, err)

	_, err = New(testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	client.Close()

	ethereumSym.MockClientNew(int32(testClientID) * -1)

	_, err = New(testRPCURL)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestBlock(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	client, _, err := newMockBlock()
	assert.NilError(t, err)

	ethereumSym.MockCurrentBlockNumber(testClientID, testCurrentBlockNumber)
	ethereumSym.MockBlockByNumber(testClientID, testBlockID)

	block, err := client.BlockByNumber(nil)
	assert.NilError(t, err)

	if block.id != testBlockID {
		t.Errorf("Expected block id `%d` got `%d`", testBlockID, block.id)
		return
	}

	ethereumSym.MockBlockByNumber(testClientID+10, testBlockID)

	if _, err = client.BlockByNumber(nil); err == nil {
		t.Errorf("Expected error")
		return
	}

	ethereumSym.MockCurrentBlockNumber(testClientID+10, testCurrentBlockNumber)

	if _, err = client.BlockByNumber(nil); err == nil {
		t.Errorf("Expected error")
		return
	}
}

func TestCurrentChain(t *testing.T) {
	err := setTestVars()
	assert.NilError(t, err)

	err = ethereumSym.MockCurrentChainId(testClientID, testClientID, testChain)
	assert.NilError(t, err)

	client, err := newMockClient()
	assert.NilError(t, err)

	chain, err := client.CurrentChainID()
	assert.NilError(t, err)

	if chain.Cmp(testChain) != 0 {
		t.Errorf("Expected chain id `%d` got `%d`", testChain, chain)
		return
	}

	err = ethereumSym.MockCurrentChainId(testClientID, testClientID+10, testChain)
	assert.NilError(t, err)

	_, err = client.CurrentChainID()
	if err == nil {
		t.Errorf("Expected error")
		return
	}

	err = ethereumSym.MockCurrentChainId(testClientID+10, testClientID+10, testChain)
	assert.NilError(t, err)

	_, err = client.CurrentChainID()
	if err == nil {
		t.Errorf("Expected error")
		return
	}

	err = ethereumSym.MockCurrentChainId(testClientID, testClientID, big.NewInt(0))
	assert.NilError(t, err)

	_, err = client.CurrentChainID()
	if err == nil {
		t.Errorf("Expected error")
		return
	}
}
