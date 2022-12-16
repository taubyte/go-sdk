package ethereum

import (
	"math/big"
	"testing"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

func TestClient(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	client, err := newMockClient()
	if err != nil {
		t.Errorf("Creating new rpc client failed with: %s", err)
		return
	}

	_, err = New(testString)
	if err == nil {
		t.Error("Expected error")
		return
	}

	client.Close()

	ethereumSym.MockClientNew(int32(testClientId) * -1)

	_, err = New(testRpcUrl)
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestBlock(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	client, _, err := newMockBlock()
	if err != nil {
		t.Errorf("Getting mocked block failed with: %s", err)
		return
	}

	ethereumSym.MockCurrentBlockNumber(testClientId, testCurrentBlockNumber)
	ethereumSym.MockBlockByNumber(testClientId, testBlockId)

	block, err := client.BlockByNumber(nil)
	if err != nil {
		t.Errorf("Getting block by number with nil block failed with: %s", err)
		return
	}

	if block.id != testBlockId {
		t.Errorf("Expected block id `%d` got `%d`", testBlockId, block.id)
		return
	}

	ethereumSym.MockBlockByNumber(testClientId+10, testBlockId)

	if _, err = client.BlockByNumber(nil); err == nil {
		t.Errorf("Expected error")
		return
	}

	ethereumSym.MockCurrentBlockNumber(testClientId+10, testCurrentBlockNumber)

	if _, err = client.BlockByNumber(nil); err == nil {
		t.Errorf("Expected error")
		return
	}
}

func TestCurrentChain(t *testing.T) {
	err := setTestVars()
	if err != nil {
		t.Error(err)
		return
	}

	if err := ethereumSym.MockCurrentChainId(testClientId, testClientId, testChain); err != nil {
		t.Error(err)
		return
	}

	client, err := newMockClient()
	if err != nil {
		t.Errorf("New mock client failed with: %s", err)
		return
	}

	chain, err := client.CurrentChainId()
	if err != nil {
		t.Errorf("Current chain id failed with %s", err)
		return
	}
	if chain.Cmp(testChain) != 0 {
		t.Errorf("Expected chain id `%d` got `%d`", testChain, chain)
		return
	}

	if err = ethereumSym.MockCurrentChainId(testClientId, testClientId+10, testChain); err != nil {
		t.Error(err)
		return
	}

	_, err = client.CurrentChainId()
	if err == nil {
		t.Errorf("Expected error")
		return
	}

	if err = ethereumSym.MockCurrentChainId(testClientId+10, testClientId+10, testChain); err != nil {
		t.Error(err)
		return
	}

	_, err = client.CurrentChainId()
	if err == nil {
		t.Errorf("Expected error")
		return
	}

	if err = ethereumSym.MockCurrentChainId(testClientId, testClientId, big.NewInt(0)); err != nil {
		t.Error(err)
		return
	}

	_, err = client.CurrentChainId()
	if err == nil {
		t.Errorf("Expected error")
		return
	}
}
