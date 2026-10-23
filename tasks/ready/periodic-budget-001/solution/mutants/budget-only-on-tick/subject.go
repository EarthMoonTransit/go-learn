package subject

import "time"

func RunPeriodic(period, budget time.Duration, visit func(time.Duration, time.Duration) bool) bool {
	origin := time.Now()
	ticker := time.NewTicker(period)
	deadline := time.NewTimer(budget)
	defer ticker.Stop()
	defer deadline.Stop()
	for {
		select {
		case <-ticker.C:
			elapsed := time.Since(origin)
			if elapsed >= budget {
				return true
			}
			if !visit(elapsed, budget-elapsed) {
				return false
			}
		}
	}
}
