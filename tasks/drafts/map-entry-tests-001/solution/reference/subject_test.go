package subject_test

import (
	"maps"
	"task.local/map-entry-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	got := subject.Entries(map[string]int{"a": 1})
	if len(got) != 1 || got[0] != (subject.Entry{Key: "a", Value: 1}) {
		t.Errorf("single got%v", got)
	}
}

func TestContractCase02(t *testing.T) {
	in := map[string]int{"a": 1, "": 0, "λ": -2, " x ": 7, "last": 9}
	before := maps.Clone(in)
	got := subject.Entries(in)
	if len(got) != len(in) {
		t.Fatalf("entries length got%d want%d", len(got), len(in))
	}
	seen := map[string]bool{}
	for _, e := range got {
		v, ok := in[e.Key]
		if !ok || seen[e.Key] || v != e.Value {
			t.Errorf("missing/duplicate/wrong entry%v", e)
		}
		seen[e.Key] = true
	}
	if !maps.Equal(in, before) {
		t.Errorf("input changed %v", in)
	}
	got[0].Value = 99
	if !maps.Equal(in, before) {
		t.Error("result shares mutable state")
	}
}

func TestContractCase03(t *testing.T) {
	for _, in := range []map[string]int{nil, {}} {
		if got := subject.Entries(in); len(got) != 0 {
			t.Errorf("empty got%v", got)
		}
	}
}
