package subject

import (
	"slices"
)

func ScaledCopy(values []float64, units int) []float64 {
	out := make([]float64, len(values))
	for i, n := range values {
		out[i] = n * float64(units) / 2
	}
	slices.Reverse(out)
	return out
}
