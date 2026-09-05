package subject_test

import (
	"task.local/exact-prefix-001/subject"
	"testing"
)

func TestExactPrefixMatch(t *testing.T) {
	got, ok := subject.RemovePrefix("abacus", "ab")
	if !ok || got != "acus" {
		t.Fatalf("match=(%q,%v), want (acus,true)", got, ok)
	}
}
func TestPrefixIsNotACharacterSet(t *testing.T) {
	got, ok := subject.RemovePrefix("ba", "ab")
	if ok || got != "ba" {
		t.Fatalf("nonmatch=(%q,%v), want original and false", got, ok)
	}
}
func TestEmptyPrefixMatches(t *testing.T) {
	got, ok := subject.RemovePrefix("value", "")
	if !ok || got != "value" {
		t.Fatalf("empty prefix=(%q,%v)", got, ok)
	}
}
