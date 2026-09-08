package subject_test

import (
	"slices"
	"task.local/initialization-values-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	a, b := subject.Values()
	if a != 3 || b != 6 {
		t.Errorf("values got%d,%d want3,6", a, b)
	}
	if got := subject.Trace(); !slices.Equal(got, []string{"base", "double", "ready"}) {
		t.Errorf("init order got%v want[base double ready]", got)
	}
}

func TestPublicCase02(t *testing.T) {
	subject.Values()
	subject.Values()
	got := subject.Trace()
	if !slices.Equal(got, []string{"base", "double", "ready"}) {
		t.Errorf("repeated access changed init trace%v", got)
	}
	got[0] = "changed"
	if !slices.Equal(subject.Trace(), []string{"base", "double", "ready"}) {
		t.Error("trace accessor shares storage")
	}
}
