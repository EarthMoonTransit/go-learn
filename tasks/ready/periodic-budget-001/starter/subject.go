package subject

import "time"

func RunPeriodic(period, budget time.Duration, visit func(time.Duration, time.Duration) bool) bool {
	return false
}
