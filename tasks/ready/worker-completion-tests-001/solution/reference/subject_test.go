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

func TestWorker1(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		released := false
		unblock := func() {
			if !released {
				close(release)
				released = true
			}
		}
		defer unblock()
		started := make(chan struct{}, 2)
		calls := 0
		returned := make(chan (<-chan struct{}), 1)
		go func() { returned <- subject.Start(func() { calls++; started <- struct{}{}; <-release }) }()
		synctest.Wait()
		var done <-chan struct{}
		select {
		case done = <-returned:
		default:
			t.Fatal("Start blocked waiting for callback")
		}
		if done == nil {
			t.Fatal("nil completion")
		}
		select {
		case <-started:
		default:
			t.Fatal("callback did not start")
		}
		select {
		case _, ok := <-done:
			t.Errorf("completion before release, open=%v", ok)
		default:
		}
		unblock()
		awaitClosed(t, done)
		if calls != 1 {
			t.Errorf("controlled callback calls%d want1", calls)
		}
	})
}

func TestWorker2(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		released := false
		unblock := func() {
			if !released {
				close(release)
				released = true
			}
		}
		defer unblock()
		a, b := 0, 0
		first := subject.Start(func() { a++; <-release })
		second := subject.Start(func() { b++; <-release })
		synctest.Wait()
		if a != 1 || b != 1 {
			t.Errorf("overlapping callbacks %d,%d", a, b)
		}
		if first == second {
			t.Error("active calls share completion channel")
		}
		unblock()
		awaitClosed(t, first)
		awaitClosed(t, second)
	})
}
