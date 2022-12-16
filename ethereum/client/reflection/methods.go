package reflection

import (
	"fmt"

	"github.com/taubyte/go-sdk/errno"
)

func (t transactionMethod) String() string {
	return string(t)
}

// ReflectiveTransaction takes a method name for corresponding to a Transaction
// object's allowed methods.
// Returns method detail containing methods defining internal usage handling.
func ReflectiveTransaction(method string) (MethodDetail, error) {
	methodDetail, ok := transactionMethodMap[method]
	if ok == false {
		return nil, fmt.Errorf("Method `%s` is unsupported", method)
	}

	return methodDetail, nil
}

// IsBytesMethod method returns bool defining whether given transaction method returns []byte.
func (t transactionMethodDetail) IsBytesMethod() bool {
	if t.methodType == BytesMethod || t.methodType == ByteConvertibleMethod || t.methodType == BigIntMethod {
		return true
	}

	return false
}

// IsUint64Method method returns bool defining whether given transaction method returns a uint64.
func (t transactionMethodDetail) IsUint64Method() bool {
	if t.methodType == Uint64Method {
		return true
	}

	return false
}

// Type method returns the data type value returned by the transaction method.
func (t transactionMethodDetail) Type() transactionMethodType {
	return t.methodType
}

// Error returns the errno.Error value for the given transaction method.
func (t transactionMethodDetail) Error() errno.Error {
	return t.errorType
}
