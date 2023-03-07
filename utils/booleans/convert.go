package booleans

func FromBool(val bool) (parsed uint32) {
	if val {
		parsed = 1
	}

	return
}

func ToBool(val uint32) (parsed bool) {
	if val == 1 {
		parsed = true
	}

	return parsed
}
