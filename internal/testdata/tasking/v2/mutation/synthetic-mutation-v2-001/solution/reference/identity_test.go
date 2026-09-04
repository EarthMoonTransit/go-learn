package subject_test

import (
	"task.local/synthetic-mutation-v2-001/subject"
	"testing"
)

func TestIdentityNonzero(t *testing.T) {
	if got := subject.Identity(7); got != 7 {
		t.Errorf("got %d", got)
	}
}
