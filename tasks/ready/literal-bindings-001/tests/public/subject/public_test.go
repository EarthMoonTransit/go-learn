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

func TestTextAndBooleanBindings(t *testing.T) {
	_, _, label, enabled := subject.Snapshot()
	if v, ok := label.(string); !ok || v != "G" {
		t.Errorf("label = %v (%T), want string G", label, label)
	}
	if v, ok := enabled.(bool); !ok || !v {
		t.Errorf("enabled = %v (%T), want bool true", enabled, enabled)
	}
}
