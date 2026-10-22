package subject_test

import (
	"reflect"
	subject "task.local/reset-selected-state-001/subject"
	"testing"
)

func TestEverySelectionMask(t *testing.T) {
	for mask := 0; mask < 16; mask++ {
		states := make([]subject.State, 4)
		selected := make([]bool, 4)
		wantCount := 0
		for i := range states {
			states[i] = subject.State{Count: i + 1, Active: i%2 == 0, Name: "x"}
			selected[i] = mask&(1<<i) != 0
			if selected[i] {
				wantCount++
			}
		}
		before := append([]subject.State(nil), states...)
		got := subject.ResetSelected(states, selected)
		if got != wantCount {
			t.Fatalf("mask=%d count=%d; want %d", mask, got, wantCount)
		}
		for i := range states {
			want := before[i]
			if selected[i] {
				want = subject.State{}
			}
			if !reflect.DeepEqual(states[i], want) {
				t.Fatalf("mask=%d state[%d]=%+v; want %+v", mask, i, states[i], want)
			}
		}
	}
}
