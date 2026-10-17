package subject_test

import (
	"task.local/lazy-probe-001/subject"
	"testing"
)

func TestProbeGate(t *testing.T) {
	for _, c := range []struct {
		name                 string
		enabled, value, want bool
		calls                int
	}{
		{"disabled/false", false, false, false, 0},
		{"disabled/true", false, true, false, 0},
		{"enabled/false", true, false, false, 1},
		{"enabled/true", true, true, true, 1},
	} {
		t.Run(c.name, func(t *testing.T) {
			subject.ResetProbe(c.value)
			got := subject.CheckEnabled(c.enabled)
			calls := subject.ProbeCalls()
			if got != c.want {
				t.Errorf("CheckEnabled(%t) with Probe=%t returned %t, want %t", c.enabled, c.value, got, c.want)
			}
			if calls != c.calls {
				t.Errorf("Probe calls = %d, want %d: disabled must avoid eager calls; enabled must call exactly once", calls, c.calls)
			}
		})
	}
}
