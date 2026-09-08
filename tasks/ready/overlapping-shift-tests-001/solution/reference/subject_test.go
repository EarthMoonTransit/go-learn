package subject_test

import (
	"slices"
	"task.local/overlapping-shift-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	v := []int{1, 2, 3, 4}
	alias := v
	n := subject.ShiftRight(v, 1)
	if n != 4 || !slices.Equal(v, []int{1, 0, 2, 3}) || !slices.Equal(alias, v) {
		t.Errorf("shift got%v removed%d", v, n)
	}
}

func TestContractCase02(t *testing.T) {
	for _, c := range []struct {
		in   []int
		i    int
		want []int
		last int
	}{{[]int{-1, 2, 3}, 0, []int{0, -1, 2}, 3}, {[]int{7}, 0, []int{0}, 7}, {[]int{4, 4, -2}, 2, []int{4, 4, 0}, -2}} {
		n := subject.ShiftRight(c.in, c.i)
		if n != c.last || !slices.Equal(c.in, c.want) {
			t.Errorf("index%d got%v,%d want%v,%d", c.i, c.in, n, c.want, c.last)
		}
	}
}
