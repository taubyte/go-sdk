package siwe

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/ethereum/siwe"
	"github.com/taubyte/go-sdk/utils/codec"
)

func InitMessage(domain, uri, address, nonce string, options map[string]string) (string, error) {
	if len(domain) == 0 || len(uri) == 0 || len(address) == 0 {
		return "", fmt.Errorf("Domain, uri, and address are required")
	}
	var encodedOptions []byte
	var optionsLen uint32
	if options != nil {
		err := codec.Convert(options).To(&encodedOptions)
		if err != nil {
			return "", fmt.Errorf("Encoding options failed with: %s", err)
		}
		optionsLen = uint32(len(encodedOptions))
	} else {
		encodedOptions = []byte{0}
	}
	var size uint32
	err0 := symbols.SiweInitMessageLen(domain, uri, address, nonce, &encodedOptions[0], optionsLen, &size)
	if err0 != 0 {
		return "", fmt.Errorf("Siwe message init failed with: %s", err0)
	}
	messageBytes := make([]byte, size)
	err0 = symbols.SiweInitMessage(domain, uri, address, nonce, &encodedOptions[0], optionsLen, &messageBytes[0])
	if err0 != 0 {
		return "", fmt.Errorf("Getting siwe message data failed with: %s", err0)
	}
	return string(messageBytes), nil
}
func VerifyEIP191(message string, signature []byte) error {
	err := symbols.SiweVerifyEIP191(message, &signature[0], uint32(len(signature)))
	if err != 0 {
		return fmt.Errorf("Verifying signature failed with: %s", err)
	}
	return nil
}
