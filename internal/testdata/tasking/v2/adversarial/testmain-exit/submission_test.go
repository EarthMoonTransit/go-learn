package subject_test

import (
	"os"
	"testing"

	"task.local/synthetic/subject"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestIdentity(t *testing.T) {
	if subject.Identity(1) != 1 {
		t.Error("wrong result")
	}
}
