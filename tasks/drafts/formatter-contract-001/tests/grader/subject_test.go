package subject_test

import (
	"testing"

	"task.local/formatter-contract-001/subject"
)

type wrappedFormatter struct{ value string }

func (value wrappedFormatter) Format() string { return "[" + value.value + "]" }

func TestPairFormat(t *testing.T) {
	pair := subject.Pair{Left: "left", Right: "right"}
	if got := pair.Format(); got != "left:right" {
		t.Fatalf("Pair.Format returned %q", got)
	}
	var formatter subject.Formatter = pair
	if got := formatter.Format(); got != "left:right" {
		t.Fatalf("interface call returned %q", got)
	}
}

func TestRenderDelegatesToInterface(t *testing.T) {
	if got := subject.Render(wrappedFormatter{value: "x"}); got != "[x]" {
		t.Fatalf("Render returned %q", got)
	}
}
