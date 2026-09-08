package subject_test

import (
	"task.local/decision-table-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if got := subject.SelectAction("ace", "ace", "jack"); got != "P" {
		t.Errorf("pair got%q wantP", got)
	}
}

func TestContractCase02(t *testing.T) {
	for _, c := range []struct{ a, b, v, want string }{{"king", "ace", "queen", "S"}, {"ace", "ten", "five", "W"}, {"ace", "king", "nine", "W"}, {"ten", "king", "ace", "S"}, {"seven", "king", "ace", "S"}, {"ten", "six", "six", "S"}, {"ten", "six", "seven", "H"}, {"ten", "two", "six", "S"}, {"ten", "two", "seven", "H"}, {"nine", "two", "two", "H"}, {"two", "two", "five", "H"}} {
		if got := subject.SelectAction(c.a, c.b, c.v); got != c.want {
			t.Errorf("decision %s/%s/%s got%q want%q", c.a, c.b, c.v, got, c.want)
		}
	}
}

func TestContractCase03(t *testing.T) {
	for _, c := range []struct{ a, b, v, want string }{{"Ace", "king", "two", "H"}, {"ace ", "ten", "two", "H"}, {"", "ace", "", "H"}, {"ace", "ten", "king ", "W"}, {"ten", "two", "unknown", "S"}, {"ten", "two", "Seven", "S"}} {
		if got := subject.SelectAction(c.a, c.b, c.v); got != c.want {
			t.Errorf("exact token %q/%q/%q got%q want%q", c.a, c.b, c.v, got, c.want)
		}
	}
}
