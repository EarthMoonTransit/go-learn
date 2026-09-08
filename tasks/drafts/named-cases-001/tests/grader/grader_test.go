package subject_test

import (
	"task.local/named-cases-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	calls := 0
	t.Run("fresh", func(parent *testing.T) {
		subject.RunNamed(parent, []subject.Case{{Name: "again", Check: func(*testing.T) { calls++ }}})
	})
	if calls != 1 {
		t.Errorf("fresh calls%d", calls)
	}
}
