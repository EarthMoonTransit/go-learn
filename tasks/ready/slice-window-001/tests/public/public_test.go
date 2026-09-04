package subject_test

import (
	"task.local/slice-window-001/subject"
	"testing"
)

func TestWindowExamples(t *testing.T) {
	input := make([]int, 3, 6)
	copy(input, []int{5, 6, 7})
	out, ok := subject.Window(input, 1, 3)
	if !ok || len(out) != 2 || cap(out) != 5 || out[0] != 6 {
		t.Fatalf("wrong ordinary window")
	}
	if _, ok := subject.Window(input, 1, 4); ok {
		t.Fatal("high beyond length must be rejected even with spare capacity")
	}
}

func TestInvalidWindowPreservesInput(t *testing.T) {
	input := []int{3, 4}
	out, ok := subject.Window(input, -1, 1)
	if ok || out != nil || input[0] != 3 || input[1] != 4 {
		t.Fatal("invalid bounds must return nil,false and preserve input")
	}
}
