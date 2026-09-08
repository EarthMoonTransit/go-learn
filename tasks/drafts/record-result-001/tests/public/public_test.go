package subject_test

import (
	"task.local/record-result-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	r, e := subject.NewRecord("data", 7)
	if e != nil || r == nil || *r != (subject.Record{Label: "data", Count: 7}) {
		t.Fatalf("record got%v,%v", r, e)
	}
}

func TestPublicCase02(t *testing.T) {
	a, e := subject.NewRecord("", -1)
	b, f := subject.NewRecord("", -1)
	if e != nil || f != nil || a == nil || b == nil || *a != (subject.Record{Count: -1}) || *b != (subject.Record{Count: -1}) {
		t.Fatalf("empty negative got%v,%v %v,%v", a, e, b, f)
	}
	a.Count = 99
	a.Label = "x"
	if *b != (subject.Record{Count: -1}) {
		t.Errorf("results share storage %v", b)
	}
	c, e := subject.NewRecord("", -1)
	if e != nil || c == nil || *c != (subject.Record{Count: -1}) || a.Count != 99 {
		t.Errorf("later changed state a%v c%v err%v", a, c, e)
	}
}
