package subject_test

import (
	"testing"
	"time"
)

func TestGraderDurationUnits(t *testing.T) {
	budgetRun(t, 1700*time.Microsecond, 7200*time.Microsecond)
	budgetRun(t, 9*time.Millisecond, 2*time.Millisecond)
	slowPeriodic(t, 200*time.Microsecond, 1100*time.Microsecond, 250*time.Microsecond, 550*time.Microsecond)
}
