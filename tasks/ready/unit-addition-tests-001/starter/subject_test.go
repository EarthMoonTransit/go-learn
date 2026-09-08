package subject_test

import (
	"task.local/unit-addition-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	v := map[string]int{}
	u := map[string]int{"dozen": 12}
	if !subject.AddUnit(v, u, "item", "dozen") || v["item"] != 12 {
		t.Errorf("add got %v want item:12 and true", v)
	}
}

// Add cases for the remaining requirements in TASK.md.
