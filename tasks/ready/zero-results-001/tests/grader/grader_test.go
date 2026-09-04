package subject_test

import (
	"task.local/zero-results-001/subject"
	"testing"
)

func TestExactDefaults(t *testing.T) {
	for attempt := 0; attempt < 3; attempt++ {
		n, b, s, v, m := subject.Defaults()
		if n != 0 || b || s != "" || v != nil || m != nil {
			t.Fatalf("Defaults must return the exact zero value of every result type")
		}
	}
}
