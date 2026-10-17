package subject

import (
	"context"
	"time"
)

func RunWithTimeout(parent context.Context, budget time.Duration, operation func(context.Context) (int, error)) (int, error) {
	if parent.Err() != nil {
		return 0, parent.Err()
	}
	ctx, cancel := context.WithTimeout(parent, budget)
	defer cancel()
	return operation(ctx)
}
