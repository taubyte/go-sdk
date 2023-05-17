package ethereum

import (
	"fmt"
	"math/big"
	"math/rand"
	"sync"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	"github.com/taubyte/go-sdk/ethereum/client/bytes"
	"github.com/taubyte/go-sdk/utils/codec"
)

var (
	testString = "test string"
	testBytes  = []byte("test bytes")

	testTransactionID   uint32
	testTransactionHash = make([]byte, 32)
	testTransactions    []uint32

	testNonce uint64

	testClientID uint32

	testCurrentBlockNumber uint64
	testBlockNumber        *big.Int
	testBlockID            uint64

	testChain = big.NewInt(rand.Int63())

	testRPCURL        = "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	testAddressRaw, _ = bytes.AddressFromHex(testAddress)
	testAddressBytes  = testAddressRaw.Bytes()

	testContractID         uint32
	testContractMethodSize uint32
	testContractEventsSize uint32

	testInputFailureMethod  = "inputFailureMethod"
	testOutputFailureMethod = "outputFailureMethod"
	testIncompatibleVar     = new(sync.WaitGroup)

	testPassingMethod = "passingMethod"
	testPassingInput  uint32
	testPassingOutput uint32

	testContract ethereumSym.MockContract

	testPrivateKeyHex = "91e4a13e5a30ad353cdf5ea7bb909dfdf967122e3b43e331ad947b68a3899b2c"
	testAddress       = "0x3b8C3C867B2bc1AEb6ea2f565C82F27116c3c54E"
)

func setTestVars() error {
	testTransactionID = uint32(rand.Int31())
	testTransactions = []uint32{uint32(rand.Int31()), uint32(rand.Int31()), uint32(rand.Int31())}

	testNonce = uint64(rand.Int63())

	testClientID = uint32(rand.Int31())

	testCurrentBlockNumber = uint64(rand.Int63())
	testBlockNumber = big.NewInt(int64(testCurrentBlockNumber))
	testBlockID = uint64(rand.Int63())

	testChain = big.NewInt(rand.Int63())

	testContractID = uint32(rand.Int31())

	testPassingInput = uint32(rand.Int31())
	testPassingOutput = uint32(rand.Int31())

	rand.Read(testTransactionHash)

	testContract = ethereumSym.MockContract{
		Methods: map[string]ethereumSym.MockContractMethod{
			testPassingMethod: {
				Inputs:  []interface{}{testPassingInput},
				Outputs: []interface{}{testPassingOutput},
			},
			testInputFailureMethod: {
				Inputs:  []interface{}{testIncompatibleVar},
				Outputs: []interface{}{testIncompatibleVar},
			},
			testOutputFailureMethod: {
				Inputs:  []interface{}{testPassingInput},
				Outputs: []interface{}{testIncompatibleVar},
			},
		},

		ContractSizeClientId: testClientID,
		ContractDataClientId: testClientID,
		MethodSizeClientId:   testClientID,
		MethodDataClientId:   testClientID,
		CallSizeClientId:     testClientID,
		CallDataClientId:     testClientID,
	}

	methods := []string{}
	for method := range testContract.Methods {
		methods = append(methods, method)
	}

	var encoded []byte

	err := codec.Convert(methods).To(&encoded)
	if err != nil {
		return fmt.Errorf("Converting methods list to bytes failed with: " + err.Error())
	}

	testContractMethodSize = uint32(len(encoded))

	return nil
}

func newMockTransaction() (*Transaction, *Block, Client, error) {
	client, block, err := newMockBlock()
	if err != nil {
		return nil, nil, 0, err
	}

	ethereumSym.MockBlockTransaction(testClientID, testTransactionID)

	tx, err := block.Transaction(testTransactionHash)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Getting mocked transaction from block failed with: %s", err)
	}

	return tx, block, client, nil
}

func newMockClient() (Client, error) {
	ethereumSym.MockClientNew(int32(testClientID))

	client, err := New(testRPCURL)
	if err != nil {
		return 0, err
	}

	if client != Client(testClientID) {
		return 0, fmt.Errorf("Expected client value `%d` got `%d`", testClientID, client)
	}

	return client, nil
}

func newMockBlock() (Client, *Block, error) {
	client, err := newMockClient()
	if err != nil {
		return 0, nil, err
	}

	ethereumSym.MockBlockNumber(testClientID, testClientID, testBlockNumber)

	ethereumSym.MockBlockByNumber(testClientID, testBlockID)

	block, err := client.BlockByNumber(testBlockNumber)
	if err != nil {
		return 0, nil, err
	}

	if block.id != testBlockID {
		return 0, nil, fmt.Errorf("Expected block id to be `%d` got `%d`", testBlockID, block.id)
	}

	return client, block, nil
}
