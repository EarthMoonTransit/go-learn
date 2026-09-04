package subject_test

import (
	"task.local/literal-bindings-001/subject"
	"testing"
)

func TestAllBindings(t *testing.T) {
	count, scale, label, enabled := subject.Snapshot()
	if v, ok := count.(int); !ok || v != 12 {
		t.Errorf("count has incorrect type or value")
	}
	if v, ok := scale.(float64); !ok || v != 1.25 {
		t.Errorf("scale has incorrect type or value")
	}
	if v, ok := label.(string); !ok || v != "G" {
		t.Errorf("label has incorrect type or value")
	}
	if v, ok := enabled.(bool); !ok || !v {
		t.Errorf("enabled has incorrect type or value")
	}
}
