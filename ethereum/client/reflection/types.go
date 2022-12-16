package reflection

import "github.com/taubyte/go-sdk/errno"

type MethodDetail interface {
	Type() transactionMethodType
	Error() errno.Error
	IsBytesMethod() bool
	IsUint64Method() bool
}

type transactionMethodType uint32
type transactionMethodDetail struct {
	methodType transactionMethodType
	errorType  errno.Error
}

type transactionMethod string
