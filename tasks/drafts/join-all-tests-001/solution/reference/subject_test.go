package subject_test

import (
	"slices"
	"task.local/join-all-tests-001/subject"
	"testing"
	"testing/synctest"
)

func awaitJoined(t *testing.T, done <-chan struct{}) {
	t.Helper()
	if done == nil {
		t.Fatal("nil completion")
	}
	synctest.Wait()
	select {
	case _, ok := <-done:
		if ok {
			t.Error("completion sent a value")
		}
	default:
		t.Fatal("completion did not close")
	}
}
func TestContractCase01(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		calls := 0
		done := subject.JoinAll([]func(){func() { calls++ }})
		awaitJoined(t, done)
		if calls != 1 {
			t.Errorf("single callback calls%d", calls)
		}
	})
}

func TestContractCase02(t *testing.T) {
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
		counts := make([]int, 3)
		work := make([]func(), 3)
		for i := range work {
			work[i] = func() { counts[i]++; <-release }
		}
		returned := make(chan (<-chan struct{}), 1)
		go func() { returned <- subject.JoinAll(work) }()
		synctest.Wait()
		var done <-chan struct{}
		select {
		case done = <-returned:
		default:
			t.Fatal("JoinAll blocked before callbacks finished")
		}
		if done == nil {
			t.Fatal("nil completion")
		}
		if !slices.Equal(counts, []int{1, 1, 1}) {
			t.Errorf("callbacks not independently started: %v", counts)
		}
		select {
		case <-done:
			t.Error("completion before workers returned")
		default:
		}
		unblock()
		awaitJoined(t, done)
		if !slices.Equal(counts, []int{1, 1, 1}) {
			t.Errorf("callback counts%v want[1 1 1]", counts)
		}
	})
}

func TestContractCase03(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		a := subject.JoinAll(nil)
		b := subject.JoinAll([]func(){})
		if a == b {
			t.Error("independent completions shared")
		}
		awaitJoined(t, a)
		awaitJoined(t, b)
	})
}
