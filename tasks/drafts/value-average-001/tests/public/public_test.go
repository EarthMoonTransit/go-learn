package subject_test

import (
	"math"
	"task.local/value-average-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	s := subject.Sample{Label: "a", Attempts: 4, Hits: 1}
	got := s.Average()
	if math.IsNaN(got) || math.Abs(got-0.25) > 1e-9 {
		t.Errorf("average got %g want .25", got)
	}
	if s != (subject.Sample{Label: "a", Attempts: 4, Hits: 1}) {
		t.Errorf("receiver changed: %v", s)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := (subject.Sample{}).Average(); got != 0 {
		t.Errorf("zero got %g want0", got)
	}
}

func TestPublicCase03(t *testing.T) {
	s := subject.Sample{Attempts: 3, Hits: 2}
	for _, got := range []float64{s.Average(), (&s).Average(), (subject.Sample{Attempts: 3, Hits: 2}).Average()} {
		if math.IsNaN(got) || math.Abs(got-2.0/3) > 1e-9 {
			t.Errorf("value/pointer got %g want2/3", got)
		}
	}
}
