package subject_test

import "fmt"

import (
	"errors"
	"task.local/error-context-wrap-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	cause := errors.New("failed")
	wrapped := subject.Wrap("read", cause)
	if wrapped == nil || wrapped.Error() != "read: failed" || errors.Unwrap(wrapped) != cause {
		t.Errorf("wrapped got%v unwrap%v", wrapped, errors.Unwrap(wrapped))
	}
}

func TestPublicCase02(t *testing.T) {
	cause := errors.New("bad")
	first := subject.Wrap("", cause)
	second := subject.Wrap(" outer ", first)
	if first == nil || second == nil || first.Error() != ": bad" || second.Error() != " outer : : bad" || errors.Unwrap(first) != cause || errors.Unwrap(second) != first || cause.Error() != "bad" {
		t.Errorf("layers got%v,%v", first, second)
	}
}

func TestPublicCase03(t *testing.T) {
	for _, op := range []string{"", "read"} {
		if got := subject.Wrap(op, nil); got != nil {
			t.Errorf("nil cause got%v", got)
		}
	}
}

type formattedCause struct{}

func (*formattedCause) Error() string                 { return "error text" }
func (*formattedCause) Format(s fmt.State, verb rune) { fmt.Fprint(s, "different formatter text") }
func TestUsesErrorText(t *testing.T) {
	cause := &formattedCause{}
	got := subject.Wrap("read", cause)
	if got == nil || got.Error() != "read: error text" || errors.Unwrap(got) != cause {
		t.Errorf("Error text: got %v, want read: error text with original cause", got)
	}
}
