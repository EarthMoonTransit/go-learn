package subject_test

import (
	"reflect"
	subject "task.local/guarded-stage-values-001/subject"
	"testing"
)

func TestStageOrderCommitAndStop(t *testing.T) {
	trace := []int{}
	a := func(v int) (int, bool) { trace = append(trace, 1000+v); return v + 3, true }
	b := func(v int) (int, bool) { trace = append(trace, 2000+v); return v * 4, true }
	reject := func(v int) (int, bool) { trace = append(trace, 3000+v); return -999, false }
	late := func(v int) (int, bool) { trace = append(trace, 4000+v); return v + 100, true }
	steps := []func(int) (int, bool){nil, a, nil, b, reject, late}
	for _, initial := range []int{2, -1} {
		trace = nil
		got, count := subject.RunStages(initial, steps)
		want := (initial + 3) * 4
		wantTrace := []int{1000 + initial, 2000 + initial + 3, 3000 + want}
		if got != want || count != 2 || !reflect.DeepEqual(trace, wantTrace) {
			t.Fatalf("initial=%d: got value/count=%d/%d calls=%v; want %d/2 calls=%v", initial, got, count, trace, want, wantTrace)
		}
		if steps[0] != nil || steps[2] != nil {
			t.Fatal("nil entries changed")
		}
		for _, i := range []int{1, 3, 4, 5} {
			if steps[i] == nil {
				t.Fatalf("step %d cleared", i)
			}
		}
	}
}
func TestStageNilRepeatedAndFirstRejection(t *testing.T) {
	for _, steps := range [][]func(int) (int, bool){nil, {}, {nil, nil}} {
		value, count := subject.RunStages(37, steps)
		if value != 37 || count != 0 {
			t.Fatalf("empty/nil stages: value/count=%d/%d; want 37/0", value, count)
		}
	}
	calls := 0
	next := func(v int) (int, bool) { calls++; return v + calls, true }
	value, count := subject.RunStages(0, []func(int) (int, bool){next, next, next})
	if value != 6 || count != 3 || calls != 3 {
		t.Fatalf("repeated function: value/count/calls=%d/%d/%d; want 6/3/3", value, count, calls)
	}
	calls = 0
	reject := func(v int) (int, bool) { calls++; return 999, false }
	value, count = subject.RunStages(-7, []func(int) (int, bool){reject, next})
	if value != -7 || count != 0 || calls != 1 {
		t.Fatalf("first rejection: value/count/calls=%d/%d/%d; want -7/0/1", value, count, calls)
	}
}
func TestStageEntriesRetained(t *testing.T) {
	steps := []func(int) (int, bool){func(v int) (int, bool) { return v + 2, true }, func(v int) (int, bool) { return v - 3, true }}
	subject.RunStages(10, steps)
	for i, delta := range []int{2, -3} {
		if steps[i] == nil {
			t.Fatalf("entry %d became nil", i)
		}
		got, ok := steps[i](20)
		if !ok || got != 20+delta {
			t.Fatalf("entry %d behavior changed: (%d,%t); want (%d,true)", i, got, ok, 20+delta)
		}
	}
}
