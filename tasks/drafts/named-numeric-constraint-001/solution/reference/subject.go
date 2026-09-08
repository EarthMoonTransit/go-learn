package subject

import "math"

type supported interface{ ~int | ~int64 | ~float64 }
type numeric interface{ supported }

func Double[T numeric](value T) T {
	if math.IsNaN(float64(value)) {
		return value
	}
	return value + value
}
