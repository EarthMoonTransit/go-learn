package subject_test

import (
	subject "task.local/array-copy-stages-001/subject"
	"testing"
)

func checkArrayTuple(t *testing.T, input [5]int) {
	t.Helper()
	source, assigned, inside, after := subject.ArrayStages(input[0], input[1], input[2], input[3], input[4])
	if source != [3]int{input[0], input[1], input[2]} ||
		assigned != [3]int{input[3], input[1], input[2]} ||
		inside != [3]int{input[3], input[4], input[2]} || after != assigned {
		t.Fatalf("input=%v: source/assigned/callee/caller=%v/%v/%v/%v; want independent stage values", input, source, assigned, inside, after)
	}
}

func TestGraderFiniteArrayInputs(t *testing.T) {
	values := [3]int{-2, 0, 3}
	for code := 0; code < 243; code++ {
		var input [5]int
		remainder := code
		for i := range input {
			input[i] = values[remainder%3]
			remainder /= 3
		}
		checkArrayTuple(t, input)
	}
	for _, input := range [][5]int{{-9, 9, -9, 9, -9}, {9, -9, 9, -9, 9}} {
		checkArrayTuple(t, input)
	}
}

func TestGraderArrayCallsHaveNoRetainedState(t *testing.T) {
	first := [5]int{8, -7, 6, -5, 4}
	checkArrayTuple(t, first)
	checkArrayTuple(t, [5]int{-1, 2, -3, 4, -5})
	checkArrayTuple(t, first)
}
