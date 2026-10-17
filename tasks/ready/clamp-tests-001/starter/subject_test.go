package subject_test

import (
	"task.local/clamp-tests-001/subject"
	"testing"
)

// Go discovers Test functions in _test.go files. Keep the supplied signature;
// t is the runner's handle for reporting a failure.
func TestClampInsideRange(t *testing.T) {
	if got := subject.Clamp(5, 0, 10); got != 5 {
		t.Fatalf("Clamp(5,0,10)=%d, want 5", got)
	}
}

func TestClampBelowRange(t *testing.T) {
	if got := subject.Clamp(-3, -1, 8); got != -1 {
		// %d shows the actual result; the arguments and want value identify
		// the case. This branch runs only when the comparison fails.
		t.Fatalf("Clamp(-3,-1,8)=%d, want -1", got)
	}
}

// Add cases for the remaining requirements. No table or helper is required.
