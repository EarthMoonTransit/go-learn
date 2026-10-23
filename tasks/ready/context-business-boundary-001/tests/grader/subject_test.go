package subject_test

import (
	"context"
	subject "task.local/context-business-boundary-001/subject"
	"testing"
)

func TestProductGrid(t *testing.T) {
	for v := -8; v <= 8; v++ {
		for m := -8; m <= 8; m++ {
			calls := 0
			ctx := context.WithValue(context.Background(), "multiplier", 8)
			err := subject.Apply(ctx, v, m, func(_ context.Context, got int) error {
				calls++
				if got != v*m {
					t.Errorf("product(%d,%d)=%d", v, m, got)
				}
				return nil
			})
			if err != nil || calls != 1 {
				t.Errorf("calls=%d err=%v", calls, err)
			}
		}
	}
}
