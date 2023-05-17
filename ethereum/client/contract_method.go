package ethereum

import (
	"fmt"
	"math/big"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	ec "github.com/taubyte/go-sdk/ethereum/client/codec"
	"github.com/taubyte/go-sdk/utils/codec"
)

// Call invokes the (constant) contract method with params as input values
func (c *ContractMethod) Call(inputParameters ...interface{}) ([]interface{}, error) {
	encoded, err := c.handleInputs(inputParameters...)
	if err != nil {
		return nil, fmt.Errorf("handling input parameters failed with: %s", err)
	}

	var size uint32
	err0 := ethereumSym.EthCallContractSize(uint32(c.client), c.contractID, c.name, &encoded[0], uint32(len(encoded)), &size)
	if err0 != 0 {
		return nil, fmt.Errorf("calling contract failed with: %s", err0)
	}
	if size == 0 {
		return nil, nil
	}

	buf := make([]byte, size)
	err0 = ethereumSym.EthCallContract(uint32(c.client), c.contractID, c.name, &buf[0])
	if err0 != 0 {
		return nil, fmt.Errorf("called contract but unable to write data with: %s", err0)
	}

	var outputsBytes [][]byte
	err = codec.Convert(buf).To(&outputsBytes)
	if err != nil {
		return nil, fmt.Errorf("encoded output bytes to bytes slice failed with: %s", err)
	}

	if len(outputsBytes) != len(c.outputs) {
		return nil, fmt.Errorf("unexpected number of outputs: got `%d` expected `%d`", len(outputsBytes), len(c.outputs))
	}

	var outputs []interface{}
	for idx, value := range outputsBytes {
		outputType := c.outputs[idx]
		decoder, err := ec.Converter(outputType).Decoder()
		if err != nil {
			return nil, fmt.Errorf("getting output decoder for value %v failed with: %s", value, err)
		}

		output, err := decoder(value)
		if err != nil {
			return nil, fmt.Errorf("decoding output buf to type `%s` failed with: %s", outputType, err)
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}

// Transact invokes the (paid) contract method with params as input values. If chain id is nil, then current chain Id is used.
func (c *ContractMethod) Transact(chainID *big.Int, privateKey []byte, inputParameters ...interface{}) (*Transaction, error) {
	var err error
	if len(privateKey) == 0 {
		return nil, fmt.Errorf("private key cannot be empty")
	}

	if chainID == nil {
		chainID, err = c.client.CurrentChainID()
		if err != nil {
			return nil, fmt.Errorf("getting current chain Id failed with: %s", err)
		}
	}

	chainBytes := chainID.Bytes()
	var transactionID uint32

	encoded, err := c.handleInputs(inputParameters...)
	if err != nil {
		return nil, fmt.Errorf("handling inputs failed with: %s", err)
	}

	err0 := ethereumSym.EthTransactContract(uint32(c.client), c.contractID, &chainBytes[0], uint32(len(chainBytes)), c.name, &privateKey[0], uint32(len(privateKey)), &encoded[0], uint32(len(encoded)), 0, &transactionID)
	if err0 != 0 {
		return nil, fmt.Errorf("transacting contract method `%s` failed with: %s", c.name, err0)
	}

	return &Transaction{id: transactionID, contractID: c.contractID}, nil
}

func (c *ContractMethod) handleInputs(inputParameters ...interface{}) ([]byte, error) {
	if len(inputParameters) == 0 {
		return []byte{0}, nil
	}

	inputs := make([][]byte, 0)
	for idx, param := range inputParameters {
		encoder, err := ec.Converter(c.inputs[idx]).Encoder()
		if err != nil {
			return nil, fmt.Errorf("getting encoder for param `%v` failed with: %s", param, err)
		}

		encodedInput, err := encoder(param)
		if err != nil {
			return nil, fmt.Errorf("encoding input for `%v` failed with: %s", param, err)
		}

		inputs = append(inputs, encodedInput)
	}

	var encoded []byte
	err := codec.Convert(inputs).To(&encoded)
	if err != nil {
		return nil, fmt.Errorf("encoding input list failed with: %s", err)
	}

	return encoded, nil
}
