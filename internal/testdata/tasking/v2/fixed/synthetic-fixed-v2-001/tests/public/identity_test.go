package subject_test

import (
	"task.local/synthetic-fixed-v2-001/subject"
	"testing"
)

func TestIdentityPublic(t *testing.T) {
	if got := subject.Identity(7); got != 7 {
		t.Fatalf("got %d", got)
	}
}
