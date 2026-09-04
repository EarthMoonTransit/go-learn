package subject

import "math"

func EffectiveRate(itemsPerHour int, successPercent float64) int {
	return int(math.Round(float64(itemsPerHour) * successPercent / 100))
}
