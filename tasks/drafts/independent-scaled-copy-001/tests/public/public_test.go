package subject_test

import (
	"math"
	"slices"
	"task.local/independent-scaled-copy-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	in := []float64{0.5, 250, 150, 3, 0.5}
	want := []float64{1.5, 750, 450, 9, 1.5}
	got := subject.ScaledCopy(in, 6)
	if !slices.Equal(got, want) {
		t.Errorf("scaled got %v want %v", got, want)
	}
	if !slices.Equal(in, []float64{0.5, 250, 150, 3, 0.5}) {
		t.Errorf("input changed: %v", in)
	}
}

func TestPublicCase02(t *testing.T) {
	got := subject.ScaledCopy([]float64{0.6, 1, 0.5}, 3)
	want := []float64{0.9, 1.5, 0.75}
	if len(got) != len(want) {
		t.Fatalf("length got %d want3", len(got))
	}
	for i := range want {
		if math.IsNaN(got[i]) || math.Abs(got[i]-want[i]) > 1e-9 {
			t.Errorf("value%d got %g want %g", i, got[i], want[i])
		}
	}
}

func TestPublicCase03(t *testing.T) {
	for _, units := range []int{0, 2} {
		in := []float64{4, 2}
		got := subject.ScaledCopy(in, units)
		if len(got) != 2 {
			t.Fatalf("length got %d want2", len(got))
		}
		want := []float64{4 * float64(units) / 2, 2 * float64(units) / 2}
		if !slices.Equal(got, want) {
			t.Errorf("units%d got %v want %v", units, got, want)
		}
		got[0] = 99
		if in[0] != 4 || in[1] != 2 {
			t.Errorf("units%d shared or modified input %v", units, in)
		}
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.ScaledCopy(nil, 100); len(got) != 0 {
		t.Errorf("nil got %v want empty", got)
	}
}
