package subject_test

import (
	"errors"
	"testing"

	"task.local/error-chain-query-001/subject"
)

func TestWrapFieldPreservesCause(t *testing.T) {
	cause := errors.New("denied")
	if err := subject.WrapField("mode", cause); !errors.Is(err, cause) {
		t.Fatal("wrapped error does not preserve its cause")
	}
}
