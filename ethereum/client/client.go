package ethereum

import (
	"errors"
	"fmt"
	"math/big"

	goUrl "net/url"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

// New connects a client to the given rpc URL.
func New(url string) (Client, error) {
	_, err := goUrl.ParseRequestURI(url)
	if err != nil {
		return 0, fmt.Errorf("Parsing url failed with: %s", err)
	}

	var clientID uint32
	err0 := ethereumSym.EthNew(&clientID, url)
	if err0 != 0 {
		return 0, fmt.Errorf("Creating new ethereum client failed with %s", err0)
	}

	return Client(clientID), nil
}

// CurrentBlockNumber returns the most recent block number.
func (c Client) CurrentBlockNumber() (number uint64, err error) {
	if err := ethereumSym.EthCurrentBlockNumber(uint32(c), &number); err != 0 {
		return 0, fmt.Errorf("Getting current block number failed with: %s", err)
	}

	return number, nil
}

// BlockByNumber returns a block from the current canonical chain. If number is nil, the latest known block is returned.
func (c Client) BlockByNumber(blockNumber *big.Int) (*Block, error) {
	if blockNumber == nil {
		current, err := c.CurrentBlockNumber()
		if err != nil {
			return nil, fmt.Errorf("Getting current block number failed with: %s", err)
		}

		blockNumber = big.NewInt(int64(current))
	}

	blockNumberBytes := blockNumber.Bytes()

	block := Block{
		client: c,
	}

	if err := ethereumSym.EthBlockByNumber(uint32(c), uint32(len(blockNumberBytes)), &blockNumberBytes[0], &block.id); err != 0 {
		return nil, fmt.Errorf("Getting block by block number failed with: %s", err)
	}

	return &block, nil
}

// CurrentChainID retrieves the current chain ID for transaction replay protection.
func (c Client) CurrentChainID() (*big.Int, error) {
	var size uint32
	if err := ethereumSym.EthCurrentChainIdSize(uint32(c), &size); err != 0 {
		return nil, fmt.Errorf("Getting current chain id failed with: %s", err)
	}

	if size != 0 {
		bytes := make([]byte, size)
		if err := ethereumSym.EthCurrentChainId(uint32(c), &bytes[0]); err != 0 {
			return nil, fmt.Errorf("Getting current chain id failed with: %s", err)
		}

		return new(big.Int).SetBytes(bytes), nil
	}

	return nil, errors.New("Chain not found")
}

// Close will close the Ethereum Client
func (c Client) Close() {
	ethereumSym.EthCloseClient(uint32(c))
}
