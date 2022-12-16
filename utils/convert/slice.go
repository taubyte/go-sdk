package convert

func Count(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	var result = 1
	for _, d := range data {
		if d == 0 {
			result++
		}

	}
	return result
}
