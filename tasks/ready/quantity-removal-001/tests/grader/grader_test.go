package subject_test

import (
	"task.local/quantity-removal-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if subject.RemoveQuantity(nil, map[string]int{"u": 1}, "x", "u") {
		t.Error("nil items accepted")
	}
	items := map[string]int{"x": 1}
	if subject.RemoveQuantity(items, nil, "x", "u") || items["x"] != 1 {
		t.Errorf("nil units changed %v", items)
	}
}
