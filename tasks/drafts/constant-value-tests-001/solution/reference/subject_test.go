package subject_test

import (
	"math"
	"task.local/constant-value-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if subject.Rate != 0.05 {
		t.Errorf("Rate got%v want.05", subject.Rate)
	}
}

func TestContractCase02(t *testing.T) {
	if subject.Days != 365 || subject.DayCount() != 365 {
		t.Errorf("days constant%d function%d want365", subject.Days, subject.DayCount())
	}
	if got := subject.Rate32(); got != float32(.05) {
		t.Errorf("Rate32 got%g want%g", got, float32(.05))
	}
}

func TestContractCase03(t *testing.T) {
	for _, days := range []int{0, 1, 365, 730} {
		want := float64(days) / 365
		got := subject.AnnualFraction(days)
		if math.IsNaN(got) || math.Abs(got-want) > 1e-9 {
			t.Errorf("days%d got%g want%g", days, got, want)
		}
	}
}
