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
		return nil, fmt.Errorf("Handling input parameters failed with: %s", err)
	}

	var size uint32
	err0 := ethereumSym.EthCallContractSize(uint32(c.client), c.contractId, c.name, &encoded[0], uint32(len(encoded)), &size)
	if err0 != 0 {
		return nil, fmt.Errorf("Calling contract failed with: %s", err0)
	}
	if size == 0 {
		return nil, nil
	}

	buf := make([]byte, size)
	err0 = ethereumSym.EthCallContract(uint32(c.client), c.contractId, c.name, &buf[0])
	if err0 != 0 {
		return nil, fmt.Errorf("Called contract but unable to write data with: %s", err0)
	}

	var outputsBytes [][]byte
	err = codec.Convert(buf).To(&outputsBytes)
	if err != nil {
		return nil, fmt.Errorf("Encoded output bytes to bytes slice failed with: %s", err)
	}

	if len(outputsBytes) != len(c.outputs) {
		return nil, fmt.Errorf("Unexpected number of outputs: got `%d` expected `%d`", len(outputsBytes), len(c.outputs))
	}

	var outputs []interface{}
	for idx, value := range outputsBytes {
		outputType := c.outputs[idx]
		decoder, err := ec.Converter(outputType).Decoder()
		if err != nil {
			return nil, fmt.Errorf("Getting output decoder for value %v failed with: %s", value, err)
		}

		output, err := decoder(value)
		if err != nil {
			return nil, fmt.Errorf("Decoding output buf to type `%s` failed with: %s", outputType, err)
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}

// Transact invokes the (paid) contract method with params as input values. If chain id is nil, then current chain Id is used.
func (c *ContractMethod) Transact(chainId *big.Int, privateKey string, inputParameters ...interface{}) (*Transaction, error) {
	var err error
	if privateKey == "" {
		return nil, fmt.Errorf("Private key cannot be empty")
	}

	if chainId == nil {
		chainId, err = c.client.CurrentChainId()
		if err != nil {
			return nil, fmt.Errorf("Getting current chain Id failed with: %s", err)
		}
	}

	chainBytes := chainId.Bytes()
	var transactionId uint32

	encoded, err := c.handleInputs(inputParameters...)
	if err != nil {
		return nil, fmt.Errorf("Handling inputs failed with: %s", err)
	}

	err0 := ethereumSym.EthTransactContract(uint32(c.client), c.contractId, &chainBytes[0], uint32(len(chainBytes)), c.name, privateKey, &encoded[0], uint32(len(encoded)), &transactionId)
	if err0 != 0 {
		return nil, fmt.Errorf("Transacting contract method `%s` failed with: %s", c.name, err0)
	}

	return &Transaction{id: transactionId, contractId: c.contractId}, nil
}

func (c *ContractMethod) handleInputs(inputParameters ...interface{}) ([]byte, error) {
	if len(inputParameters) == 0 {
		return []byte{0}, nil
	}

	inputs := make([][]byte, 0)
	for idx, param := range inputParameters {
		encoder, err := ec.Converter(c.inputs[idx]).Encoder()
		if err != nil {
			return nil, fmt.Errorf("Getting encoder for param `%v` failed with: %s", param, err)
		}

		encodedInput, err := encoder(param)
		if err != nil {
			return nil, fmt.Errorf("Encoding input for `%v` failed with: %s", param, err)
		}

		inputs = append(inputs, encodedInput)
	}

	var encoded []byte
	err := codec.Convert(inputs).To(&encoded)
	if err != nil {
		return nil, fmt.Errorf("Encoding input list failed with: %s", err)
	}

	return encoded, nil
}
