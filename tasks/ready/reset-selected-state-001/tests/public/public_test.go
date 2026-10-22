package subject_test

import (
	"reflect"
	subject "task.local/reset-selected-state-001/subject"
	"testing"
)

func TestSelectedZeroValuesAndSharing(t *testing.T) {
	backing := []int{7, 11, 19}
	labels := map[string]int{"keep": 23}
	target := 29
	shared := subject.State{Count: 5, Active: true, Name: " value ", Values: backing[:2], Labels: labels, Target: &target}
	states := []subject.State{shared, shared, {}, shared}
	selected := []bool{true, false, true, true}
	saved := states[1]
	got := subject.ResetSelected(states, selected)
	if got != 3 {
		t.Fatalf("selected count=%d; want 3 including the already-zero state", got)
	}
	for _, i := range []int{0, 2, 3} {
		if !reflect.DeepEqual(states[i], subject.State{}) {
			t.Fatalf("state[%d]=%+v (nil slice/map/pointer=%t/%t/%t); want exact State zero value", i, states[i], states[i].Values == nil, states[i].Labels == nil, states[i].Target == nil)
		}
	}
	if !reflect.DeepEqual(selected, []bool{true, false, true, true}) {
		t.Fatalf("selection changed: %v; want [true false true true]", selected)
	}
	if !reflect.DeepEqual(backing, []int{7, 11, 19}) || !reflect.DeepEqual(labels, map[string]int{"keep": 23}) || target != 29 {
		t.Fatalf("old shared data changed: slice=%v map=%v pointee=%d; want [7 11 19], map[keep:23], 29", backing, labels, target)
	}
	u := states[1]
	if !reflect.DeepEqual(u, saved) || len(u.Values) != 2 || cap(u.Values) != 3 || &u.Values[0] != &backing[0] || u.Target != &target {
		t.Fatalf("unselected state/header/pointer changed: %+v len/cap=%d/%d; want original state,2/3 and original identities", u, len(u.Values), cap(u.Values))
	}
	u.Labels["new"] = 31
	if labels["new"] != 31 {
		t.Fatalf("unselected map detached: old map[new]=%d; want 31", labels["new"])
	}
	u.Values[0] = 37
	if backing[0] != 37 {
		t.Fatalf("unselected slice detached: old[0]=%d; want 37", backing[0])
	}
}
func TestResetCountsAndRepeatedCalls(t *testing.T) {
	if got := subject.ResetSelected(nil, nil); got != 0 {
		t.Fatalf("empty reset=%d; want 0", got)
	}
	states := []subject.State{{Count: 8}, {Values: []int{}, Labels: map[string]int{}}}
	before := append([]subject.State(nil), states...)
	if got := subject.ResetSelected(states, []bool{false, false}); got != 0 || !reflect.DeepEqual(states, before) {
		t.Fatalf("no selection: count=%d states=%+v; want 0 and unchanged %+v", got, states, before)
	}
	for pass := 0; pass < 2; pass++ {
		if got := subject.ResetSelected(states, []bool{true, true}); got != 2 || !reflect.DeepEqual(states, []subject.State{{}, {}}) {
			t.Fatalf("repeat %d: count=%d states=%+v; want 2 and two zero states", pass, got, states)
		}
	}
	states[1] = subject.State{Count: -4, Active: true, Name: "again"}
	if got := subject.ResetSelected(states, []bool{false, true}); got != 1 || !reflect.DeepEqual(states[1], subject.State{}) {
		t.Fatalf("repopulated state: count=%d state=%+v; want 1,zero", got, states[1])
	}
}
