package storage

import "strings"

func trimInvisibleChar(entry string) string {
	return strings.Trim(entry, "\x00")
}
