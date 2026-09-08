package subject_test

import (
	"task.local/afternoon-window-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	v, e := subject.IsAfternoon("Monday, January 1, 2024 00:00:00")
	if v || e != nil {
		t.Errorf("midnight got%v,%v", v, e)
	}
}
