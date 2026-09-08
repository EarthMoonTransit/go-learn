package subject_test

import (
	"context"
	"errors"
	"task.local/propagate-context-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	type key struct{}
	ctx := context.WithValue(context.Background(), key{}, "v")
	cause := errors.New("operation failed")
	calls := 0
	n, e := subject.CallWithContext(ctx, func(got context.Context) (int, error) {
		calls++
		if got != ctx || got.Value(key{}) != "v" {
			t.Error("context identity/value lost")
		}
		return 7, cause
	})
	if n != 7 || e != cause || calls != 1 {
		t.Errorf("result got%d,%v calls%d", n, e, calls)
	}
}

func TestPublicCase02(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	calls := 0
	n, e := subject.CallWithContext(ctx, func(got context.Context) (int, error) {
		calls++
		if got != ctx || got.Err() != context.Canceled {
			t.Error("canceled context lost")
		}
		return -2, nil
	})
	if n != -2 || e != nil || calls != 1 {
		t.Errorf("canceled operation got%d,%v calls%d", n, e, calls)
	}
}

func TestPublicCase03(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if n, e := subject.CallWithContext(ctx, func(context.Context) (int, error) { return 0, nil }); n != 0 || e != nil || ctx.Err() != nil {
		t.Error("successful composition changed context")
	}
}
