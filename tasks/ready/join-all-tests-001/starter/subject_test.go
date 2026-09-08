package subject_test

import (
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
func TestStartingCase01(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		calls := 0
		done := subject.JoinAll([]func(){func() { calls++ }})
		awaitJoined(t, done)
		if calls != 1 {
			t.Errorf("single callback calls%d", calls)
		}
	})
}

// Add cases for the remaining requirements in TASK.md.
