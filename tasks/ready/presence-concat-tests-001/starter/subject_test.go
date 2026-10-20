package subject_test

import (
	subject "task.local/presence-concat-tests-001/subject"
	"testing"
)

func TestNoParts(t *testing.T) {
	if subject.ConcatPresent() != nil {
		t.Fatal("no parts: want nil")
	}
}
