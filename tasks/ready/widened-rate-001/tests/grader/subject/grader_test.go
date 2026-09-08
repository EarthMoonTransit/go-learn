package subject_test

import (
	"math"
	"task.local/widened-rate-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	value := 999999999.123
	want := value * float64(subject.BandRate(value)) / 100
	got := subject.BandAmount(value)
	if math.IsNaN(got) || math.Abs(got-want) > math.Max(1e-8, math.Abs(want)*1e-12) {
		t.Errorf("wide amount got %.15g want %.15g", got, want)
	}
}
