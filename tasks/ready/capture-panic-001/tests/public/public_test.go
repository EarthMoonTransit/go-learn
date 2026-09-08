package subject_test

import (
	"errors"
	"fmt"
	"task.local/capture-panic-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	calls := 0
	sentinel := errors.New("work failed")
	for _, cause := range []error{nil, sentinel} {
		err, p := subject.Capture(func() error { calls++; return cause })
		if err != cause || p {
			t.Errorf("normal got%v,%v want%v,false", err, p, cause)
		}
	}
	if calls != 2 {
		t.Errorf("normal calls%d want2", calls)
	}
}

func TestPublicCase02(t *testing.T) {
	for _, v := range []any{"broken", "", 7, 0} {
		calls := 0
		err, p := subject.Capture(func() error { calls++; panic(v) })
		want := fmt.Sprintf("panic: %v", v)
		if err == nil || err.Error() != want || !p || calls != 1 {
			t.Errorf("panic got%v,%v calls%d want%q,true,1", err, p, calls, want)
		}
	}
}
