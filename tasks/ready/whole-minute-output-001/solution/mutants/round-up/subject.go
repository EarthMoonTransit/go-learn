package subject

import (
	"math"
)

func SuccessfulPerMinute(rate int, percent float64) int {
	return int(math.Ceil(float64(rate) * percent / 100 / 60))
}
