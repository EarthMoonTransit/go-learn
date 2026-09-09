package subject_test

import (
	"context"
	"errors"
	"task.local/cancel-worker-001/subject"
	"testing"
	"testing/synctest"
)

func awaitCanceled(t *testing.T, done <-chan struct{}) {
	t.Helper()
	synctest.Wait()
	select {
	case _, ok := <-done:
		if ok {
			t.Error("worker completion sent value")
		}
	default:
		t.Fatal("worker did not finish after cancellation")
	}
}
func TestPublicCase01(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, parentCancel := context.WithCancel(context.Background())
		defer parentCancel()
		child, cancel, done := subject.StartCancelable(parent)
		if child == nil || cancel == nil || done == nil {
			t.Fatal("nil child/cancel/completion")
		}
		defer cancel()
		synctest.Wait()
		if child.Err() != nil {
			t.Error("child canceled early")
		}
		select {
		case <-done:
			t.Error("worker completed before cancel")
		default:
		}
		cancel()
		cancel()
		awaitCanceled(t, done)
		if child.Err() != context.Canceled || parent.Err() != nil {
			t.Errorf("child cancellation child%v parent%v", child.Err(), parent.Err())
		}
	})
}

func TestPublicCase02(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		type key struct{}
		parent, parentCancel := context.WithCancelCause(context.WithValue(context.Background(), key{}, "value"))
		cause := errors.New("parent stopped")
		defer parentCancel(cause)
		child, cancel, done := subject.StartCancelable(parent)
		if child == nil || cancel == nil || done == nil {
			t.Fatal("nil result")
		}
		defer cancel()
		if child.Value(key{}) != "value" {
			t.Error("parent value lost")
		}
		parentCancel(cause)
		awaitCanceled(t, done)
		if child.Err() != context.Canceled || context.Cause(child) != cause {
			t.Errorf("parent cancel child%v cause%v", child.Err(), context.Cause(child))
		}
	})
}

func TestPublicCase03(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, stop := context.WithCancel(context.Background())
		stop()
		child, cancel, done := subject.StartCancelable(parent)
		if child == nil || cancel == nil || done == nil {
			t.Fatal("nil result")
		}
		defer cancel()
		awaitCanceled(t, done)
		if child.Err() != context.Canceled {
			t.Error("already canceled parent ignored")
		}
	})
}
