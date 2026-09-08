package subject_test

import (
	"context"
	"task.local/cancel-worker-001/subject"
	"testing"
	"testing/synctest"
	"time"
)

func TestSiblingLifetimesAndDeadline(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		deadline := time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC)
		parent, stop := context.WithDeadline(context.Background(), deadline)
		defer stop()
		a, cancelA, doneA := subject.StartCancelable(parent)
		b, cancelB, doneB := subject.StartCancelable(parent)
		if a == nil || b == nil || cancelA == nil || cancelB == nil || doneA == nil || doneB == nil {
			t.Fatal("nil sibling result")
		}
		defer cancelA()
		defer cancelB()
		if d, ok := a.Deadline(); !ok || !d.Equal(deadline) {
			t.Error("parent deadline lost")
		}
		cancelA()
		awaitCanceled(t, doneA)
		if b.Err() != nil || parent.Err() != nil {
			t.Error("child cancellation escaped its scope")
		}
		select {
		case <-doneB:
			t.Error("sibling worker finished early")
		default:
		}
		cancelB()
		awaitCanceled(t, doneB)
	})
}
