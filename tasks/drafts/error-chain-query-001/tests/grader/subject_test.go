package subject_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"task.local/error-chain-query-001/subject"
)

func TestErrorChainQueries(t *testing.T) {
	cause := errors.New("denied")
	wrapped := subject.WrapField("mode", cause)
	outer := fmt.Errorf("configure: %w", wrapped)

	if !subject.HasCause(outer, cause) {
		t.Error("HasCause missed wrapped target")
	}
	lookalike := errors.New("denied")
	if subject.HasCause(outer, lookalike) {
		t.Error("HasCause matched message instead of identity")
	}
	fieldErr, ok := subject.FindField(outer)
	if !ok || fieldErr.Field != "mode" || !errors.Is(fieldErr.Err, cause) {
		t.Fatalf("FindField returned %#v, %t", fieldErr, ok)
	}
	if !strings.Contains(fieldErr.Error(), "mode") || !strings.Contains(fieldErr.Error(), "denied") {
		t.Fatalf("Error omitted context: %q", fieldErr.Error())
	}
}

func TestWrapFieldNil(t *testing.T) {
	if err := subject.WrapField("mode", nil); err != nil {
		t.Fatalf("WrapField returned %v; want nil", err)
	}
}
