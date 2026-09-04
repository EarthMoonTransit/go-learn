package subject_test

import (
	"testing"

	"task.local/synthetic/subject"
)

func TestBypass(t *testing.T) {
	t.Fatal("always")
}

var _ = subject.Identity
