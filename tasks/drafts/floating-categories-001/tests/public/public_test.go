package subject_test

import (
	"math"
	"task.local/floating-categories-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		v    float64
		want string
	}{{math.NaN(), "nan"}, {math.Float64frombits(0xfff8000000000001), "nan"}, {math.Inf(1), "positive-infinity"}, {math.Inf(-1), "negative-infinity"}} {
		if got := subject.FloatKind(c.v); got != c.want {
			t.Errorf("kind got%q want%q", got, c.want)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	for _, v := range []float64{0, math.Copysign(0, -1), -1, 1, math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64, math.MaxFloat64, -math.MaxFloat64} {
		if got := subject.FloatKind(v); got != "finite" {
			t.Errorf("finite %g got%q", v, got)
		}
	}
}
