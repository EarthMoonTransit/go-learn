package subject

import (
	"strings"
)

func Quantity(values map[string]int, key string) (int, bool) {
	n, ok := values[strings.ToUpper(key)]
	return n, ok
}
