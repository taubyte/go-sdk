package codec

import (
	"fmt"
	"strings"

	"github.com/taubyte/go-sdk/utils/codec"
)

// Encoder method returns the encoder func, if given type is supported.
func (c converterType) Encoder() (encoder, error) {
	if converter, ok := converterMap[c]; ok {
		return converter.Encoder, nil
	}

	if strings.HasPrefix(c.String(), "[]") {
		return func(i interface{}) ([]byte, error) {
			ifaceList, ok := i.([]interface{})
			if !ok {
				return nil, fmt.Errorf("inputted data `%s` is not a list", i)
			}

			converter := Converter(c.String()[2:])
			encoder, err := converter.Encoder()
			if err != nil {
				return nil, fmt.Errorf("getting encoder for data type `%s` failed with: %s", converter, err)
			}

			encodedList := make([][]byte, 0)
			for _, iface := range ifaceList {
				encoded, err := encoder(iface)
				if err != nil {
					return nil, fmt.Errorf("encoding to slice element `%s` to buf failed with: %s", iface, err)
				}

				encodedList = append(encodedList, encoded)
			}

			var encoded []byte
			codec.Convert(encodedList).To(&encoded)

			return encoded, nil
		}, nil
	}

	return nil, fmt.Errorf("getting encoder method failed with: type %s not supported", c)
}
