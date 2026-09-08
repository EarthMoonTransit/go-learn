package subject_test

import (
	"task.local/number-box-tests-001/subject"
	"testing"
)

type numberBox struct{ value, calls int }

func (b *numberBox) Number() int { b.calls++; return b.value }

type alternateBox struct{}

func (alternateBox) Number() int { return 9 }

func TestStartingCase01(t *testing.T) {
	b := &numberBox{value: 4}
	if got := subject.DescribeNumberBox(b); got != "box 4.0" || b.calls != 1 {
		t.Errorf("box got%q calls%d", got, b.calls)
	}
}

// Add cases for the remaining requirements in TASK.md.
