package subject

import (
	"context"
	"time"
)

type Result struct {
	Value int
	Err   error
}

func RunBudgeted(parent context.Context, budget time.Duration, ceiling time.Time, operations []func(context.Context) (int, error)) []Result {
	results := make([]Result, 0, len(operations))
	for _, operation := range operations {
		base, releaseBase := context.WithTimeout(parent, budget)
		defer releaseBase()
		scope := base
		if !ceiling.IsZero() {
			var releaseCap context.CancelFunc
			scope, releaseCap = context.WithDeadline(base, ceiling)
			defer releaseCap()
		}
		value, err := operation(scope)
		results = append(results, Result{value, err})
	}
	return results
}
