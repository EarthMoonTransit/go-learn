package subject_test

import (
	"errors"
	"fmt"
	"task.local/error-chain-match-001/subject"
	"testing"
)

type semanticError struct{ target error }

func (e *semanticError) Error() string        { return "semantic" }
func (e *semanticError) Is(target error) bool { return target == e.target }

func TestPublicCase01(t *testing.T) {
	target := errors.New("x")
	other := errors.New("x")
	wrapped := fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", target))
	for _, c := range []struct {
		err, target error
		want        bool
	}{{nil, nil, true}, {target, nil, false}, {nil, target, false}, {target, target, true}, {other, target, false}, {wrapped, target, true}, {errors.Join(other, wrapped), target, true}, {errors.Join(other), target, false}} {
		if got := subject.Matches(c.err, c.target); got != c.want {
			t.Errorf("match %v/%v got%v want%v", c.err, c.target, got, c.want)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	target := errors.New("marker")
	err := &semanticError{target: target}
	if !subject.Matches(err, target) || subject.Matches(err, errors.New("marker")) {
		t.Error("custom Is semantics lost")
	}
}
