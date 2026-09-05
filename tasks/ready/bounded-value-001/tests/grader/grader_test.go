package subject_test

import (
	"task.local/bounded-value-001/subject"
	"testing"
)

func TestAllSmallIntervals(t *testing.T) {
	for lo := -5; lo <= 5; lo++ {
		for hi := -5; hi <= 5; hi++ {
			for v := -7; v <= 7; v++ {
				got, ok := subject.Bounded(v, lo, hi)
				if lo > hi {
					if got != 0 || ok {
						t.Fatal("inverted interval accepted")
					}
					continue
				}
				if !ok || got < lo || got > hi {
					t.Fatalf("invalid result (%d,%v) for [%d,%d]", got, ok, lo, hi)
				}
				if v >= lo && v <= hi && got != v {
					t.Fatal("interior value changed")
				}
				if v < lo && got != lo || v > hi && got != hi {
					t.Fatal("incorrect limiting endpoint")
				}
			}
		}
	}
}
func TestExtremeEndpoints(t *testing.T) {
	max := int(^uint(0) >> 1)
	min := -max - 1
	for _, c := range []struct {
		v, lo, hi, w int
		ok           bool
	}{{min, min, max, min, true}, {max, min, max, max, true}, {0, min, max, 0, true}, {min, max, max, max, true}, {max, min, min, min, true}, {0, max, min, 0, false}} {
		got, ok := subject.Bounded(c.v, c.lo, c.hi)
		if got != c.w || ok != c.ok {
			t.Fatalf("incorrect extreme endpoint result (%d,%v)", got, ok)
		}
	}
}
