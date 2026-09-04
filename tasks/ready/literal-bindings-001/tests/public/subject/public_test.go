package subject_test

import (
	"task.local/literal-bindings-001/subject"
	"testing"
)

func TestNumericBindings(t *testing.T) {
	count, scale, _, _ := subject.Snapshot()
	if v, ok := count.(int); !ok || v != 12 {
		t.Errorf("count = %v (%T), want int 12", count, count)
	}
	if v, ok := scale.(float64); !ok || v != 1.25 {
		t.Errorf("scale = %v (%T), want float64 1.25", scale, scale)
	}
}
