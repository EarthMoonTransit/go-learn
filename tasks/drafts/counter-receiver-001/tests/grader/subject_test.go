package subject_test

import (
	"testing"

	"task.local/counter-receiver-001/subject"
)

func TestCounterAddAccumulates(t *testing.T) {
	counter := subject.Counter{Value: 10}
	counter.Add(4)
	counter.Add(-3)
	counter.Add(0)
	if counter.Value != 11 {
		t.Fatalf("Value = %d; want 11", counter.Value)
	}
}

func TestCounterInstancesRemainIndependent(t *testing.T) {
	first := subject.Counter{Value: 1}
	second := subject.Counter{Value: 8}
	first.Add(2)
	second.Add(-2)
	if first.Value != 3 || second.Value != 6 {
		t.Fatalf("values = %d, %d", first.Value, second.Value)
	}
}
