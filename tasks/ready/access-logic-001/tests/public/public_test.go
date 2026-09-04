package subject_test

import (
	"task.local/access-logic-001/subject"
	"testing"
)

func TestGateExamples(t *testing.T) {
	for _, c := range []struct{ e, b, o, w bool }{{true, false, false, true}, {true, true, false, false}, {false, true, true, true}} {
		if got := subject.Allowed(c.e, c.b, c.o); got != c.w {
			t.Errorf("Allowed(%v,%v,%v)=%v, want %v", c.e, c.b, c.o, got, c.w)
		}
	}
}
