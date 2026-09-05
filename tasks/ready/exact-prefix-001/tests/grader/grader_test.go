package subject_test

import (
	"task.local/exact-prefix-001/subject"
	"testing"
)

func TestPrefixExactBoundaries(t *testing.T) {
	cases := []struct {
		s, prefix, want string
		ok              bool
	}{
		{"", "", "", true}, {"", "x", "", false}, {"ab", "abcd", "ab", false},
		{"all", "all", "", true}, {"preprex", "pre", "prex", true},
		{"xpre", "pre", "xpre", false}, {"abZ", "Ab", "abZ", false},
		{" x ", " ", "x ", true}, {"x ", " ", "x ", false},
		{"\x00\xffA", "\x00\xff", "A", true},
		{"aaaa", "aa", "aa", true}, {"plain", "none", "plain", false},
	}
	for _, c := range cases {
		got, ok := subject.RemovePrefix(c.s, c.prefix)
		if got != c.want || ok != c.ok {
			t.Errorf("RemovePrefix(%q,%q)=(%q,%v), want (%q,%v)", c.s, c.prefix, got, ok, c.want, c.ok)
		}
	}
}
