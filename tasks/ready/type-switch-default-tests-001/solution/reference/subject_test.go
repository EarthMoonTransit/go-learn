package subject_test

import (
	"task.local/type-switch-default-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if got := subject.TypeLabel(3); got != "int" {
		t.Errorf("int got%q", got)
	}
}

func TestContractCase02(t *testing.T) {
	type alias = int
	for _, c := range []struct {
		v    any
		want string
	}{{nil, "nil"}, {0, "int"}, {alias(9), "int"}, {"", "string"}, {"int", "string"}, {subject.NamedInt(3), "other"}, {int64(3), "other"}, {3.0, "other"}, {true, "other"}, {(*int)(nil), "other"}, {[]int(nil), "other"}, {struct{}{}, "other"}} {
		if got := subject.TypeLabel(c.v); got != c.want {
			t.Errorf("dynamic %T:%v got%q want%q", c.v, c.v, got, c.want)
		}
	}
}
