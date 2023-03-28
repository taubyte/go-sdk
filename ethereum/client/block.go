package ethereum

import (
	"errors"
	"fmt"
	"math/big"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	"github.com/taubyte/go-sdk/utils/codec"
)

// Transaction returns transaction from block with the given transaction hash.
//
// Transaction hash is 32 bytes, if inputted hash is longer than 32 bytes hash will be trimmed.
func (b *Block) Transaction(hash []byte) (*Transaction, error) {
	if len(hash) < 32 {
		return nil, errors.New("invalid hash")
	}

	var transactionID uint32
	if err0 := ethereumSym.EthGetTransactionFromBlockByHash(uint32(b.client), &b.id, &transactionID, &hash[0]); err0 != 0 {
		return nil, fmt.Errorf("getting transaction by hash from block failed with %s", err0)
	}

	return &Transaction{
		id:            transactionID,
		blockID:       b.id,
		client:        b.client,
		hash:          hash,
		rawSignatures: rawSignatures{},
	}, nil
}

// Transactions returns all transactions from the given block.
func (b *Block) Transactions() ([]*Transaction, error) {
	var size uint32
	var arrSize uint32
	if err0 := ethereumSym.EthGetTransactionsFromBlockSize(uint32(b.client), &b.id, &size, &arrSize); err0 != 0 {
		return nil, fmt.Errorf("getting transactions list size from block failed with: %s", err0)
	}
	if size == 0 {
		return nil, nil
	}

	transactionHashesBytes := make([]byte, size)
	if err0 := ethereumSym.EthGetTransactionsFromBlock(uint32(b.client), &b.id, &transactionHashesBytes[0]); err0 != 0 {
		return nil, fmt.Errorf("getting transactions from block failed with: %s", err0)
	}

	var transactionHashes []uint32
	err := codec.Convert(transactionHashesBytes).To(&transactionHashes)
	if err != nil || uint32(len(transactionHashes)) != arrSize {
		return nil, fmt.Errorf("converting transaction hashes failed with: %s", err)
	}

	var transactionList []*Transaction
	for _, transaction := range transactionHashes {
		t := &Transaction{
			id:            transaction,
			client:        b.client,
			blockID:       b.id,
			rawSignatures: rawSignatures{},
		}

		transactionList = append(transactionList, t)
	}

	return transactionList, nil
}

// Number returns the *big.Int value of the block number
func (b *Block) Number() (*big.Int, error) {
	var size uint32
	if err := ethereumSym.EthBlockNumberFromIdSize(uint32(b.client), &b.id, &size); err != 0 {
		return nil, fmt.Errorf("getting block number size failed with %s", err)
	}

	if size != 0 {
		buf := make([]byte, size)
		if err := ethereumSym.EthBlockNumberFromId(uint32(b.client), &b.id, &buf[0]); err != 0 {
			return nil, fmt.Errorf("getting block number failed with: %s", err)
		}

		return new(big.Int).SetBytes(buf), nil
	}

	return nil, errors.New("unable to get block number from block")
}
