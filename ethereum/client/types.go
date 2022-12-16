package ethereum

import (
	"math/big"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client uint32

// Block defines wrappers for a block retrieved by the client.
type Block struct {
	id     uint64
	client Client
}

// Contract defines typed wrappers for a contract with given abi.
type Contract struct {
	id      uint32
	client  Client
	methods map[string]*ContractMethod
	address string
}

type ContractMethod struct {
	contractId uint32
	client     Client
	name       string
	inputs     []string
	outputs    []string
}

// Transaction defines wrappers for a transaction retrieved by the client.
type Transaction struct {
	id            uint32
	client        Client
	blockId       uint64
	contractId    uint32
	json          []byte
	nonce         uint64
	gasPrice      *big.Int
	gasTipCap     *big.Int
	gasFeeCap     *big.Int
	gas           uint64
	value         *big.Int
	data          []byte
	rawSignatures rawSignatures
	toAddress     []byte
	chain         *big.Int
	hash          []byte
}

type rawSignatures struct {
	VSig *big.Int
	RSig *big.Int
	SSig *big.Int
}

type abi struct {
	json    string
	inputs  []string
	outputs []string
}
