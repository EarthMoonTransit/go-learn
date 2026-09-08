package subject_test

import (
	"task.local/typed-nil-result-001/subject"
	"testing"
)

type otherCoder struct{ code, calls int }

func (c *otherCoder) CodeValue() int { c.calls++; return c.code }

func TestPublicCase01(t *testing.T) {
	for _, v := range []subject.Coder{nil, (*subject.Result)(nil)} {
		n, ok := subject.ReadCode(v)
		if n != 0 || ok {
			t.Errorf("nil %T got%d,%v", v, n, ok)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	for _, code := range []int{0, -3, 9} {
		r := &subject.Result{Code: code}
		n, ok := subject.ReadCode(r)
		if n != code || !ok || r.Code != code {
			t.Errorf("result got%d,%v stored%d want%d,true", n, ok, r.Code, code)
		}
	}
}

func TestPublicCase03(t *testing.T) {
	c := &otherCoder{code: 4}
	n, ok := subject.ReadCode(c)
	if n != 4 || !ok || c.calls != 1 {
		t.Errorf("other got%d,%v calls%d", n, ok, c.calls)
	}
}
