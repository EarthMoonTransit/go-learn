package subject

import (
	"context"
)

func CallWithContext(ctx context.Context, operation func(context.Context) (int, error)) (int, error) {
	n, e := operation(ctx)
	if e != nil {
		return 0, e
	}
	return n, e
}
