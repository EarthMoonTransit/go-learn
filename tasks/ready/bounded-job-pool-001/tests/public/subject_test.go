package subject_test

import (
	"errors"
	"sync/atomic"
	subject "task.local/bounded-job-pool-001/subject"
	"testing"
	"testing/synctest"
)

func checkPoolResults(t *testing.T, got []subject.Result, values []int, errs []error) {
	t.Helper()
	if len(got) != len(values) {
		t.Errorf("result count=%d; want one result for each of %d jobs", len(got), len(values))
	}
	seen := make([]bool, len(values))
	for _, r := range got {
		if r.Index < 0 || r.Index >= len(values) {
			t.Errorf("result index=%d; want index in [0,%d)", r.Index, len(values))
			continue
		}
		if seen[r.Index] {
			t.Errorf("duplicate result index=%d; want every job index exactly once", r.Index)
		}
		seen[r.Index] = true
		if r.Value != values[r.Index] {
			t.Errorf("job %d value=%d; want %d, including when an error is present", r.Index, r.Value, values[r.Index])
		}
		if r.Err != errs[r.Index] {
			t.Errorf("job %d error=%v; want exact callback error %v", r.Index, r.Err, errs[r.Index])
		}
	}
	for i, ok := range seen {
		if !ok {
			t.Errorf("missing result for job %d", i)
		}
	}
}
func heldPool(t *testing.T, workers int, values []int, errs []error) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		var active atomic.Int32
		calls := make([]atomic.Int32, len(values))
		jobs := make([]func() (int, error), len(values))
		for i := range jobs {
			jobs[i] = func() (int, error) {
				calls[i].Add(1)
				active.Add(1)
				defer active.Add(-1)
				<-release
				return values[i], errs[i]
			}
		}
		finished := make(chan []subject.Result, 1)
		go func() { finished <- subject.RunPool(jobs, workers) }()
		synctest.Wait()
		wantActive := min(workers, len(jobs))
		if got := active.Load(); got != int32(wantActive) {
			t.Errorf("active callbacks while held=%d; want %d for workers=%d jobs=%d", got, wantActive, workers, len(jobs))
		}
		var got []subject.Result
		returned := false
		select {
		case got = <-finished:
			returned = true
			if len(jobs) > 0 {
				t.Error("RunPool returned before held callbacks completed")
			}
		default:
		}
		close(release)
		synctest.Wait()
		if !returned {
			select {
			case got = <-finished:
				returned = true
			default:
				t.Error("RunPool did not finish after all callbacks were released")
			}
		}
		if !returned {
			return
		}
		if n := active.Load(); n != 0 {
			t.Errorf("active callbacks after RunPool returned=%d; want 0", n)
		}
		for i := range calls {
			if n := calls[i].Load(); n != 1 {
				t.Errorf("job %d callback calls=%d; want exactly 1 even after an earlier error", i, n)
			}
		}
		checkPoolResults(t, got, values, errs)
	})
}
func TestWorkersStartTogetherWithinBound(t *testing.T) {
	marker := errors.New("callback failure")
	heldPool(t, 3, []int{14, 0, -8, 21, 4, -2}, []error{marker, nil, marker, nil, nil, nil})
}
func TestSingleWorkerContinuesAfterError(t *testing.T) {
	marker := errors.New("first callback failed")
	heldPool(t, 1, []int{17, -3, 0}, []error{marker, nil, nil})
}
func TestFewerJobsThanWorkers(t *testing.T) {
	heldPool(t, 4, []int{-5, 8}, []error{nil, nil})
}
func TestEmptyPool(t *testing.T) { heldPool(t, 4, nil, nil) }
