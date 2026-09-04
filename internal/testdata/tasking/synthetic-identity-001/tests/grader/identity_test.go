package subject_test

import (
	"testing"

	"task.local/synthetic-identity-001/subject"
)

func TestIdentityGrader(t *testing.T) {
	for _, value := range []int{0, -9, 42} {
		if got := subject.Identity(value); got != value {
			t.Errorf("Identity(%d) = %d, want %d", value, got, value)
		}
	}
}
