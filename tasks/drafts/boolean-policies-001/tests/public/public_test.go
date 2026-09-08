package subject_test

import (
	"task.local/boolean-policies-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.FastAllowed(true); got != false {
		t.Errorf("FastAllowed: got %v, want %v", got, false)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.FastAllowed(false); got != true {
		t.Errorf("FastAllowed: got %v, want %v", got, true)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.AnyActive(false, false, false); got != false {
		t.Errorf("AnyActive: got %v, want %v", got, false)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.AnyActive(true, false, false); got != true {
		t.Errorf("AnyActive: got %v, want %v", got, true)
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.AnyActive(false, true, false); got != true {
		t.Errorf("AnyActive: got %v, want %v", got, true)
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.AnyActive(false, false, true); got != true {
		t.Errorf("AnyActive: got %v, want %v", got, true)
	}
}

func TestPublicCase07(t *testing.T) {
	if got := subject.SignalAllowed(false, true); got != true {
		t.Errorf("SignalAllowed: got %v, want %v", got, true)
	}
}

func TestPublicCase08(t *testing.T) {
	if got := subject.SignalAllowed(true, true); got != false {
		t.Errorf("SignalAllowed: got %v, want %v", got, false)
	}
}

func TestPublicCase09(t *testing.T) {
	if got := subject.SignalAllowed(false, false); got != false {
		t.Errorf("SignalAllowed: got %v, want %v", got, false)
	}
}

func TestPublicCase10(t *testing.T) {
	if got := subject.ReleaseAllowed(true, false, false, true); got != true {
		t.Errorf("ReleaseAllowed: got %v, want %v", got, true)
	}
}

func TestPublicCase11(t *testing.T) {
	if got := subject.ReleaseAllowed(false, true, true, true); got != false {
		t.Errorf("ReleaseAllowed: got %v, want %v", got, false)
	}
}

func TestPublicCase12(t *testing.T) {
	if got := subject.ReleaseAllowed(false, false, true, false); got != true {
		t.Errorf("ReleaseAllowed: got %v, want %v", got, true)
	}
}

func TestPublicCase13(t *testing.T) {
	if got := subject.ReleaseAllowed(true, false, true, false); got != false {
		t.Errorf("ReleaseAllowed: got %v, want %v", got, false)
	}
}

func TestPublicCase14(t *testing.T) {
	if got := subject.ReleaseAllowed(false, false, false, false); got != false {
		t.Errorf("ReleaseAllowed: got %v, want %v", got, false)
	}
}
