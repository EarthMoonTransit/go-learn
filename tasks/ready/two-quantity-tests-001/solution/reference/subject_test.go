package subject_test

import (
	"math"
	"slices"
	"task.local/two-quantity-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	gotA, gotB := subject.Quantities([]string{"solid", "liquid", "solid"})
	if gotA != 100 || math.IsNaN(gotB) || math.Abs(gotB-0.2) > 1e-9 {
		t.Errorf("Quantities got (%d,%g) want (100,0.2)", gotA, gotB)
	}
}

func TestContractCase02(t *testing.T) {
	gotA, gotB := subject.Quantities([]string{"liquid", "other", "liquid"})
	if gotA != 0 || math.IsNaN(gotB) || math.Abs(gotB-0.4) > 1e-9 {
		t.Errorf("Quantities got (%d,%g) want (0,0.4)", gotA, gotB)
	}
}

func TestContractCase03(t *testing.T) {
	gotA, gotB := subject.Quantities([]string{"solid", "other", "solid"})
	if gotA != 100 || math.IsNaN(gotB) || math.Abs(gotB-0) > 1e-9 {
		t.Errorf("Quantities got (%d,%g) want (100,0)", gotA, gotB)
	}
}

func TestContractCase04(t *testing.T) {
	gotA, gotB := subject.Quantities(nil)
	if gotA != 0 || math.IsNaN(gotB) || math.Abs(gotB-0) > 1e-9 {
		t.Errorf("Quantities got (%d,%g) want (0,0)", gotA, gotB)
	}
}

func TestContractCase05(t *testing.T) {
	gotA, gotB := subject.Quantities([]string{"SOLID", "LIQUID", ""})
	if gotA != 0 || math.IsNaN(gotB) || math.Abs(gotB-0) > 1e-9 {
		t.Errorf("Quantities got (%d,%g) want (0,0)", gotA, gotB)
	}
}

func TestContractCase06(t *testing.T) {
	in := []string{"solid", "liquid"}
	subject.Quantities(in)
	if !slices.Equal(in, []string{"solid", "liquid"}) {
		t.Errorf("input changed: %q", in)
	}
}
