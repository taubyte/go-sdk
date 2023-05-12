//go:generate go run ./generate

package errno

type Error uint32

func (e Error) String() string {
	return errorStrings[e]
}

/***************************** ERRORS ****************************/

// When adding new errno, make sure to add to the string list
const (
	ErrorNone Error = iota
	ErrorEventNotFound
	ErrorBufferTooSmall
	ErrorAddressOutOfMemory
	ErrorNilAddress
	ErrorHttpWrite
	ErrorHttpReadBody
	ErrorCloseBody
	ErrorEOF
	ErrorReadHeaders
	ErrorClientNotFound
	ErrorParseUrlFailed
	ErrorMemoryWriteFailed
	ErrorHttpRequestFailed
	ErrorInvalidMethod
	ErrorNewRequestFailed
	ErrorHeaderNotFound
	ErrorHttpWriteBodyFailed
	ErrorCidNotFoundOnIpfs
	ErrorInvalidCid
	ErrorSubscribeFailed
	ErrorPublishFailed
	ErrorDatabaseCreateFailed
	ErrorDatabaseGetFailed
	ErrorDatabasePutFailed
	ErrorDatabaseNotFound
	ErrorDatabaseDeleteFailed
	ErrorDatabaseListFailed
	ErrorKeystoreCreateFailed
	ErrorDatabaseKeyNotFound
	ErrorKeystoreNotFound
	ErrorAddFileFailed
	ErrorGetFileFailed
	ErrorDeleteFileFailed
	ErrorCloseFileFailed
	ErrorFileNameNotFound
	ErrorListFileVersionsFailed
	ErrorListingUsedSpaceFailed
	ErrorGetWebSocketURLFailed
	ErrorByteConversionFailed
	ErrorChannelNotFound
	ErrorNewStreamFailed
	ErrorCommandCreateFailed
	ErrorP2PSendFailed
	ErrorP2PCommandNotFound
	ErrorP2PProtocolNotFound
	ErrorP2PFromNotFound
	ErrorP2PToNotFound
	ErrorP2PListenFailed
	ErrorMarshalDataFailed
	ErrorZeroSize
	ErrorEthereumNewClient
	ErrorEthereumBlockNotFound
	ErrorEthereumChainIdNotFound
	ErrorEthereumInvalidHexKey
	ErrorEthereumNonceNotFound
	ErrorEthereumGasPriceNotFound
	ErrorEthereumGasTipCapNotFound
	ErrorEthereumGasFeeCapNotFound
	ErrorEthereumGasNotFound
	ErrorEthereumValueNotFound
	ErrorEthereumDataNotFound
	ErrorEthereumAddressNotFound
	ErrorEthereumChainNotFound
	ErrorEthereumHashNotFound
	ErrorEthereumTransactionNotFound
	ErrorEthereumSendTransactionFailed
	ErrorEthereumMarshalJSON
	ErrorEthereumMethodNotSupported
	ErrorConvertibleConversionFailed
	ErrorEthereumContractNotFound
	ErrorEthereumParsingAbiFailed
	ErrorEthereumParsingECDSAFailed
	ErrorEthereumBindTransactorFailed
	ErrorEthereumCallContractFailed
	ErrorEthereumParseInputTypeFailed
	ErrorEthereumParseOutputTypeFailed
	ErrorEthereumContractMethodNotFound
	ErrorEthereumInvalidContractMethodInput
	ErrorEthereumInvalidContractMethodOutput
	ErrorEthereumUnsupportedDataType
	ErrorEthereumCannotCallPaidMutatorTransaction
	ErrorEthereumCannotTransactFreeMethod
	ErrorEthereumTransactMethodFailed
	ErrorEthereumDeployFailed
	ErrorEthereumSignFailed
	ErrorEthereumInvalidPublicKey
	ErrorEthereumInvalidPrivateKey
	ErrorEthereumRecoverPubKeyFailed
	ErrorSizeMismatch
	ErrorStorageGetMetaFailed
	ErrorAddFileToIpfsFailed
	ErrorStorageNotFound
	ErrorStorageListFailed
	ErrorCreatingNewFile
	ErrorWritingFile
	ErrorContentNotFound
	ErrorReadingFile
	ErrorInvalidWhence
	ErrorSeekingFile
	ErrorCidNotFound
	ErrorResolverNotFound
	ErrorFailedTxTLookup
	ErrorFailedAddressLookup
	ErrorFailedCNAMELookup
	ErrorFailedMXLookup
	ErrorCachedResponseTypeNotFound
	ErrorCachedResponseNotFound
	SmartOpErrorResourceNotFound
	SmartOpErrorWrongResourceInterface
	ErrorRandRead
	ErrorMemoryViewNotFound
	ErrorMemoryViewNotCloser
	ErrorSeekMethodNotFound
	ErrorInvalidBool
	ErrorFifoNotFound
	ErrorFifoDatatypeInvalid
	ErrorP2PDiscoverFailed
	ErrorEthereumWatchEventFailed
	ErrorEthereumRPCOptionUnmarshalFailed
	ErrorCap
)
