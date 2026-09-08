package subject

import (
	"context"
)

func CallWithContext(ctx context.Context, operation func(context.Context) (int, error)) (int, error) {
	operation(ctx)
	return operation(ctx)
}
