package subject_test

import (
	"task.local/gcd-reduction-001/subject"
	"testing"
)

// A descending divisor scan is independent of the reference's remainder loop.
func divisorOracle(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	limit := a
	if b < limit {
		limit = b
	}
	for d := limit; d > 0; d-- {
		if a%d == 0 && b%d == 0 {
			return d
		}
	}
	return 0
}

func TestGCDBoundedDomainAndLargePairs(t *testing.T) {
	for a := 0; a <= 64; a++ {
		for b := 0; b <= 64; b++ {
			want := divisorOracle(a, b)
			if got := subject.GCD(a, b); got != want {
				t.Fatalf("GCD(%d, %d) = %d, want %d", a, b, got, want)
			}
		}
	}
	for _, c := range []struct{ a, b, want int }{
		{1000000, 0, 1000000}, {0, 1000000, 1000000}, {1000000, 999999, 1},
		{999999, 1000000, 1}, {750000, 500000, 250000}, {500000, 750000, 250000},
		{832040, 514229, 1}, {1000000, 1000000, 1000000},
	} {
		if got := subject.GCD(c.a, c.b); got != c.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
