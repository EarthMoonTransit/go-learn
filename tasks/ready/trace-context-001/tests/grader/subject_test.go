package subject_test

import (
	"context"
	subject "task.local/trace-context-001/subject"
	"testing"
)

func TestRepeatedShadow(t *testing.T) {
	a := subject.WithTrace(context.Background(), "original")
	b := subject.WithTrace(subject.WithTrace(a, "middle"), "")
	checkTrace(t, b, "", true)
	checkTrace(t, a, "original", true)
}
