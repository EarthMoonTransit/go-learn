package subject_test

import (
	"task.local/array-copy-address-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	source := &[5]string{"a", "b", "c", "d", "e"}
	out, ptrs := subject.CopyArray(source)
	if out == nil || *out != *source {
		t.Fatalf("copy got%v want%v", out, source)
	}
	for i, p := range ptrs {
		if p != &out[i] {
			t.Errorf("address%d not copied element", i)
		}
	}
}

// Add cases for the remaining requirements in TASK.md.
