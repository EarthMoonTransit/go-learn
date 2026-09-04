package subject_test

import (
	"strings"
	"task.local/encoded-width-001/subject"
	"testing"
)

func TestExactTagsAndDefault(t *testing.T) {
	tags := []string{"byte", "word", "long", "quad"}
	widths := []int{1, 2, 4, 8}
	for i, tag := range tags {
		got, ok := subject.Width(tag)
		if got != widths[i] || !ok {
			t.Fatalf("incorrect recognized tag %q", tag)
		}
		for _, bad := range []string{" " + tag, tag + " ", strings.ToUpper(tag), tag + "s", tag + "\x00", tag + "\n"} {
			got, ok := subject.Width(bad)
			if got != 0 || ok {
				t.Errorf("unexpected normalized tag %q", bad)
			}
		}
	}
	for _, bad := range []string{"", "unknown", "Byte", "4", "λ"} {
		got, ok := subject.Width(bad)
		if got != 0 || ok {
			t.Errorf("unknown tag %q accepted", bad)
		}
	}
}
