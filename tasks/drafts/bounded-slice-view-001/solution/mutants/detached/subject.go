package subject

import (
	"slices"
)

func CappedView(in []int, start, end int) []int { return slices.Clone(in[start:end]) }
