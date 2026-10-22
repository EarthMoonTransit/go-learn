package subject_test

import (
	subject "task.local/replace-rune-range-001/subject"
	"testing"
)

func TestRuneSplicing(t *testing.T) {
	cases := []struct {
		s                 string
		lo, hi            int
		replacement, want string
	}{
		{"A界🙂Z", 1, 3, "éΩ", "AéΩZ"}, {"A界🙂Z", 2, 2, "漢字", "A界漢字🙂Z"}, {"A界🙂Z", 4, 4, "!", "A界🙂Z!"},
		{"A界🙂Z", 1, 3, "", "AZ"}, {"abc", 0, 3, "λ", "λ"}, {"", 0, 0, "界🙂", "界🙂"}, {"", 0, 0, "", ""},
		{"é界", 1, 2, "!", "e!界"}, {"left", 0, 0, "前", "前left"}, {"A界🙂Z", 1, 2, "XY", "AXY🙂Z"},
	}
	for _, c := range cases {
		got, ok := subject.ReplaceRuneRange(c.s, c.lo, c.hi, c.replacement)
		if !ok || got != c.want {
			t.Fatalf("ReplaceRuneRange(%q,%d,%d,%q)=(%q,%t); want (%q,true)", c.s, c.lo, c.hi, c.replacement, got, ok, c.want)
		}
	}
}
func TestInvalidRuneBounds(t *testing.T) {
	s := "A界🙂Z"
	max := int(^uint(0) >> 1)
	for _, b := range [][2]int{{-1, 1}, {0, -1}, {3, 2}, {0, 5}, {5, 5}, {0, max}, {max, max}} {
		got, ok := subject.ReplaceRuneRange(s, b[0], b[1], "x")
		if ok || got != s {
			t.Fatalf("invalid bounds=%v: got (%q,%t); want (%q,false)", b, got, ok, s)
		}
	}
}
