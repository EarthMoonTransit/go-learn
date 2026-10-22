package subject

import "time"

func UntilIdle(input <-chan int, quiet time.Duration) (values []int, exhausted bool) {
	timer := time.NewTimer(quiet)
	defer timer.Stop()
	for {
		select {
		case value, ok := <-input:
			if !ok {
				return values, true
			}
			values = append(values, value)

		case <-timer.C:
			return values, false
		}
	}
}
