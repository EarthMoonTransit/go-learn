package subject

import (
	"slices"
)

func ReplaceFirst(replacement int, values ...int) bool {
	if len(values) == 0 {
		return false
	}
	values = slices.Clone(values)
	values[0] = replacement
	return true
}
