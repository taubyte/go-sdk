package ethereum

import (
	"fmt"
	"math/big"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	"github.com/taubyte/go-sdk/ethereum/client/bytes"
	"github.com/taubyte/go-sdk/ethereum/client/reflection"
	"github.com/taubyte/go-sdk/utils/ints"
	"github.com/taubyte/go-sdk/utils/slices"
)

func (t *Transaction) callBytesMethod(method string) ([]byte, error) {
	var size uint32
	err := ethereumSym.EthGetTransactionMethodSize(uint32(t.client), &t.blockID, t.contractID, t.id, method, &size)
	if err != 0 {
		return nil, fmt.Errorf("getting size failed with: %s", err)
	}

	if size != 0 {
		buf := make([]byte, size)
		err = ethereumSym.EthGetTransactionMethodBytes(uint32(t.client), &t.blockID, t.contractID, t.id, method, &buf[0])
		if err != 0 {
			return nil, fmt.Errorf("getting bytes buffer failed with: %s", err)
		}

		return buf, nil
	}

	return nil, nil
}

// Nonce returns the sender account nonce of the transaction.
func (t *Transaction) Nonce() (uint64, error) {
	if t.nonce != 0 {
		return t.nonce, nil
	}

	if err := ethereumSym.EthGetTransactionMethodUint64(uint32(t.client), &t.blockID, t.contractID, t.id, reflection.TransactionNonceMethod.String(), &t.nonce); err != 0 {
		return 0, fmt.Errorf("getting transaction nonce failed with: %s", err)
	}
	return t.nonce, nil
}

// GasPrice returns the gas price of the transaction
func (t *Transaction) GasPrice() (*big.Int, error) {
	if t.gasPrice != nil {
		return t.gasPrice, nil
	}

	buf, err := t.callBytesMethod(reflection.TransactionGasPriceMethod.String())
	if err != nil {
		return nil, fmt.Errorf("getting gas price failed with: %s", err)
	}

	t.gasPrice = ints.NewBigInt(buf)
	return t.gasPrice, nil
}

// GasTipCap returns the gasTipCap per gas of the transaction.
func (t *Transaction) GasTipCap() (*big.Int, error) {
	if t.gasTipCap != nil {
		return t.gasTipCap, nil
	}

	buf, err := t.callBytesMethod(reflection.TransactionGasTipCapMethod.String())
	if err != nil {
		return nil, fmt.Errorf("getting gas tip cap failed with: %s", err)
	}

	t.gasTipCap = ints.NewBigInt(buf)
	return t.gasTipCap, nil
}

// GasFeeCap returns the fee cap per gas of the transaction.
func (t *Transaction) GasFeeCap() (*big.Int, error) {
	if t.gasFeeCap != nil {
		return t.gasFeeCap, nil
	}

	buf, err := t.callBytesMethod(reflection.TransactionGasFeeCapMethod.String())
	if err != nil {
		return nil, fmt.Errorf("getting gas fee cap failed with: %s", err)
	}

	t.gasFeeCap = ints.NewBigInt(buf)
	return t.gasFeeCap, nil
}

// Gas returns the gas limit of the transaction.
func (t *Transaction) Gas() (uint64, error) {
	if t.gas != 0 {
		return t.gas, nil
	}

	if err := ethereumSym.EthGetTransactionMethodUint64(uint32(t.client), &t.blockID, t.contractID, t.id, reflection.TransactionGasMethod.String(), &t.gas); err != 0 {
		return 0, fmt.Errorf("getting transaction gas failed with: %s", err)
	}

	return t.gas, nil
}

// Value returns the ether amount of the transaction.
func (t *Transaction) Value() (*big.Int, error) {
	if t.value != nil {
		return t.value, nil
	}

	buf, err := t.callBytesMethod(reflection.TransactionValueMethod.String())
	if err != nil {
		return nil, fmt.Errorf("getting transaction value price failed with: %s", err)
	}

	t.value = ints.NewBigInt(buf)
	return t.value, nil
}

// Data returns the input data of the transaction.
func (t *Transaction) Data() ([]byte, error) {
	if len(t.data) != 0 {
		return t.data, nil
	}

	buf, err := t.callBytesMethod(reflection.TransactionDataMethod.String())
	if err != nil {
		return nil, fmt.Errorf("getting gas price failed with: %s", err)
	}

	t.data = buf
	return t.data, nil
}

// RawSignatures returns the V, R, S signature values of the transaction. The return values should not be modified by the caller.
func (t *Transaction) RawSignatures() (rawSignatures, error) {
	if t.rawSignatures.VSig != nil && t.rawSignatures.SSig != nil && t.rawSignatures.RSig != nil {
		return t.rawSignatures, nil
	}

	sizes := make([]uint32, 3)
	if err := ethereumSym.EthTransactionRawSignaturesSize(uint32(t.client), &t.blockID, t.contractID, t.id, &sizes[0], &sizes[1], &sizes[2]); err != 0 {
		return t.rawSignatures, fmt.Errorf("getting transaction signatures failed with: %s", err)
	}

	bufList := slices.MakeByteList(sizes...)
	if err0 := ethereumSym.EthTransactionRawSignatures(uint32(t.client), &t.blockID, t.contractID, t.id, &bufList[0][0], &bufList[1][0], &bufList[2][0]); err0 != 0 {
		return t.rawSignatures, fmt.Errorf("getting nonce from transaction failed with: %s", err0)
	}

	t.rawSignatures = rawSignatures{
		VSig: new(big.Int).SetBytes(bufList[0]),
		RSig: new(big.Int).SetBytes(bufList[1]),
		SSig: new(big.Int).SetBytes(bufList[2]),
	}
	return t.rawSignatures, nil
}

// ToAddress returns the recipient address of the transaction.
func (t *Transaction) ToAddress() ([]byte, error) {
	if len(t.toAddress) != 0 {
		return t.toAddress, nil
	}

	buf := make([]byte, bytes.AddressByteLength)
	err := ethereumSym.EthGetTransactionMethodBytes(uint32(t.client), &t.blockID, t.contractID, t.id, reflection.TransactionToAddressMethod.String(), &buf[0])
	if err != 0 {
		return nil, fmt.Errorf("getting transaction recipient address failed with: %s", err)
	}

	t.toAddress = buf
	return buf, nil
}

// Chain returns the EIP155 chain ID of the transaction. The return value will always be
// non-nil. For legacy transactions which are not replay-protected, the return value is
// zero.
func (t *Transaction) Chain() (*big.Int, error) {
	if t.chain != nil {
		return t.chain, nil
	}

	buf, err := t.callBytesMethod(reflection.TransactionChainMethod.String())
	if err != nil {
		return nil, fmt.Errorf("getting transaction chain Id failed with: %s", err)
	}

	t.chain = ints.NewBigInt(buf)
	return t.chain, nil
}

// Hash returns the transaction hash.
func (t *Transaction) Hash() ([]byte, error) {
	if len(t.hash) != 0 {
		return t.hash, nil
	}
	buf, err := t.callBytesMethod(reflection.TransactionHashMethod.String())
	if err != nil {
		return nil, fmt.Errorf("getting transaction value price failed with: %s", err)
	}

	t.hash = buf
	return t.hash, nil
}

// Send injects a signed transaction into the pending pool for execution.
func (t *Transaction) Send() (err error) {
	if err0 := ethereumSym.EthSendTransaction(uint32(t.client), &t.blockID, t.contractID, t.id); err0 != 0 {
		err = fmt.Errorf("sending transaction failed with: %s", err)
	}

	return
}
