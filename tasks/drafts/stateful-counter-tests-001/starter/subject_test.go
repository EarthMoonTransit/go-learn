package subject_test

import (
	"task.local/stateful-counter-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	f := subject.NewCounter(10)
	if f == nil {
		t.Fatal("nil closure")
	}
	if got := f(2); got != 12 {
		t.Errorf("first got%d want12", got)
	}
}

// Add cases for the remaining requirements in TASK.md.
