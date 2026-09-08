package subject_test

import (
	"task.local/generic-stack-001/subject"
	"testing"
)

func TestGraderBoundary(t *testing.T) {
	var s subject.Stack[bool]
	s.Push(false)
	got, ok := s.Pop()
	if got || !ok || !s.IsEmpty() {
		t.Errorf("stored zero got%v,%v", got, ok)
	}
}
