package subject_test

import (
	"task.local/access-logic-001/subject"
	"testing"
)

func TestTruthTable(t *testing.T) {
	cases := []struct{ e, b, o, w bool }{
		{false, false, false, false}, {false, false, true, true}, {false, true, false, false}, {false, true, true, true},
		{true, false, false, true}, {true, false, true, true}, {true, true, false, false}, {true, true, true, true},
	}
	for _, c := range cases {
		if subject.Allowed(c.e, c.b, c.o) != c.w {
			t.Errorf("incorrect boolean combination")
		}
	}
}
