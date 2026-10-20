package subject_test

import (
	subject "task.local/mask-state-tests-001/subject"
	"testing"
)

func TestMaskStates(t *testing.T) {
	cases := []struct {
		bits, mask uint8
		want       subject.State
	}{
		{0, 0, subject.State{false, true, 0, 0, 0}},
		{255, 0, subject.State{false, true, 255, 255, 255}},
		{0, 255, subject.State{false, false, 255, 0, 255}},
		{255, 255, subject.State{true, true, 255, 0, 0}},
		{165, 60, subject.State{true, false, 189, 129, 153}},
		{195, 3, subject.State{true, true, 195, 192, 192}},
		{128, 1, subject.State{false, false, 129, 128, 129}},
		{85, 170, subject.State{false, false, 255, 85, 255}},
		{6, 5, subject.State{true, false, 7, 2, 3}},
	}
	for _, c := range cases {
		got := subject.Observe(c.bits, c.mask)
		if got != c.want {
			t.Errorf("Observe(%08b,%08b)=%+v; want %+v", c.bits, c.mask, got, c.want)
		}
	}
}
