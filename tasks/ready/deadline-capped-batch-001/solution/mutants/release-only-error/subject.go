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
	failed := false
	defer func() {
		if failed {
			releaseBase()
		}
	}()
	scope := base
	if !ceiling.IsZero() {
		var releaseCap context.CancelFunc
		scope, releaseCap = context.WithDeadline(base, ceiling)
		defer func() {
			if failed {
				releaseCap()
			}
		}()
	}
	results := make([]Result, 0, len(operations))
	for _, operation := range operations {
		value, err := operation(scope)
		if err != nil {
			failed = true
		}
		results = append(results, Result{Value: value, Err: err})
	}
	return results
}
