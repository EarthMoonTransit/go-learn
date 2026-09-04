package adapter_test

import (
	"task.local/synthetic-workspace-v2-001/adapter"
	"testing"
)

func TestTransformPublic(t *testing.T) {
	if got := adapter.Transform(7); got != 7 {
		t.Fatalf("got %d", got)
	}
}
