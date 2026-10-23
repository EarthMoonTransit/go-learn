package subject_test

import (
	"errors"
	"sync/atomic"
	subject "task.local/stream-job-outcomes-001/subject"
	"testing"
	"testing/synctest"
)

func collect(t *testing.T, output <-chan subject.Result) []subject.Result {
	t.Helper()
	if output == nil {
		t.Error("StreamJobs output=nil; want nonnil result channel")
		return nil
	}
	var all []subject.Result
	for range 40 {
		synctest.Wait()
		select {
		case item, ok := <-output:
			if !ok {
				return all
			}
			all = append(all, item)
		default:
			t.Error("output still open after released callbacks and drained results; want closure")
			return all
		}
	}
	t.Error("result stream exceeded finite observation allowance")
	return all
}
func outcomes(t *testing.T, got, want []subject.Result) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("result count=%d; want %d", len(got), len(want))
	}
	seen := make([]bool, len(want))
	for _, r := range got {
		if r.Index < 0 || r.Index >= len(want) {
			t.Errorf("result index=%d outside job range", r.Index)
			continue
		}
		if seen[r.Index] {
			t.Errorf("duplicate result index=%d", r.Index)
		}
		seen[r.Index] = true
		w := want[r.Index]
		if r.Value != w.Value {
			t.Errorf("job %d Value=%d; want %d", r.Index, r.Value, w.Value)
		}
		if r.Err != w.Err {
			t.Errorf("job %d Err=%v; want exact error %v", r.Index, r.Err, w.Err)
		}
		if r.Panicked != w.Panicked {
			t.Errorf("job %d Panicked=%t; want %t", r.Index, r.Panicked, w.Panicked)
		}
		if r.Panic != w.Panic {
			t.Errorf("job %d Panic=%#v; want exact panic value %#v", r.Index, r.Panic, w.Panic)
		}
	}
	for i, ok := range seen {
		if !ok {
			t.Errorf("missing result for job %d", i)
		}
	}
}
func heldJobs(t *testing.T, workers, n int) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		var active atomic.Int32
		calls := make([]atomic.Int32, n)
		jobs := make([]func() (int, error), n)
		want := make([]subject.Result, n)
		marker := errors.New("ordinary callback error")
		for i := range jobs {
			want[i] = subject.Result{Index: i, Value: i*5 - 7}
			if i%2 == 0 {
				want[i].Err = marker
			}
			jobs[i] = func() (int, error) {
				calls[i].Add(1)
				active.Add(1)
				defer active.Add(-1)
				<-release
				return want[i].Value, want[i].Err
			}
		}
		created := make(chan (<-chan subject.Result), 1)
		go func() { created <- subject.StreamJobs(jobs, workers) }()
		synctest.Wait()
		var output <-chan subject.Result
		returned := false
		select {
		case output = <-created:
			returned = true
		default:
			t.Error("StreamJobs did not return output while callbacks were held")
		}
		if a := active.Load(); a != int32(min(workers, n)) {
			t.Errorf("active held callbacks=%d; want %d", a, min(workers, n))
		}
		var early []subject.Result
		if output != nil && n > 0 {
			select {
			case r, ok := <-output:
				if ok {
					early = append(early, r)
				}
				t.Error("output became readable before held callbacks completed")
			default:
			}
		}
		close(release)
		synctest.Wait()
		if !returned {
			select {
			case output = <-created:
			default:
				t.Error("StreamJobs did not return after all callback releases")
			}
		}
		got := append(early, collect(t, output)...)
		synctest.Wait()
		if a := active.Load(); a != 0 {
			t.Errorf("active callbacks after drain=%d; want 0", a)
		}
		for i := range calls {
			if n := calls[i].Load(); n != 1 {
				t.Errorf("job %d calls=%d; want 1", i, n)
			}
		}
		outcomes(t, got, want)
	})
}
func TestPublicHeldWorkersAndPromptOutput(t *testing.T) {
	heldJobs(t, 2, 5)
	heldJobs(t, 3, 2)
	heldJobs(t, 3, 0)
}
func TestPublicCompletedResultStreamsBeforeOtherJob(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		jobs := []func() (int, error){func() (int, error) { <-release; return 17, nil }, func() (int, error) { return 43, nil }}
		created := make(chan (<-chan subject.Result), 1)
		go func() { created <- subject.StreamJobs(jobs, 2) }()
		synctest.Wait()
		var output <-chan subject.Result
		returned := false
		select {
		case output = <-created:
			returned = true
		default:
			t.Error("output handle unavailable before unrelated callback release")
		}
		var early []subject.Result
		select {
		case r, ok := <-output:
			if ok {
				early = append(early, r)
				if r.Index != 1 {
					t.Errorf("early result index=%d; want completed job 1", r.Index)
				}
			} else {
				t.Error("output closed while job 0 remained held")
			}
		default:
			t.Error("completed job 1 not published while unrelated job 0 remained held")
		}
		close(release)
		synctest.Wait()
		if !returned {
			select {
			case output = <-created:
			default:
				t.Error("output handle unavailable after release")
			}
		}
		outcomes(t, append(early, collect(t, output)...), []subject.Result{{Index: 0, Value: 17}, {Index: 1, Value: 43}})
	})
}
func panicSequence(t *testing.T, panics []any) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		marker := errors.New("normal error after panic")
		jobs := make([]func() (int, error), 2*len(panics))
		want := make([]subject.Result, len(jobs))
		calls := make([]atomic.Int32, len(jobs))
		for i, value := range panics {
			p, n := 2*i, 2*i+1
			want[p] = subject.Result{Index: p, Panicked: true, Panic: value}
			want[n] = subject.Result{Index: n, Value: 11 - i, Err: marker}
			jobs[p] = func() (int, error) { calls[p].Add(1); panic(value) }
			jobs[n] = func() (int, error) { calls[n].Add(1); return 11 - i, marker }
		}
		output := subject.StreamJobs(jobs, 1)
		got := collect(t, output)
		for i := range calls {
			if n := calls[i].Load(); n != 1 {
				t.Errorf("job %d calls=%d after earlier panic; want 1", i, n)
			}
		}
		outcomes(t, got, want)
	})
}
func TestPublicWorkerContinuesAfterEveryPanic(t *testing.T) {
	panicSequence(t, []any{errors.New("raw panic identity"), 0, ""})
}
func TestPublicIndependentCalls(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		first := subject.StreamJobs([]func() (int, error){func() (int, error) { panic("first") }}, 1)
		outcomes(t, collect(t, first), []subject.Result{{Index: 0, Panicked: true, Panic: "first"}})
		second := subject.StreamJobs([]func() (int, error){func() (int, error) { return 0, nil }}, 1)
		outcomes(t, collect(t, second), []subject.Result{{Index: 0}})
	})
}
