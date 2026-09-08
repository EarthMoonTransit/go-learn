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
	if got := (subject.Sample{}).Average(); math.IsNaN(got) || math.Abs(got) > 1e-9 {
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

func TestMethodFunctions(t *testing.T) {
	value, pointer := subject.MethodFunctions()
	if value == nil || pointer == nil {
		t.Fatal("method functions must both be non-nil")
	}
	sample := subject.Sample{Label: "x", Attempts: 4, Hits: 1}
	for _, hits := range []int{1, 3, 0, 4} {
		sample.Hits = hits
		want := float64(hits) / 4
		for _, got := range []float64{value(sample), pointer(&sample)} {
			if math.IsNaN(got) || math.Abs(got-want) > 1e-9 {
				t.Errorf("method function hits%d got %g want %g", hits, got, want)
			}
		}
		if sample != (subject.Sample{Label: "x", Attempts: 4, Hits: hits}) {
			t.Errorf("function changed argument: %v", sample)
		}
	}
	other := subject.Sample{Attempts: 2, Hits: 1}
	if got := value(other); math.IsNaN(got) || math.Abs(got-.5) > 1e-9 {
		t.Errorf("independent value got %g want.5", got)
	}
	if got := pointer(&other); math.IsNaN(got) || math.Abs(got-.5) > 1e-9 {
		t.Errorf("independent pointer got %g want.5", got)
	}
}
