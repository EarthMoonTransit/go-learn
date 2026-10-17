package subject_test

import (
	"task.local/formatter-contract-001/subject"
	"testing"
)

type fixedFormatter string

func (value fixedFormatter) Format() string { return string(value) }

type countedFormatter struct {
	calls  int
	output string
}

func (value *countedFormatter) Format() string { value.calls++; return value.output }

func TestPairAndImplicitContract(t *testing.T) {
	for _, c := range []struct{ left, right, want string }{{"L", "R", "L:R"}, {"", "", ":"}, {"x:", " y ", "x:: y "}} {
		pair := subject.Pair{Left: c.left, Right: c.right}
		if got := pair.Format(); got != c.want {
			t.Errorf("Pair.Format(%q,%q) = %q, want %q", c.left, c.right, got, c.want)
		}
		var value subject.Formatter = pair
		if got := subject.Render(value); got != c.want {
			t.Errorf("Render(Pair) = %q, want %q", got, c.want)
		}
	}
}

func TestCallerDefinedFormatters(t *testing.T) {
	for _, want := range []string{"custom", "", " \tvalue:\n"} {
		if got := subject.Render(fixedFormatter(want)); got != want {
			t.Errorf("Render caller formatter = %q, want unchanged %q", got, want)
		}
		value := &countedFormatter{output: want}
		if got := subject.Render(value); got != want {
			t.Errorf("Render pointer formatter = %q, want unchanged %q", got, want)
		}
		if value.calls != 1 {
			t.Errorf("Format calls = %d, want 1", value.calls)
		}
	}
}
