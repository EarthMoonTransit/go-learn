package subject_test

import (
	subject "task.local/once-result-001/subject"
	"testing"
)

func TestPositiveSuccess(t *testing.T) {
	n := 0
	memo := subject.Memoize(func() (int, error) { n++; return 19, nil })
	for range 3 {
		v, e := memo()
		if v != 19 || e != nil {
			t.Errorf("result=(%d,%v); want(19,nil)", v, e)
		}
	}
	if n != 1 {
		t.Errorf("calls=%d; want1", n)
	}
}
