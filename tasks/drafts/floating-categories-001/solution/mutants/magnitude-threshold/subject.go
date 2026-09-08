package subject

import (
	"math"
)

func FloatKind(value float64) string {
	if math.Abs(value) > 1e308 {
		return "positive-infinity"
	}
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
