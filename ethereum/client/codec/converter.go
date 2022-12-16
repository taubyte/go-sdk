package codec

// Converter takes the reflect type string and returns a corresponding convertType
func Converter(typeString string) converterType {
	return converterType(typeString)
}

func (c converterType) String() string {
	return string(c)
}
