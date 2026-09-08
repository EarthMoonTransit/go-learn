package subject_test

import (
	"task.local/period-predicate-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	f := subject.WithinDays(subject.DaysPeriod{From: 1, To: 15})
	if f == nil {
		t.Fatal("got nil predicate")
	}
	for _, x := range []struct {
		day  int
		want bool
	}{{0, false}, {1, true}, {13, true}, {15, true}, {16, false}} {
		if got := f(subject.Record{Day: x.day, Amount: 3, Category: "a"}); got != x.want {
			t.Errorf("day%d got %v want %v", x.day, got, x.want)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	f := subject.WithinDays(subject.DaysPeriod{From: 2, To: 2})
	if f == nil || !f(subject.Record{Day: 2}) {
		t.Error("single-day period rejected")
	}
}

func TestPublicCase03(t *testing.T) {
	f := subject.WithinDays(subject.DaysPeriod{From: 3, To: 1})
	if f == nil {
		t.Fatal("nil predicate")
	}
	for _, day := range []int{0, 1, 2, 3, 4} {
		if f(subject.Record{Day: day}) {
			t.Errorf("reversed range accepted day%d", day)
		}
	}
}

func TestPublicCase04(t *testing.T) {
	p := subject.DaysPeriod{From: 1, To: 2}
	a := subject.WithinDays(p)
	p.From = 9
	p.To = 10
	b := subject.WithinDays(p)
	if a == nil || b == nil {
		t.Fatal("nil predicates")
	}
	if !a(subject.Record{Day: 1}) || a(subject.Record{Day: 9}) || !b(subject.Record{Day: 9}) || b(subject.Record{Day: 1}) {
		t.Error("predicates lost their independent periods")
	}
}

func TestPublicCase05(t *testing.T) {
	f := subject.WithinDays(subject.DaysPeriod{From: 1, To: 2})
	if f == nil || !f(subject.Record{Day: 1, Amount: -99, Category: ""}) {
		t.Error("non-Day fields affected acceptance")
	}
}
