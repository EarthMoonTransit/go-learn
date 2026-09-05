package subject_test

import (
	"task.local/byte-window-tests-001/subject"
	"testing"
)

func TestByteWindowContract(t *testing.T) {
	cases := []struct {
		s          string
		start, end int
		want       string
		ok         bool
	}{
		{"header", 1, 4, "ead", true}, {"xyz", 0, 3, "xyz", true},
		{"", 0, 0, "", true}, {"ab", 2, 2, "", true}, {"abcd", 2, 2, "", true},
		{"ab", -1, 1, "", false}, {"ab", 0, -1, "", false},
		{"ab", 0, 3, "", false}, {"ab", 3, 3, "", false}, {"abcd", 3, 1, "", false},
		{"A\x00\xffZ", 1, 3, "\x00\xff", true},
		{"\xc3\xa9X", 0, 1, "\xc3", true}, {"\xc3\xa9X", 1, 2, "\xa9", true},
	}
	for _, c := range cases {
		got, ok := subject.ByteWindow(c.s, c.start, c.end)
		if got != c.want || ok != c.ok {
			t.Errorf("ByteWindow(%q,%d,%d)=(%q,%v), want (%q,%v)", c.s, c.start, c.end, got, ok, c.want, c.ok)
		}
	}
}
