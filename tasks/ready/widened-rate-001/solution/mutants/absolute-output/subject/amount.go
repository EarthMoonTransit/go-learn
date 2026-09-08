package subject

import (
	"math"
)

func BandAmount(value float64) float64 { return math.Abs(value) * float64(BandRate(value)) / 100 }
