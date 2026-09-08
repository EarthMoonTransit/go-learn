package subject_test

import (
	"slices"
	"task.local/predicate-filter-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	in := []subject.Record{{Day: 1, Amount: 5.15, Category: "a"}, {Day: 13, Amount: 55.67, Category: "b"}, {Day: 15, Amount: 11, Category: "a"}}
	got := subject.Filter(in, func(r subject.Record) bool { return r.Category == "a" })
	want := []subject.Record{in[0], in[2]}
	if !slices.Equal(got, want) {
		t.Errorf("filtered got %+v want %+v", got, want)
	}
}

func TestPublicCase02(t *testing.T) {
	in := []subject.Record{{Day: 3}, {Day: 1}, {Day: 3}}
	var called []int
	got := subject.Filter(in, func(r subject.Record) bool { called = append(called, r.Day); return true })
	if !slices.Equal(called, []int{3, 1, 3}) || !slices.Equal(got, in) {
		t.Errorf("order/calls got %+v,%v want %+v,[3 1 3]", got, called, in)
	}
}

func TestPublicCase03(t *testing.T) {
	in := []subject.Record{{Day: 1}, {Day: 2}}
	before := append([]subject.Record(nil), in...)
	subject.Filter(in, func(r subject.Record) bool { return r.Day == 2 })
	if !slices.Equal(in, before) {
		t.Errorf("input changed: %+v", in)
	}
}

func TestPublicCase04(t *testing.T) {
	calls := 0
	got := subject.Filter(nil, func(r subject.Record) bool { calls++; return true })
	if len(got) != 0 || calls != 0 {
		t.Errorf("empty got %+v calls%d want empty/0", got, calls)
	}
}

func TestPublicCase05(t *testing.T) {
	got := subject.Filter([]subject.Record{{Day: 1}}, func(r subject.Record) bool { return false })
	if len(got) != 0 {
		t.Errorf("reject all got %+v", got)
	}
}
