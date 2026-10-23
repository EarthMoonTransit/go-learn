package subject_test

import (
	"reflect"
	"slices"
	"task.local/selected-window-copies-001/subject"
	"testing"
)

func windowsCheck(t *testing.T, values []int, windows []subject.Window) {
	t.Helper()
	input := make([]int, len(values), len(values)+17)
	copy(input, values)
	before := append([]int(nil), input...)
	var want [][]int
	for _, w := range windows {
		if w.Start < w.End {
			want = append(want, append([]int(nil), values[w.Start:w.End]...))
		}
	}
	got := subject.SelectedCopies(input, windows)
	if !slices.Equal(input, before) {
		t.Errorf("input changed during selection: got %v; want %v", input, before)
	}
	if len(got) != len(want) {
		t.Errorf("selected result count=%d; want %d without empty/prefix entries", len(got), len(want))
		return
	}
	if cap(got) != len(want) {
		t.Errorf("outer capacity=%d; want selected count %d", cap(got), len(want))
	}
	for i := range want {
		if len(got[i]) != len(want[i]) {
			t.Errorf("window %d length=%d; want %d", i, len(got[i]), len(want[i]))
			return
		}
		if cap(got[i]) != len(want[i]) {
			t.Errorf("window %d capacity=%d; want %d", i, cap(got[i]), len(want[i]))
		}
		if !slices.Equal(got[i], want[i]) {
			t.Errorf("window %d contents=%v; want %v in requested order", i, got[i], want[i])
		}
	}
	for i := range input {
		input[i] = 10000 + i
	}
	for i := range want {
		if !slices.Equal(got[i], want[i]) {
			t.Errorf("input mutation changed window %d: got %v; want %v", i, got[i], want[i])
		}
	}
	copy(input, before)
	for i := range got {
		for at := range got[i] {
			old := got[i][at]
			got[i][at] = 20000 + i + at
			if !slices.Equal(input, before) {
				t.Errorf("window mutation changed input: window %d index %d", i, at)
			}
			for peer := range got {
				if peer != i && !slices.Equal(got[peer], want[peer]) {
					t.Errorf("window %d mutation changed peer %d: got %v; want %v", i, peer, got[peer], want[peer])
				}
			}
			got[i][at] = old
		}
	}
}
func TestSelectNonemptyInRequestOrder(t *testing.T) {
	windowsCheck(t, []int{8, -2, 0, 5, 11, 19}, []subject.Window{{1, 4}, {2, 5}, {0, 0}, {1, 4}, {5, 6}, {3, 3}})
	windowsCheck(t, []int{4, 7, -3}, []subject.Window{{2, 3}, {0, 2}})
}
func TestEmptySelections(t *testing.T) {
	windowsCheck(t, nil, nil)
	windowsCheck(t, []int{}, []subject.Window{{0, 0}})
	windowsCheck(t, []int{1, 2, 3}, nil)
	windowsCheck(t, []int{1, 2, 3}, []subject.Window{{0, 0}, {2, 2}, {3, 3}})
}
func TestIndependentRepeatedCalls(t *testing.T) {
	input := []int{7, -1, 3, 9}
	windows := []subject.Window{{1, 3}, {2, 4}}
	first := subject.SelectedCopies(input, windows)
	second := subject.SelectedCopies(input, windows)
	if len(first) != 2 || len(second) != 2 || len(first[0]) != 2 || len(second[0]) != 2 {
		t.Error("repeated call shape mismatch")
		return
	}
	saved := [][]int{{-1, 3}, {3, 9}}
	first[0][0] = 88
	if !reflect.DeepEqual(second, saved) {
		t.Errorf("one call changed another: got %v; want %v", second, saved)
	}
	if !slices.Equal(input, []int{7, -1, 3, 9}) {
		t.Error("repeated call result aliases input")
	}
}
