package subject_test

import (
	"task.local/typed-line-001/subject"
	"testing"
)

func TestStatusExactRepresentations(t *testing.T) {
	cases := []struct {
		label   string
		count   int
		enabled bool
		want    string
	}{
		{"", 0, true, "label=; count=0; enabled=true"},
		{"%s%%", 123, false, "label=%s%%; count=123; enabled=false"},
		{" a b ", -120, true, "label= a b ; count=-120; enabled=true"},
		{"a\"b\\c", 32, false, "label=a\"b\\c; count=32; enabled=false"},
		{"x\ny", 4, true, "label=x\ny; count=4; enabled=true"},
		{"\x00\xff", 8, false, "label=\x00\xff; count=8; enabled=false"},
	}
	for _, c := range cases {
		if got := subject.FormatStatus(c.label, c.count, c.enabled); got != c.want {
			t.Errorf("FormatStatus(%q,%d,%v)=%q, want %q", c.label, c.count, c.enabled, got, c.want)
		}
	}
}
