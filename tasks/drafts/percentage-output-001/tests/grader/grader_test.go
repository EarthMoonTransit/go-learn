package subject_test

import (
	"math"
	"task.local/percentage-output-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.SuccessfulPerHour(7, 12.5); math.IsNaN(got) || math.Abs(got-(0.875)) > 1e-9 {
		t.Errorf("SuccessfulPerHour: got %.12g want %.12g", got, float64(0.875))
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.SuccessfulPerHour(999999, 33.3); math.IsNaN(got) || math.Abs(got-(332999.667)) > 1e-9 {
		t.Errorf("SuccessfulPerHour: got %.12g want %.12g", got, float64(332999.667))
	}
}
