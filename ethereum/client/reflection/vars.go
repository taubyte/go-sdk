package reflection

import "github.com/taubyte/go-sdk/errno"

const (
	BytesMethod transactionMethodType = iota
	ByteConvertibleMethod
	BigIntMethod
	Uint64Method
)

const (
	TransactionNonceMethod     transactionMethod = "Nonce"
	TransactionGasPriceMethod  transactionMethod = "GasPrice"
	TransactionGasTipCapMethod transactionMethod = "GasTipCap"
	TransactionGasFeeCapMethod transactionMethod = "GasFeeCap"
	TransactionGasMethod       transactionMethod = "Gas"
	TransactionValueMethod     transactionMethod = "Value"
	TransactionDataMethod      transactionMethod = "Data"
	TransactionToAddressMethod transactionMethod = "To"
	TransactionChainMethod     transactionMethod = "ChainId"
	TransactionHashMethod      transactionMethod = "Hash"
)

var transactionMethodMap = map[string]transactionMethodDetail{
	TransactionNonceMethod.String():     {methodType: Uint64Method, errorType: errno.ErrorEthereumNonceNotFound},
	TransactionGasPriceMethod.String():  {methodType: BigIntMethod, errorType: errno.ErrorEthereumGasPriceNotFound},
	TransactionGasTipCapMethod.String(): {methodType: BigIntMethod, errorType: errno.ErrorEthereumGasTipCapNotFound},
	TransactionGasFeeCapMethod.String(): {methodType: BigIntMethod, errorType: errno.ErrorEthereumGasFeeCapNotFound},
	TransactionGasMethod.String():       {methodType: Uint64Method, errorType: errno.ErrorEthereumGasNotFound},
	TransactionValueMethod.String():     {methodType: BigIntMethod, errorType: errno.ErrorEthereumValueNotFound},
	TransactionDataMethod.String():      {methodType: BytesMethod, errorType: errno.ErrorEthereumDataNotFound},
	TransactionToAddressMethod.String(): {methodType: ByteConvertibleMethod, errorType: errno.ErrorEthereumAddressNotFound},
	TransactionChainMethod.String():     {methodType: BigIntMethod, errorType: errno.ErrorEthereumChainNotFound},
	TransactionHashMethod.String():      {methodType: ByteConvertibleMethod, errorType: errno.ErrorEthereumHashNotFound},
}
