package subject_test

import (
	"task.local/whole-minute-output-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.SuccessfulPerMinute(0, 100); got != 0 {
		t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 0)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.SuccessfulPerMinute(221, 100); got != 3 {
		t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 3)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.SuccessfulPerMinute(426, 80); got != 5 {
		t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 5)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.SuccessfulPerMinute(6824, 20.5); got != 23 {
		t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 23)
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.SuccessfulPerMinute(8000, 0); got != 0 {
		t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 0)
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.SuccessfulPerMinute(119, 100); got != 1 {
		t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 1)
	}
}

func TestPublicCase07(t *testing.T) {
	if got := subject.SuccessfulPerMinute(359, 50.5); got != 3 {
		t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 3)
	}
}
