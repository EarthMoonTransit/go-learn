package subject_test

import (
	"task.local/typed-line-001/subject"
	"testing"
)

func TestTypedStatus(t *testing.T) {
	got := subject.FormatStatus("cache", 65, true)
	if got != "label=cache; count=65; enabled=true" {
		t.Fatalf("status=%q", got)
	}
}
func TestNegativeStatus(t *testing.T) {
	got := subject.FormatStatus("", -7, false)
	if got != "label=; count=-7; enabled=false" {
		t.Fatalf("status=%q", got)
	}
}
