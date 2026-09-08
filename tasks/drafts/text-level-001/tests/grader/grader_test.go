package subject_test

import (
	"task.local/text-level-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	var l subject.Level
	if e := l.UnmarshalText([]byte("0")); e != nil || l != 0 {
		t.Errorf("zero got%d,%v", l, e)
	}
}
