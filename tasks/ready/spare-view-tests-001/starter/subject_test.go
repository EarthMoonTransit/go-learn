package subject_test

import (
	subject "task.local/spare-view-tests-001/subject"
	"testing"
)

func TestEmptyView(t *testing.T) {
	got := subject.SpareView(nil, 0, 0, 0)
	if len(got) != 0 || cap(got) != 0 {
		t.Fatalf("empty len/cap=%d/%d; want 0/0", len(got), cap(got))
	}
}
