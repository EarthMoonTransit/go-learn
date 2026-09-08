package subject

import (
	"math"
)

func FloatKind(value float64) string {
	if value == math.NaN() {
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
