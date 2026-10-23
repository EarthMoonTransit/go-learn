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

	base, releaseBase := context.WithTimeout(parent, budget)
	defer releaseBase()
	scope := base
	if !ceiling.IsZero() {
		var releaseCap context.CancelFunc
		scope, releaseCap = context.WithDeadline(base, ceiling)
		defer releaseCap()
	}
	results := make([]Result, 0, len(operations))
	for i := len(operations) - 1; i >= 0; i-- {
		operation := operations[i]
		value, err := operation(scope)
		results = append(results, Result{Value: value, Err: err})
	}
	return results
}
