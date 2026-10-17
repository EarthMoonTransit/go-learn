package subject_test

import (
	"task.local/captured-method-values-001/subject"
	"testing"
)

func TestSavedReceivers(t *testing.T) {
	for _, initial := range []int{-7, 0, 11} {
		counter := subject.Counter{Value: initial}
		value, pointer := subject.Capture(&counter)
		if value == nil || pointer == nil {
			t.Fatal("Capture returned a nil function, want two non-nil saved calls")
		}
		for _, current := range []int{-9, 0, 23} {
			counter.Value = current
			if got := value(); got != initial {
				t.Errorf("value-bound call after mutation = %d, want captured %d", got, initial)
			}
			if got := pointer(); got != current {
				t.Errorf("pointer-bound call after mutation = %d, want current %d", got, current)
			}
			if counter.Value != current {
				t.Errorf("saved calls changed receiver to %d, want unchanged %d", counter.Value, current)
			}
		}
	}
}

func TestIndependentCaptures(t *testing.T) {
	a, b := subject.Counter{Value: 5}, subject.Counter{Value: 12}
	av, ap := subject.Capture(&a)
	bv, bp := subject.Capture(&b)
	if av == nil || ap == nil || bv == nil || bp == nil {
		t.Fatal("Capture returned nil, want independent non-nil calls")
	}
	a.Value, b.Value = 6, 13
	if got := av(); got != 5 {
		t.Errorf("first saved value = %d, want 5", got)
	}
	if got := ap(); got != 6 {
		t.Errorf("first saved pointer = %d, want 6", got)
	}
	if got := bv(); got != 12 {
		t.Errorf("second saved value = %d, want 12", got)
	}
	if got := bp(); got != 13 {
		t.Errorf("second saved pointer = %d, want 13", got)
	}
	later, _ := subject.Capture(&a)
	if later == nil {
		t.Fatal("later capture is nil, want a saved call")
	}
	a.Value = 8
	if got := later(); got != 6 {
		t.Errorf("later saved value = %d, want 6", got)
	}
	if got := av(); got != 5 {
		t.Errorf("older saved value = %d, want 5", got)
	}
}

func TestReadMethodsPreserveReceiver(t *testing.T) {
	c := subject.Counter{Value: 17}
	if got := c.ReadValue(); got != 17 {
		t.Errorf("ReadValue = %d, want 17", got)
	}
	if got := c.ReadPointer(); got != 17 {
		t.Errorf("ReadPointer = %d, want 17", got)
	}
	if c.Value != 17 {
		t.Errorf("method changed receiver to %d, want 17", c.Value)
	}
}
