package subject_test

import (
	"errors"
	"sync/atomic"
	subject "task.local/indexed-job-results-001/subject"
	"testing"
	"testing/synctest"
)

func p17AssertRows(t *testing.T, label string, got, want []subject.Result) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s result length = %d, want %d records", label, len(got), len(want))
		return
	}
	for i := range want {
		if got[i].Index != want[i].Index {
			t.Errorf("%s result[%d].Index = %d, want %d", label, i, got[i].Index, want[i].Index)
		}
		if got[i].Value != want[i].Value {
			t.Errorf("%s result[%d].Value = %d, want %d", label, i, got[i].Value, want[i].Value)
		}
		if got[i].Err != want[i].Err {
			t.Errorf("%s result[%d] error identity differs: got %T (%v), want original job error %T (%v)", label, i, got[i].Err, got[i].Err, want[i].Err, want[i].Err)
		}
	}
}

func p17Jobs(rows []subject.Result) ([]func() (int, error), []atomic.Int32) {
	calls := make([]atomic.Int32, len(rows))
	jobs := make([]func() (int, error), len(rows))
	for i, row := range rows {
		jobs[i] = func() (int, error) { calls[i].Add(1); return row.Value, row.Err }
	}
	return jobs, calls
}

func p17AssertCalls(t *testing.T, label string, calls []atomic.Int32) {
	t.Helper()
	for i := range calls {
		if got := calls[i].Load(); got != 1 {
			t.Errorf("%s job[%d] calls = %d, want exactly 1", label, i, got)
		}
	}
}

func TestEveryIndexedValueAndError(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		first, second := errors.New("same error text"), errors.New("same error text")
		want := []subject.Result{{Index: 0, Value: 35}, {Index: 1, Value: -9, Err: first}, {Index: 2, Value: 0}, {Index: 3, Value: 73, Err: second}}
		jobs, calls := p17Jobs(want)
		got := subject.CollectResults(jobs)
		// Synchronize any wrongly unjoined workers before inspecting their storage.
		// The controlled-gate test independently rejects premature return.
		synctest.Wait()
		p17AssertRows(t, "mixed jobs", got, want)
		p17AssertCalls(t, "mixed jobs", calls)
	})
}

func TestConcurrentStartAndWaitAll(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		first, early := errors.New("job error"), errors.New("job error")
		want := []subject.Result{{Index: 0, Value: 31}, {Index: 1, Value: 0, Err: first}, {Index: 2, Value: -7}, {Index: 3, Value: 82, Err: early}}
		calls := make([]atomic.Int32, len(want))
		gates := make([]chan struct{}, len(want))
		released := make([]bool, len(want))
		jobs := make([]func() (int, error), len(want))
		for i, row := range want {
			gates[i] = make(chan struct{})
			gate := gates[i]
			jobs[i] = func() (int, error) { calls[i].Add(1); <-gate; return row.Value, row.Err }
		}
		release := func(i int) {
			if !released[i] {
				close(gates[i])
				released[i] = true
			}
		}
		defer func() {
			for i := range gates {
				release(i)
			}
			synctest.Wait()
		}()
		done := make(chan []subject.Result, 1)
		go func() { done <- subject.CollectResults(jobs) }()
		synctest.Wait()
		p17AssertCalls(t, "before any job is released", calls)
		var got []subject.Result
		received, reported := false, false
		observe := func(stage string, allReleased bool) {
			if !received {
				select {
				case got = <-done:
					received = true
				default:
				}
			}
			if received && !allReleased && !reported {
				t.Errorf("CollectResults returned %s while jobs remain blocked; want every job to return first", stage)
				reported = true
			}
			if allReleased && !received {
				t.Error("CollectResults has not returned after every job finished; want all indexed records")
			}
		}
		observe("before release", false)
		// Job3 reports an error first. Other jobs remain held after that error.
		for i := len(gates) - 1; i >= 0; i-- {
			release(i)
			synctest.Wait()
			observe("after releasing a job", i == 0)
		}
		if received {
			p17AssertRows(t, "reverse completion", got, want)
		}
		p17AssertCalls(t, "after every job is released", calls)
	})
}

func TestEmptyBatches(t *testing.T) {
	for _, c := range []struct {
		name string
		jobs []func() (int, error)
	}{{"nil", nil}, {"empty", []func() (int, error){}}} {
		t.Run(c.name, func(t *testing.T) {
			if got := subject.CollectResults(c.jobs); len(got) != 0 {
				t.Errorf("empty batch result length = %d, want 0", len(got))
			}
		})
	}
}

func TestIndependentResultRecords(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		marker := errors.New("first batch error")
		firstWant := []subject.Result{{Index: 0, Value: 5, Err: marker}, {Index: 1, Value: 14}}
		secondWant := []subject.Result{{Index: 0, Value: -3}, {Index: 1, Value: 0}}
		firstJobs, firstCalls := p17Jobs(firstWant)
		first := subject.CollectResults(firstJobs)
		synctest.Wait()
		secondJobs, secondCalls := p17Jobs(secondWant)
		second := subject.CollectResults(secondJobs)
		synctest.Wait()
		p17AssertRows(t, "first batch after second call", first, firstWant)
		p17AssertRows(t, "second batch", second, secondWant)
		p17AssertCalls(t, "first batch", firstCalls)
		p17AssertCalls(t, "second batch", secondCalls)
		if len(first) > 0 {
			first[0].Value = 999
		}
		p17AssertRows(t, "second batch after first result edit", second, secondWant)
	})
}

func TestConcurrentIndependentBatches(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		firstWant := []subject.Result{{Index: 0, Value: -42}, {Index: 1, Value: 17, Err: errors.New("first")}, {Index: 2, Value: 0}}
		secondWant := []subject.Result{{Index: 0, Value: 8, Err: errors.New("second")}, {Index: 1, Value: -6}}
		firstJobs, firstCalls := p17Jobs(firstWant)
		secondJobs, secondCalls := p17Jobs(secondWant)
		firstDone, secondDone := make(chan []subject.Result, 1), make(chan []subject.Result, 1)
		go func() { firstDone <- subject.CollectResults(firstJobs) }()
		go func() { secondDone <- subject.CollectResults(secondJobs) }()
		synctest.Wait()
		p17AssertRows(t, "concurrent first", <-firstDone, firstWant)
		p17AssertRows(t, "concurrent second", <-secondDone, secondWant)
		p17AssertCalls(t, "concurrent first", firstCalls)
		p17AssertCalls(t, "concurrent second", secondCalls)
	})
}
