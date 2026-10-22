package subject_test

import (
	subject "task.local/replace-rune-range-001/subject"
	"testing"
)

func TestEverySmallRuneWindow(t *testing.T) {
	for _, s := range []string{"", "a", "界🙂éz"} {
		r := []rune(s)
		for lo := 0; lo <= len(r); lo++ {
			for hi := lo; hi <= len(r); hi++ {
				for _, replacement := range []string{"", "Ω", "xy界"} {
					want := string(r[:lo]) + replacement + string(r[hi:])
					got, ok := subject.ReplaceRuneRange(s, lo, hi, replacement)
					if !ok || got != want {
						t.Fatalf("s=%q range=%d:%d replacement=%q: (%q,%t); want (%q,true)", s, lo, hi, replacement, got, ok, want)
					}
				}
			}
		}
	}
}
