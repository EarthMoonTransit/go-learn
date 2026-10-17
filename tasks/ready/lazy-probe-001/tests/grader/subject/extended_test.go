package subject_test

import (
	"task.local/lazy-probe-001/subject"
	"testing"
)

func TestProbeSequence(t *testing.T) {
	for _, value := range []bool{true, false, true, false} {
		subject.ResetProbe(value)
		calls := 0
		for _, enabled := range []bool{false, true, false, true, true, false} {
			want := false
			if enabled {
				calls++
				want = value
			}
			if got := subject.CheckEnabled(enabled); got != want {
				t.Errorf("sequence enabled=%t Probe=%t: got %t, want %t", enabled, value, got, want)
			}
			if got := subject.ProbeCalls(); got != calls {
				t.Errorf("sequence calls = %d, want %d", got, calls)
			}
		}
	}
}
