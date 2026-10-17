package subject

import (
	"context"
	"time"
)

func RunWithTimeout(parent context.Context, budget time.Duration, operation func(context.Context) (int, error)) (int, error) {
	ctx, cancel := context.WithTimeout(parent, budget)
	defer cancel()
	type result struct {
		n   int
		err error
	}
	finished := make(chan result, 1)
	go func() { n, err := operation(ctx); finished <- result{n, err} }()
	select {
	case value := <-finished:
		return value.n, value.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
