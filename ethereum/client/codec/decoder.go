package codec

import (
	"fmt"
	"strings"

	"github.com/taubyte/go-sdk/utils/codec"
)

// Decoder method returns the decoder func if given type is supported.
func (c converterType) Decoder() (decoder, error) {
	if converter, ok := converterMap[c]; ok {
		return converter.Decoder, nil
	}

	if strings.HasPrefix(c.String(), "[]") {
		return func(b []byte) (interface{}, error) {
			byteList := [][]byte{}
			err := codec.Convert(b).To(byteList)
			if err != nil {
				return nil, fmt.Errorf("getting decoder method for slice failed with: converting to list failed with: %s", err)
			}

			var ifaceList []interface{}
			converter := Converter(c.String()[2:])
			decoder, err := converter.Decoder()
			if err != nil {
				return nil, fmt.Errorf("getting decoder method for slice failed with: %s", err)
			}

			for _, buf := range byteList {

				iface, err := decoder(buf)
				if err != nil {
					return nil, fmt.Errorf("decoding buf for slice failed with: %s", err)
				}

				ifaceList = append(ifaceList, iface)
			}

			return ifaceList, nil
		}, nil
	}

	return nil, fmt.Errorf("getting decoder method failed with: type %s not supported", c)
}
