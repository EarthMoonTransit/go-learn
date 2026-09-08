package subject_test

import (
	"task.local/category-closure-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	f := subject.InCategory("a")
	if f == nil {
		t.Fatal("nil predicate")
	}
	if !f(subject.Record{Category: "a"}) || f(subject.Record{Category: "b"}) {
		t.Error("predicate did not match exact category")
	}
}

func TestContractCase02(t *testing.T) {
	f := subject.InCategory("a")
	if f == nil {
		t.Fatal("nil predicate")
	}
	for _, c := range []string{"A", "a ", " a"} {
		if f(subject.Record{Category: c}) {
			t.Errorf("unexpected category %q", c)
		}
	}
}

func TestContractCase03(t *testing.T) {
	f := subject.InCategory("")
	if f == nil || !f(subject.Record{}) {
		t.Error("empty category rejected")
	}
}

func TestContractCase04(t *testing.T) {
	a := subject.InCategory("a")
	b := subject.InCategory("b")
	if a == nil || b == nil {
		t.Fatal("nil predicates")
	}
	for i := 0; i < 2; i++ {
		if !a(subject.Record{Category: "a"}) || a(subject.Record{Category: "b"}) || !b(subject.Record{Category: "b"}) || b(subject.Record{Category: "a"}) {
			t.Error("predicates share captured category")
		}
	}
}

func TestContractCase05(t *testing.T) {
	f := subject.InCategory("a")
	if f == nil || !f(subject.Record{Day: -3, Amount: -2, Category: "a"}) {
		t.Error("unrelated fields affected match")
	}
}
