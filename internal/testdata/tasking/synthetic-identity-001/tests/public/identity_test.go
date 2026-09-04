package subject_test

import (
	"testing"

	"task.local/synthetic-identity-001/subject"
)

func TestIdentityPublic(t *testing.T) {
	if got := subject.Identity(7); got != 7 {
		t.Errorf("Identity(7) = %d, want 7", got)
	}
}
