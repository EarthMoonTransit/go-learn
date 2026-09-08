package subject_test

import (
	"task.local/typed-value-observations-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	a, b, c := subject.ZeroValues()
	s, okS := a.(string)
	n, okN := b.(int)
	v, okV := c.(bool)
	if !okS || !okN || !okV || s != "" || n != 0 || v {
		t.Errorf("zero got%T:%v %T:%v %T:%v", a, a, b, b, c, c)
	}
}

func TestPublicCase02(t *testing.T) {
	a, b, c, d := subject.LiteralValues()
	s, okS := a.(string)
	n, okN := b.(int)
	v, okV := c.(bool)
	if !okS || !okN || !okV || s != "go" || n != 42 || !v || d != float32(3.14) {
		t.Errorf("literal got%T:%v %T:%v %T:%v float%g", a, a, b, b, c, c, d)
	}
}
