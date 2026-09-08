package subject_test

import (
	"task.local/parse-timestamp-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	v, e := subject.ParseTimestamp("1/1/2000 00:00:00")
	if e != nil || v.Year() != 2000 || v.Hour() != 0 {
		t.Errorf("midnight got%v,%v", v, e)
	}
}
