package ethereum

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/big"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	"github.com/taubyte/go-sdk/ethereum/client/bytes"
	"github.com/taubyte/go-sdk/utils/codec"
)

// DeployContract method deploys the contract on given chain,returns low level contract interface through which calls
// and transactions may be made through.
func (c Client) DeployContract(abi, byteCode io.Reader, chainID *big.Int, privateKey []byte) (*Contract, *Transaction, error) {
	var err error
	if abi == nil {
		return nil, nil, fmt.Errorf("ABI is nil")
	}
	if byteCode == nil {
		return nil, nil, fmt.Errorf("Byte code is nil")
	}
	if len(privateKey) == 0 {
		return nil, nil, fmt.Errorf("Private key is empty")
	}
	if chainID == nil {
		chainID, err = c.CurrentChainID()
		if err != nil {
			return nil, nil, fmt.Errorf("Getting current chain failed with: %s", err)
		}
	}

	chainBytes := chainID.Bytes()
	_byteCode, err := ioutil.ReadAll(byteCode)
	if err != nil {
		return nil, nil, fmt.Errorf("Reading byte code failed with: %s", err)
	}

	abiBytes, err := ioutil.ReadAll(abi)
	if err != nil {
		return nil, nil, fmt.Errorf("Reading abi failed with: %s", err)
	}

	var methodSize uint32
	address := make([]byte, 20)
	var contractID uint32
	var transactionID uint32

	err0 := ethereumSym.EthDeployContract(uint32(c), &chainBytes[0], uint32(len(chainBytes)), string(_byteCode), &abiBytes[0], uint32(len(abiBytes)), &privateKey[0], uint32(len(privateKey)), &address[0], &methodSize, &contractID, &transactionID)
	if err0 != 0 {
		return nil, nil, fmt.Errorf("Deploying contract failed with: %s", err0)
	}

	addressString, _ := bytes.AddressToHex(address)

	contract, err := c.getContract(addressString, contractID, methodSize)
	if err != nil {
		return nil, nil, fmt.Errorf("Getting contract failed with: %s", err)
	}

	return contract, &Transaction{id: transactionID, contractID: contractID}, err
}

// NewBoundContract method creates a low level contract interface through which calls
// and transactions may be made through.
func (c Client) NewBoundContract(abi io.Reader, address string) (*Contract, error) {
	if address == "" {
		return nil, fmt.Errorf("Address cannot be empty")
	}
	if abi == nil {
		return nil, fmt.Errorf("ABI cannot be nil")
	}

	abiBytes, err := ioutil.ReadAll(abi)
	if err != nil {
		return nil, fmt.Errorf("Reading abi failed with: %s", err)
	}

	var (
		contractID uint32
		methodSize uint32
	)

	err0 := ethereumSym.EthNewContractSize(uint32(c), &abiBytes[0], uint32(len(abiBytes)), address, &methodSize, &contractID)
	if err0 != 0 {
		return nil, fmt.Errorf("Creating new bound contract failed with: %s", err0)
	}

	return c.getContract(address, contractID, methodSize)
}

// Methods lists the available methods for within the given contract
func (c *Contract) Methods() []*ContractMethod {
	var methods []*ContractMethod
	for _, method := range c.methods {
		methods = append(methods, method)
	}

	return methods
}

// Name returns the name of the method.
func (c *ContractMethod) Name() string {
	return c.name
}

// Method returns the contract method with the corresponding inputted name.
func (c *Contract) Method(name string) (*ContractMethod, error) {
	contract, ok := c.methods[name]
	if ok == false {
		return nil, fmt.Errorf("Contract method `%s` not found", name)
	}

	return contract, nil
}

// Address returns the wallet address of the Contract
func (c *Contract) Address() string {
	return c.address
}

// getContract method is a helper method for NewBoundContract, and DeployContract methods.
// Gets the contract methods, their input, and output types for the given contract.
func (c Client) getContract(address string, contractID, methodSize uint32) (*Contract, error) {
	methods := make(map[string]*ContractMethod, 0)
	if methodSize != 0 {
		encodedMethods := make([]byte, methodSize)
		err0 := ethereumSym.EthNewContract(uint32(c), contractID, &encodedMethods[0])
		if err0 != 0 {
			return nil, fmt.Errorf("Getting contract methods data failed with: %s", err0)
		}

		var methodList []string

		err := codec.Convert(encodedMethods).To(&methodList)
		if err != nil || len(methodList) == 0 {
			return nil, fmt.Errorf("Converting encoded methods failed with: %s", err)
		}

		for _, method := range methodList {
			inputs, outputs, err := c.methodInputOutput(method, contractID)
			if err != nil {
				return nil, err
			}

			methods[method] = &ContractMethod{
				client:     c,
				contractID: contractID,
				name:       method,
				inputs:     inputs,
				outputs:    outputs,
			}
		}
	}

	return &Contract{id: contractID, client: c, methods: methods, address: address}, nil
}

func (c Client) methodInputOutput(method string, contractID uint32) (inputs []string, outputs []string, err error) {
	var inputSize uint32
	var outputSize uint32
	err0 := ethereumSym.EthGetContractMethodSize(uint32(c), contractID, method, &inputSize, &outputSize)
	if err0 != 0 {
		return nil, nil, fmt.Errorf("Getting contract method `%s` inputs and outputs size failed with: %s", method, err0)
	}

	var inputBuf []byte
	var outputBuf []byte
	if inputSize == 0 {
		inputBuf = make([]byte, 1)
	} else {
		inputBuf = make([]byte, inputSize)
	}
	if outputSize == 0 {
		outputBuf = make([]byte, 1)
	} else {
		outputBuf = make([]byte, outputSize)
	}

	err0 = ethereumSym.EthGetContractMethod(uint32(c), contractID, method, &inputBuf[0], &outputBuf[0])
	if err0 != 0 {
		return nil, nil, fmt.Errorf("Getting contract method `%s` inputs and outputs failed with: %s", method, err0)
	}

	if inputSize != 0 {
		err = codec.Convert(inputBuf).To(&inputs)
		if err != nil || len(inputs) == 0 {
			return nil, nil, fmt.Errorf("Converting contract method `%s` input list to readable list failed with: %s", method, err)
		}
	}

	if outputSize != 0 {
		err = codec.Convert(outputBuf).To(&outputs)
		if err != nil || len(outputs) == 0 {
			return nil, nil, fmt.Errorf("Converting contract method `%s` output list to readable list failed with: %s", method, err)
		}
	}

	return
}
