package subject_test

import (
	"task.local/formatter-contract-001/subject"
	"testing"
)

type wrappedFormatter struct{ value string }

func (value wrappedFormatter) Format() string { return "[" + value.value + "]" }

func TestFormattingBoundaries(t *testing.T) {
	for _, left := range []string{"", ":", " Go ", "λ"} {
		for _, right := range []string{"", ":", " Go ", "λ"} {
			pair := subject.Pair{Left: left, Right: right}
			want := left + ":" + right
			if got := subject.Render(pair); got != want {
				t.Errorf("Render pair = %q, want %q", got, want)
			}
		}
	}
	if got := subject.Render(wrappedFormatter{value: "x"}); got != "[x]" {
		t.Errorf("custom wrapped result = %q, want [x]", got)
	}
}
