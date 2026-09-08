package subject_test

import (
	"task.local/fractional-band-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.BandRate(0); got != float32(0.5) {
		t.Errorf("BandRate: got %v, want %v", got, float32(0.5))
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.BandRate(-0.000001); got != float32(3.213) {
		t.Errorf("BandRate: got %v, want %v", got, float32(3.213))
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.BandRate(999.99999); got != float32(0.5) {
		t.Errorf("BandRate: got %v, want %v", got, float32(0.5))
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.BandRate(1000); got != float32(1.621) {
		t.Errorf("BandRate: got %v, want %v", got, float32(1.621))
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.BandRate(4999.99999); got != float32(1.621) {
		t.Errorf("BandRate: got %v, want %v", got, float32(1.621))
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.BandRate(5000); got != float32(2.475) {
		t.Errorf("BandRate: got %v, want %v", got, float32(2.475))
	}
}
