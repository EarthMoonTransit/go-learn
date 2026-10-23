package subject_test

import (
	"errors"
	"fmt"
	"task.local/classified-stage-errors-001/subject"
	"testing"
)

func TestFiniteMixedStageTrees(t *testing.T) {
	for _, expose := range []bool{false, true} {
		for fail := 0; fail < 4; fail++ {
			detail := &subject.Detail{Code: fail + 11, Message: "boundary"}
			cause := errors.Join(fmt.Errorf("a: %w", errors.New("miss")), fmt.Errorf("b: %w", detail), subject.ErrTemporary)
			called := 0
			stages := make([]subject.Stage, 4)
			for i := range stages {
				stages[i] = subject.Stage{Name: fmt.Sprintf(" operation %d ", i), Expose: expose, Run: func() (int, error) {
					called++
					if i == fail {
						return i - 5, cause
					}
					return i + 2, nil
				}}
			}
			out := subject.RunStages(stages)
			if called != fail+1 {
				t.Errorf("finite call count=%d; want %d", called, fail+1)
			}
			wantDetail := (*subject.Detail)(nil)
			if expose {
				wantDetail = detail
			}
			checkFailure(t, out, fail-5, fmt.Sprintf(" operation %d ", fail), cause, expose, expose, wantDetail, expose)
		}
	}
}
