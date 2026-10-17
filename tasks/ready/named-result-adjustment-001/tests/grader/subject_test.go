package subject_test

import (
	"task.local/named-result-adjustment-001/subject"
	"testing"
)

func TestAdjustmentDomain(t *testing.T) {
	for base := -50; base <= 50; base++ {
		for delta := -20; delta <= 20; delta++ {
			caller, calls := 123, 0
			got := subject.Adjust(base, delta, &caller, func() { calls++ })
			if want := base + 3 + delta; got != want || caller != 123 || calls != 1 {
				t.Fatalf("Adjust(%d,%d) = %d/%d/%d, want %d/123/1", base, delta, got, caller, calls, want)
			}
		}
	}
}
