package subject

import (
	"math"
)

func BandRate(value float64) float32 {
	value = math.Abs(value)
	if value < 0 {
		return 3.213
	}
	if value < 1000 {
		return 0.5
	}
	if value < 5000 {
		return 1.621
	}
	return 2.475
}
