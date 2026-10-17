package subject_test

import (
	"task.local/conditional-binding-001/subject"
	"testing"
)

func TestVariedConditionalCalls(t *testing.T) {
	for initial := -21; initial <= 21; initial++ {
		for next := -19; next <= 19; next++ {
			if got := subject.SelectValue(initial, next, true); got != next {
				t.Fatalf("SelectValue(%d, %d, true) = %d, want %d", initial, next, got, next)
			}
			if got := subject.SelectValue(initial, next, false); got != initial {
				t.Fatalf("SelectValue(%d, %d, false) = %d, want %d", initial, next, got, initial)
			}
		}
	}
	max := int(^uint(0) >> 1)
	min := -max - 1
	if got := subject.SelectValue(min, max, true); got != max {
		t.Errorf("extreme true = %d, want %d", got, max)
	}
	if got := subject.SelectValue(min, max, false); got != min {
		t.Errorf("extreme false = %d, want %d", got, min)
	}
}
