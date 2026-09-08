package subject_test

import (
	"math"
	"task.local/percentage-output-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.SuccessfulPerHour(0, 100); math.IsNaN(got) || math.Abs(got-(0)) > 1e-9 {
		t.Errorf("SuccessfulPerHour: got %.12g want %.12g", got, float64(0))
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.SuccessfulPerHour(221, 100); math.IsNaN(got) || math.Abs(got-(221)) > 1e-9 {
		t.Errorf("SuccessfulPerHour: got %.12g want %.12g", got, float64(221))
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.SuccessfulPerHour(426, 80); math.IsNaN(got) || math.Abs(got-(340.8)) > 1e-9 {
		t.Errorf("SuccessfulPerHour: got %.12g want %.12g", got, float64(340.8))
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.SuccessfulPerHour(6824, 20.5); math.IsNaN(got) || math.Abs(got-(1398.92)) > 1e-9 {
		t.Errorf("SuccessfulPerHour: got %.12g want %.12g", got, float64(1398.92))
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.SuccessfulPerHour(8000, 0); math.IsNaN(got) || math.Abs(got-(0)) > 1e-9 {
		t.Errorf("SuccessfulPerHour: got %.12g want %.12g", got, float64(0))
	}
}
