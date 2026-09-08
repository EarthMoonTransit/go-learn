package subject_test

import (
	"errors"
	"task.local/error-context-wrap-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	cause := errors.New("")
	w := subject.Wrap("x", cause)
	if w == nil || w.Error() != "x: " || errors.Unwrap(w) != cause {
		t.Errorf("empty error got%v", w)
	}
}
