package subject

import (
	"context"
	"time"
)

func RunWithTimeout(parent context.Context, budget time.Duration, operation func(context.Context) (int, error)) (int, error) {
	ctx, cancel := context.WithTimeout(parent, budget)
	defer cancel()
	n, err := operation(ctx)
	if err != nil {
		return n, ctx.Err()
	}
	return n, nil
}
