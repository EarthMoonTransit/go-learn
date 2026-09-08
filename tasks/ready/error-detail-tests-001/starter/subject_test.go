package subject_test

import (
	"task.local/error-detail-tests-001/subject"
	"testing"
)

type projectedError struct{ detail *subject.Detail }

func (e *projectedError) Error() string { return "projection" }
func (e *projectedError) As(target any) bool {
	p, ok := target.(**subject.Detail)
	if ok {
		*p = e.detail
	}
	return ok
}

func TestStartingCase01(t *testing.T) {
	d := &subject.Detail{Code: 7, Message: "bad"}
	got, ok := subject.FindDetail(d)
	if got != d || !ok {
		t.Errorf("direct got%v,%v", got, ok)
	}
}

// Add cases for the remaining requirements in TASK.md.
