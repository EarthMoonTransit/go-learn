package subject_test

import (
	"testing"

	"task.local/formatter-contract-001/subject"
)

type fixedFormatter string

func (value fixedFormatter) Format() string { return string(value) }

func TestRenderAcceptsCallerDefinedFormatter(t *testing.T) {
	if got := subject.Render(fixedFormatter("custom")); got != "custom" {
		t.Fatalf("Render returned %q", got)
	}
}
