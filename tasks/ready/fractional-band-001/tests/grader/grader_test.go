package subject_test

import (
	"task.local/fractional-band-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.BandRate(5639998.742909); got != float32(2.475) {
		t.Errorf("BandRate: got %v, want %v", got, float32(2.475))
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.BandRate(-152964.231); got != float32(3.213) {
		t.Errorf("BandRate: got %v, want %v", got, float32(3.213))
	}
}
