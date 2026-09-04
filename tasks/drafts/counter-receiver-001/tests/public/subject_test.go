package subject_test

import (
	"testing"

	"task.local/counter-receiver-001/subject"
)

func TestCounterAddMutatesOriginal(t *testing.T) {
	counter := subject.Counter{Value: 2}
	counter.Add(3)
	if counter.Value != 5 {
		t.Fatalf("Value = %d; want 5", counter.Value)
	}
}
