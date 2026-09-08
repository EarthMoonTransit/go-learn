package subject

import (
	"slices"
)

func DetachedPrefix(in []int, n int) []int {
	out := make([]int, n)
	copy(out, in[:n])
	slices.Reverse(out)
	return out
}
