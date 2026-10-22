package subject_test

import (
	subject "task.local/array-copy-stages-001/subject"
	"testing"
)

func TestPredictAssignmentLeavesOriginal(t *testing.T) {
	source, assigned, _, _ := subject.ArrayStages(-4, 7, 2, 9, -8)
	if source != [3]int{-4, 7, 2} {
		t.Errorf("original after assignment mutation=%v; want [-4 7 2]", source)
	}
	if assigned != [3]int{9, 7, 2} {
		t.Errorf("assigned before call=%v; want [9 7 2]", assigned)
	}
}

func TestPredictCallLeavesCaller(t *testing.T) {
	_, _, inside, after := subject.ArrayStages(5, -3, -9, 0, 6)
	if inside != [3]int{0, 6, -9} {
		t.Errorf("callee local value=%v; want [0 6 -9]", inside)
	}
	if after != [3]int{0, -3, -9} {
		t.Errorf("caller after value call=%v; want [0 -3 -9]", after)
	}
}
