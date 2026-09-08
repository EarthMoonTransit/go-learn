package subject_test

import (
	"strings"
	"task.local/single-json-value-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	v, e := subject.DecodeOne(strings.NewReader("{\"count\":0}"))
	if v != (subject.Record{}) || e != nil {
		t.Errorf("zero count got%v,%v", v, e)
	}
}
