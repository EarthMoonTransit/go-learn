package subject

import "time"

func RunPeriodic(period, budget time.Duration, visit func(time.Duration, time.Duration) bool) bool {
	_ = time.Now()
	ticker := time.NewTicker(period)
	deadline := time.NewTimer(budget)
	defer ticker.Stop()
	defer deadline.Stop()
	var ticks int64
	for {
		select {
		case <-deadline.C:
			return true
		case <-ticker.C:
			ticks++
			elapsed := time.Duration(ticks) * period
			if elapsed >= budget {
				return true
			}
			if !visit(elapsed, budget-elapsed) {
				return false
			}
		}
	}
}
