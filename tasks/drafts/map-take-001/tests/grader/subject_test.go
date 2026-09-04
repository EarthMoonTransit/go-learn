package subject_test

import (
	"testing"

	"task.local/map-take-001/subject"
)

func TestTakeExistingAndPreservesOthers(t *testing.T) {
	values := map[string]int{"target": 7, "other": 9}
	value, ok := subject.Take(values, "target")
	if value != 7 || !ok {
		t.Fatalf("got %d, %t", value, ok)
	}
	if _, exists := values["target"]; exists {
		t.Fatal("target remains")
	}
	if values["other"] != 9 || len(values) != 1 {
		t.Fatalf("unrelated entries changed: %v", values)
	}
}

func TestTakeMissingAndNil(t *testing.T) {
	values := map[string]int{"other": 2}
	if value, ok := subject.Take(values, "missing"); value != 0 || ok {
		t.Fatalf("missing returned %d, %t", value, ok)
	}
	if len(values) != 1 {
		t.Fatalf("missing lookup changed map: %v", values)
	}
	if value, ok := subject.Take(nil, "missing"); value != 0 || ok {
		t.Fatalf("nil returned %d, %t", value, ok)
	}
}
