package slices

func MakeByteList(sizes ...uint32) [][]byte {
	list := [][]byte{}
	for _, size := range sizes {
		if size == 0 {
			size = 1
		}
		list = append(list, make([]byte, size))
	}

	return list
}
