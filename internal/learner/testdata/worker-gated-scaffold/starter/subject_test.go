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

// gatedWork is supplied test scaffolding. Release is idempotent so every test
// can defer it, even when an assertion stops the test early.
type gatedWork struct {
	gate     chan struct{}
	released bool
	calls    int
}

func newGatedWork() *gatedWork { return &gatedWork{gate: make(chan struct{})} }
func (work *gatedWork) run()   { work.calls++; <-work.gate }
func (work *gatedWork) release() {
	if !work.released {
		close(work.gate)
		work.released = true
	}
}

func TestWorkerBlockedPhase(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		work := newGatedWork()
		defer func() { work.release(); synctest.Wait() }()
		returned := make(chan (<-chan struct{}), 1)
		go func() { returned <- subject.Start(work.run) }()

		// At this boundary runnable work has either returned or reached its gate.
		synctest.Wait()
		var done <-chan struct{}
		select {
		case done = <-returned:
		default:
			t.Error("blocked phase: Start has not returned, want return before release")
			work.release()
			synctest.Wait()
			select {
			case done = <-returned:
			default:
				t.Fatal("released phase: Start has not returned, want completion handle")
			}
		}
		if work.calls != 1 {
			t.Errorf("blocked phase callback calls = %d, want 1", work.calls)
		}
		work.release()
		awaitClosed(t, done)
		if work.calls != 1 {
			t.Errorf("completed phase callback calls = %d, want 1", work.calls)
		}
	})
}

// Add your own premature-completion and overlapping independent-call cases.
// Use synctest.Test/Wait boundaries and release every gate on all failure paths.
