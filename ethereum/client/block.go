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
		return nil, errors.New("Invalid hash")
	}

	var transactionId uint32
	if err0 := ethereumSym.EthGetTransactionFromBlockByHash(uint32(b.client), &b.id, &transactionId, &hash[0]); err0 != 0 {
		return nil, fmt.Errorf("Getting transaction by hash from block failed with %s", err0)
	}

	return &Transaction{
		id:            transactionId,
		blockId:       b.id,
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
		return nil, fmt.Errorf("Getting transactions list size from block failed with: %s", err0)
	}
	if size == 0 {
		return nil, nil
	}

	transactionHashesBytes := make([]byte, size)
	if err0 := ethereumSym.EthGetTransactionsFromBlock(uint32(b.client), &b.id, &transactionHashesBytes[0]); err0 != 0 {
		return nil, fmt.Errorf("Getting transactions from block failed with: %s", err0)
	}

	var transactionHashes []uint32
	err := codec.Convert(transactionHashesBytes).To(&transactionHashes)
	if err != nil || uint32(len(transactionHashes)) != arrSize {
		return nil, fmt.Errorf("Converting transaction hashes failed with: %s", err)
	}

	var transactionList []*Transaction
	for _, transaction := range transactionHashes {
		t := &Transaction{
			id:            transaction,
			client:        b.client,
			blockId:       b.id,
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
		return nil, fmt.Errorf("Getting block number size failed with %s", err)
	}

	if size != 0 {
		buf := make([]byte, size)
		if err := ethereumSym.EthBlockNumberFromId(uint32(b.client), &b.id, &buf[0]); err != 0 {
			return nil, fmt.Errorf("Getting block number failed with: %s", err)
		}

		return new(big.Int).SetBytes(buf), nil
	}

	return nil, errors.New("Unable to get block number from block")
}

// NonceFromPrivateKey returns the account nonce of the account given from the secp256k1 hexkey.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (b *Block) NonceFromPrivateKey(hexKey string) (uint64, error) {
	if hexKey == "" {
		return 0, fmt.Errorf("Hex key cannot be empty")
	}

	blockNumber, err := b.Number()
	if err != nil {
		return 0, fmt.Errorf("Getting block number failed with: %s", err)
	}

	blockBuf := blockNumber.Bytes()
	var nonce uint64
	if err := ethereumSym.EthNonceFromPrivateKey(uint32(b.client), hexKey, uint32(len(blockBuf)), &blockBuf[0], &nonce); err != 0 {
		return 0, fmt.Errorf("Getting Nonce from private key failed with: %s", err)
	}

	return nonce, nil
}
