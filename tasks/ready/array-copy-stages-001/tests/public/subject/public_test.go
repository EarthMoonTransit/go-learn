package subject_test

import (
	subject "task.local/array-copy-stages-001/subject"
	"testing"
)

func publicArrayStages(t *testing.T, input [5]int) {
	t.Helper()
	source, assigned, inside, after := subject.ArrayStages(input[0], input[1], input[2], input[3], input[4])
	original := [3]int{input[0], input[1], input[2]}
	changed := [3]int{input[3], input[1], input[2]}
	local := [3]int{input[3], input[4], input[2]}
	if source != original {
		t.Errorf("source=%v; want %v after modifying only the assigned array", source, original)
	}
	if assigned != changed {
		t.Errorf("assigned before call=%v; want %v with only index 0 replaced", assigned, changed)
	}
	if inside != local {
		t.Errorf("insideCall=%v; want %v with callee index 1 replaced", inside, local)
	}
	if after != changed {
		t.Errorf("afterCall=%v; want %v: changing callee index 1 must not change its caller", after, changed)
	}
}

func TestPublicDistinctAssignmentAndCallStages(t *testing.T) {
	publicArrayStages(t, [5]int{-4, 7, 2, 9, -8})
}

func TestPublicZeroReplacementAndLastElement(t *testing.T) {
	publicArrayStages(t, [5]int{5, -3, -9, 0, 6})
}

func TestPublicZeroCalleeValue(t *testing.T) {
	publicArrayStages(t, [5]int{0, 0, 4, -2, 0})
}
