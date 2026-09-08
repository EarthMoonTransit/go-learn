package subject_test

import (
	"fmt"
	"task.local/number-box-tests-001/subject"
	"testing"
)

type numberBox struct{ value, calls int }

func (b *numberBox) Number() int { b.calls++; return b.value }

type alternateBox struct{}

func (alternateBox) Number() int { return 9 }

func TestContractCase01(t *testing.T) {
	b := &numberBox{value: 4}
	if got := subject.DescribeNumberBox(b); got != "box 4.0" || b.calls != 1 {
		t.Errorf("box got%q calls%d", got, b.calls)
	}
}

func TestContractCase02(t *testing.T) {
	b := &numberBox{value: -3}
	for _, n := range []int{-3, 0, 17} {
		b.value = n
		before := b.calls
		want := fmt.Sprintf("box %.1f", float64(n))
		if got := subject.DescribeNumberBox(b); got != want || b.calls != before+1 {
			t.Errorf("number%d got%q calls%d", n, got, b.calls)
		}
	}
	if got := subject.DescribeNumberBox(alternateBox{}); got != "box 9.0" {
		t.Errorf("alternate got%q", got)
	}
}
