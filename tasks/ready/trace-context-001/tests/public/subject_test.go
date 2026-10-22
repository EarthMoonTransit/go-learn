package subject_test

import (
	"context"
	subject "task.local/trace-context-001/subject"
	"testing"
	"time"
)

func checkTrace(t *testing.T, ctx context.Context, want string, present bool) {
	t.Helper()
	got, ok := subject.Trace(ctx)
	if got != want || ok != present {
		t.Errorf("Trace=(%q,%t); want(%q,%t)", got, ok, want, present)
	}
}
func TestAbsentEmptyAndShadow(t *testing.T) {
	parent := context.Background()
	checkTrace(t, parent, "", false)
	empty := subject.WithTrace(parent, "")
	checkTrace(t, empty, "", true)
	a := subject.WithTrace(parent, "A")
	b := subject.WithTrace(a, "B")
	checkTrace(t, a, "A", true)
	checkTrace(t, b, "B", true)
	checkTrace(t, parent, "", false)
}

type foreignKey struct{}
type wrongValue struct{ context.Context }

func (wrongValue) Value(any) any { return 23 }
func TestUnrelatedAndWrongDynamicValues(t *testing.T) {
	for _, ctx := range []context.Context{context.WithValue(context.Background(), "trace", "other"), context.WithValue(context.Background(), foreignKey{}, "foreign"), wrongValue{context.Background()}} {
		checkTrace(t, ctx, "", false)
	}
	parent := context.WithValue(context.Background(), foreignKey{}, "keep")
	child := subject.WithTrace(parent, "id")
	if got := child.Value(foreignKey{}); got != "keep" {
		t.Errorf("unrelated parent value=%v; want keep", got)
	}
}
func TestParentCancellationAndDeadline(t *testing.T) {
	deadline := time.Date(2099, 1, 2, 3, 4, 5, 0, time.UTC)
	parent, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	child := subject.WithTrace(parent, "request-7")
	got, ok := child.Deadline()
	if !ok || !got.Equal(deadline) {
		t.Errorf("deadline=(%v,%t); want inherited %v", got, ok, deadline)
	}
	cancel()
	select {
	case <-child.Done():
	default:
		t.Error("child did not preserve parent cancellation")
	}
	if child.Err() != context.Canceled {
		t.Errorf("child Err=%v; want context.Canceled", child.Err())
	}
	checkTrace(t, child, "request-7", true)
}

func TestLiteralMetadata(t *testing.T) {
	checkTrace(t, subject.WithTrace(context.Background(), " zero\nΩ "), " zero\nΩ ", true)
}
