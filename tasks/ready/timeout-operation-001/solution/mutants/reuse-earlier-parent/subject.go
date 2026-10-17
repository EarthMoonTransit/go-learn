package subject

import (
	"context"
	"time"
)

func RunWithTimeout(parent context.Context, budget time.Duration, operation func(context.Context) (int, error)) (int, error) {
	if deadline, ok := parent.Deadline(); ok && deadline.Before(time.Now().Add(budget)) {
		return operation(parent)
	}
	ctx, cancel := context.WithTimeout(parent, budget)
	defer cancel()
	return operation(ctx)
}
