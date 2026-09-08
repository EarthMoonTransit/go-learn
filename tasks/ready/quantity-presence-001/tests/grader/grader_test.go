package subject_test

import (
	"task.local/quantity-presence-001/subject"
	"testing"
)

func TestGraderEmpty(t *testing.T) {
	n, ok := subject.Quantity(map[string]int{}, "")
	if n != 0 || ok {
		t.Errorf("empty map got (%d,%v) want (0,false)", n, ok)
	}
}
