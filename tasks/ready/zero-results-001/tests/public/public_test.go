package subject_test

import (
	"task.local/zero-results-001/subject"
	"testing"
)

func TestScalarDefaults(t *testing.T) {
	n, b, s, _, _ := subject.Defaults()
	if n != 0 || b || s != "" {
		t.Fatalf("scalar defaults = (%d, %v, %q), want (0, false, empty string)", n, b, s)
	}
}
