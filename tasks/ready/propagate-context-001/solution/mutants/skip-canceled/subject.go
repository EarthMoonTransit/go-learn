package subject

import (
	"context"
)

func CallWithContext(ctx context.Context, operation func(context.Context) (int, error)) (int, error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return operation(ctx)
}
