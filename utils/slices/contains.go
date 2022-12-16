package slices

func Contains[T comparable](slice []T, value T) bool {
	for _, s := range slice {
		if value == s {
			return true
		}
	}

	return false
}
