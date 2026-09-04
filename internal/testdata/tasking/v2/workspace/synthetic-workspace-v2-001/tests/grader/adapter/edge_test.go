package adapter_test

import (
	"task.local/synthetic-workspace-v2-001/adapter"
	"testing"
)

func TestTransformEdges(t *testing.T) {
	for _, value := range []int{0, -4, 12} {
		if got := adapter.Transform(value); got != value {
			t.Errorf("got %d", got)
		}
	}
}
