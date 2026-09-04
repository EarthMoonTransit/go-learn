package subject_test

import (
	"task.local/synthetic-mutation-v2-001/subject"
	"testing"
)

func TestIdentityZero(t *testing.T) {
	if got := subject.Identity(0); got != 0 {
		t.Errorf("got %d", got)
	}
}
