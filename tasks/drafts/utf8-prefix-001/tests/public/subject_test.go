package subject_test

import (
	"testing"

	"task.local/utf8-prefix-001/subject"
)

func TestPrefixRunesCountsDecodedRunes(t *testing.T) {
	if got := subject.PrefixRunes("éx", 1); got != "é" {
		t.Fatalf("got %q; want %q", got, "é")
	}
}
