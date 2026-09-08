package subject_test

import (
	"math"
	"task.local/widened-rate-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	for _, value := range []float64{-10000, 0, 555.55, 4999.99, 34600.80} {
		want := value * float64(subject.BandRate(value)) / 100
		got := subject.BandAmount(value)
		if math.IsNaN(got) || math.Abs(got-want) > math.Max(1e-8, math.Abs(want)*1e-12) {
			t.Errorf("BandAmount(%v) got %.15g want %.15g", value, got, want)
		}
	}
}
