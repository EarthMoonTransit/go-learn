package subject_test

import (
	"task.local/worker-completion-tests-001/subject"
	"testing"
	"testing/synctest"
)

func awaitClosed(t *testing.T, done <-chan struct{}) {
	t.Helper()
	if done == nil {
		t.Fatal("nil completion")
	}
	synctest.Wait()
	select {
	case _, ok := <-done:
		if ok {
			t.Error("completion sent a value instead of closing")
		}
	default:
		t.Fatal("completion channel did not close")
	}
}
func TestWorker0(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		calls := 0
		done := subject.Start(func() { calls++ })
		awaitClosed(t, done)
		if calls != 1 {
			t.Errorf("callback calls%d want1", calls)
		}
	})
}

// Add controlled blocked-work and independent-call cases inside synctest.Test.
