package subject_test

import (
	"task.local/boolean-policies-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.AnyActive(true, true, true); got != true {
		t.Errorf("AnyActive: got %v, want %v", got, true)
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.SignalAllowed(true, false); got != false {
		t.Errorf("SignalAllowed: got %v, want %v", got, false)
	}
}

func TestGraderCase03(t *testing.T) {
	if got := subject.ReleaseAllowed(false, false, false, true); got != true {
		t.Errorf("ReleaseAllowed: got %v, want %v", got, true)
	}
}

func TestGraderCase04(t *testing.T) {
	if got := subject.ReleaseAllowed(true, true, true, false); got != false {
		t.Errorf("ReleaseAllowed: got %v, want %v", got, false)
	}
}
