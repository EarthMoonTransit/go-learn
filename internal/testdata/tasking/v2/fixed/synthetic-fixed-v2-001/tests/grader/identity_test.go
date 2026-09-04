package subject_test

import (
	"task.local/synthetic-fixed-v2-001/subject"
	"testing"
)

func TestIdentityGrader(t *testing.T) {
	for _, value := range []int{0, -9, 42} {
		if got := subject.Identity(value); got != value {
			t.Errorf("got %d", got)
		}
	}
}
