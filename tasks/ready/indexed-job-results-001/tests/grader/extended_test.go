package subject_test

import (
	"errors"
	subject "task.local/indexed-job-results-001/subject"
	"testing"
	"testing/synctest"
)

type p17TypedError struct{ Label string }

func (e *p17TypedError) Error() string {
	if e == nil {
		return "typed nil error"
	}
	return e.Label
}

func TestTypedNilErrorAndSingleton(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var typedNil *p17TypedError
		want := []subject.Result{{Index: 0, Value: -23, Err: typedNil}}
		jobs, calls := p17Jobs(want)
		got := subject.CollectResults(jobs)
		synctest.Wait()
		p17AssertRows(t, "typed nil error", got, want)
		p17AssertCalls(t, "singleton", calls)
	})
}

func TestLongBatchAndRepeatedValues(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		want := make([]subject.Result, 19)
		for i := range want {
			want[i] = subject.Result{Index: i, Value: i%5 - 2}
			if i%3 == 0 {
				want[i].Err = errors.New("repeated text")
			}
		}
		jobs, calls := p17Jobs(want)
		got := subject.CollectResults(jobs)
		synctest.Wait()
		p17AssertRows(t, "long repeated batch", got, want)
		p17AssertCalls(t, "long repeated batch", calls)
	})
}
