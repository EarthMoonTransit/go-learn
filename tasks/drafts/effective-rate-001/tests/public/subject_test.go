package subject_test

import (
	"testing"

	"task.local/effective-rate-001/subject"
)

func TestEffectiveRatePreservesFractionalPercentage(t *testing.T) {
	if got := subject.EffectiveRate(101, 99.5); got != 100 {
		t.Fatalf("EffectiveRate(101, 99.5) = %d; want 100", got)
	}
}
