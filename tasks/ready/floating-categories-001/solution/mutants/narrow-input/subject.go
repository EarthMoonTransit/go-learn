package subject

import (
	"math"
)

func FloatKind(value float64) string {
	value = float64(float32(value))
	if math.IsNaN(value) {
		return "nan"
	}
	if math.IsInf(value, 1) {
		return "positive-infinity"
	}
	if math.IsInf(value, -1) {
		return "negative-infinity"
	}
	return "finite"
}
