package subject_test

import (
	"task.local/copy-final-element-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	src := []string{"x", "tail"}
	dst := []string{"?"}
	subject.CopyLast(src, dst)
	if dst[0] != "tail" {
		t.Errorf("final got %q want tail", dst[0])
	}
}
