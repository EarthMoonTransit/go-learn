package subject_test

import (
	"sync/atomic"
	subject "task.local/recover-worker-001/subject"
	"testing"
	"testing/synctest"
)

func drainWorker(t *testing.T, output <-chan subject.Result) []subject.Result {
	t.Helper()
	if output == nil {
		t.Error("Launch output=nil; want a nonnil result channel")
		return nil
	}
	var results []subject.Result
	for range 8 {
		synctest.Wait()
		select {
		case r, ok := <-output:
			if !ok {
				return results
			}
			results = append(results, r)
		default:
			t.Error("result channel remains open after the callback completed; want closure after one result")
			return results
		}
	}
	t.Error("result channel produced at least 8 results; want exactly one and then close")
	return results
}
func checkWorkerResult(t *testing.T, got []subject.Result, value int, panicText string, panicked bool) {
	t.Helper()
	if len(got) != 1 {
		t.Errorf("worker result count=%d; want exactly 1", len(got))
		return
	}
	r := got[0]
	if r.Value != value {
		t.Errorf("worker value=%d; want %d", r.Value, value)
	}
	if r.Panicked != panicked {
		t.Errorf("worker Panicked=%t; want %t", r.Panicked, panicked)
	}
	if panicked {
		if r.Err == nil {
			t.Errorf("worker error=nil; want %q", panicText)
		} else if r.Err.Error() != panicText {
			t.Errorf("worker error=%q; want %q", r.Err.Error(), panicText)
		}
	} else if r.Err != nil {
		t.Errorf("normal worker error=%v; want nil", r.Err)
	}
}
func TestLaunchReturnsWhileWorkerIsHeld(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		var calls atomic.Int32
		created := make(chan (<-chan subject.Result), 1)
		go func() { created <- subject.Launch(func() int { calls.Add(1); <-release; return -19 }) }()
		synctest.Wait()
		var output <-chan subject.Result
		returned := false
		select {
		case output = <-created:
			returned = true
		default:
			t.Error("Launch did not return while its worker callback was held")
		}
		if n := calls.Load(); n != 1 {
			t.Errorf("held callback calls=%d; want exactly 1", n)
		}
		var early []subject.Result
		if output != nil {
			select {
			case r, ok := <-output:
				if ok {
					early = append(early, r)
				}
				t.Error("worker output became readable before held callback completed")
			default:
			}
		}
		close(release)
		synctest.Wait()
		if !returned {
			select {
			case output = <-created:
			default:
				t.Error("Launch did not return after callback release")
				return
			}
		}
		got := append(early, drainWorker(t, output)...)
		checkWorkerResult(t, got, -19, "", false)
		if n := calls.Load(); n != 1 {
			t.Errorf("completed callback calls=%d; want exactly 1", n)
		}
	})
}
func TestPanicPolicyForStringAndIntegerValues(t *testing.T) {
	cases := []struct {
		name  string
		value any
		text  string
	}{
		{"text", "broken", "worker panic: broken"}, {"empty-text", "", "worker panic: "},
		{"integer", 42, "worker panic: 42"}, {"zero", 0, "worker panic: 0"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				var calls atomic.Int32
				output := subject.Launch(func() int { calls.Add(1); panic(tc.value) })
				checkWorkerResult(t, drainWorker(t, output), 0, tc.text, true)
				if n := calls.Load(); n != 1 {
					t.Errorf("panicking callback calls=%d; want exactly 1", n)
				}
			})
		})
	}
}
func TestIndependentNormalAndPanickingCalls(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		zero := subject.Launch(func() int { return 0 })
		failed := subject.Launch(func() int { panic("separate") })
		checkWorkerResult(t, drainWorker(t, failed), 0, "worker panic: separate", true)
		positive := subject.Launch(func() int { return 31 })
		checkWorkerResult(t, drainWorker(t, positive), 31, "", false)
		checkWorkerResult(t, drainWorker(t, zero), 0, "", false)
	})
}
