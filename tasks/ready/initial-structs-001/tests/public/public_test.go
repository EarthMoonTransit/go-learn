package subject_test

import (
	"task.local/initial-structs-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	want := subject.Motor{Speed: 5, Drain: 2, Battery: 100}
	if got := subject.NewMotor(5, 2); got != want {
		t.Errorf("NewMotor got %+v want %+v", got, want)
	}
}

func TestPublicCase02(t *testing.T) {
	want := subject.Motor{Battery: 100}
	if got := subject.NewMotor(0, 0); got != want {
		t.Errorf("zero inputs got %+v want %+v", got, want)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.NewTrack(800); got.Distance != 800 {
		t.Errorf("track got %+v want distance800", got)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.NewTrack(0); got.Distance != 0 {
		t.Errorf("zero track got %+v", got)
	}
}
