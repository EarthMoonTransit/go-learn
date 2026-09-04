package subject_test

import (
	"testing"

	"task.local/map-take-001/subject"
)

func TestTakeRecognizesStoredZero(t *testing.T) {
	values := map[string]int{"zero": 0}
	value, ok := subject.Take(values, "zero")
	if value != 0 || !ok {
		t.Fatalf("Take returned %d, %t; want 0, true", value, ok)
	}
	if _, remains := values["zero"]; remains {
		t.Fatal("entry was not deleted")
	}
}
