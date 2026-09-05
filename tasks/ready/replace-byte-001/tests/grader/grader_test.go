package subject_test

import (
	"task.local/replace-byte-001/subject"
	"testing"
)

func TestReplacementBoundariesAndPreservation(t *testing.T) {
	cases := []struct {
		s     string
		index int
		value byte
		want  string
		ok    bool
	}{
		{"", 0, 'X', "", false}, {"abc", 3, 'X', "abc", false},
		{"abc", 8, 'X', "abc", false}, {"abc", -8, 'X', "abc", false},
		{"abc", 0, 'X', "Xbc", true}, {"abc", 2, 'X', "abX", true},
		{"a", 0, 0, "\x00", true}, {"ababa", 2, 'X', "abXba", true},
		{"\xc3\xa9!", 1, 'A', "\xc3A!", true},
		{"A\x00\xffZ", 0, 'B', "B\x00\xffZ", true},
		{"same", 2, 'm', "same", true},
	}
	for _, c := range cases {
		got, ok := subject.ReplaceByte(c.s, c.index, c.value)
		if got != c.want || ok != c.ok {
			t.Errorf("ReplaceByte(%q,%d,%d)=(%q,%v), want (%q,%v)", c.s, c.index, c.value, got, ok, c.want, c.ok)
		}
	}
}

func TestEveryReplacementByte(t *testing.T) {
	for value := 0; value < 256; value++ {
		got, ok := subject.ReplaceByte("pq", 0, byte(value))
		if !ok || len(got) != 2 || got[0] != byte(value) || got[1] != 'q' {
			t.Fatalf("replacement byte %d produced %q, %v", value, got, ok)
		}
	}
}
