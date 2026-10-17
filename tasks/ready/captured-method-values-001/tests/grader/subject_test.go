package subject_test

import (
	"task.local/captured-method-values-001/subject"
	"testing"
)

func TestSavedReceiverBoundaries(t *testing.T) {
	max := int(^uint(0) >> 1)
	values := []int{-max - 1, -999, -1, 0, 1, 999, max}
	for _, initial := range values {
		c := subject.Counter{Value: initial}
		v, p := subject.Capture(&c)
		if v == nil || p == nil {
			t.Fatal("Capture returned nil, want saved functions")
		}
		for _, current := range values {
			c.Value = current
			if got := v(); got != initial {
				t.Errorf("saved value = %d, want %d", got, initial)
			}
			if got := p(); got != current {
				t.Errorf("saved pointer = %d, want %d", got, current)
			}
			if c.Value != current {
				t.Errorf("receiver = %d, want %d", c.Value, current)
			}
		}
	}
}
