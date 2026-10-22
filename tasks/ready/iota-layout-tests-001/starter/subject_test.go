package subject_test

import (
	subject "task.local/iota-layout-tests-001/subject"
	"testing"
)

func TestFirstCode(t *testing.T) {
	codes, _, _ := subject.Layout()
	if codes[0] != 10 {
		t.Fatalf("first code=%d; want 10", codes[0])
	}
}
