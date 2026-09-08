package subject_test

import (
	"errors"
	"fmt"
	"task.local/error-detail-tests-001/subject"
	"testing"
)

type projectedError struct{ detail *subject.Detail }

func (e *projectedError) Error() string { return "projection" }
func (e *projectedError) As(target any) bool {
	p, ok := target.(**subject.Detail)
	if ok {
		*p = e.detail
	}
	return ok
}

func TestContractCase01(t *testing.T) {
	d := &subject.Detail{Code: 7, Message: "bad"}
	got, ok := subject.FindDetail(d)
	if got != d || !ok {
		t.Errorf("direct got%v,%v", got, ok)
	}
}

func TestContractCase02(t *testing.T) {
	a := &subject.Detail{Code: 1, Message: "a"}
	b := &subject.Detail{Code: 2, Message: "b"}
	wrapped := fmt.Errorf("outer: %w", a)
	for _, e := range []error{wrapped, errors.Join(errors.New("other"), wrapped, b)} {
		got, ok := subject.FindDetail(e)
		if got != a || !ok {
			t.Errorf("first cause got%v,%v", got, ok)
		}
	}
}

func TestContractCase03(t *testing.T) {
	for _, e := range []error{nil, errors.New("bad")} {
		d, ok := subject.FindDetail(e)
		if d != nil || ok {
			t.Errorf("miss got%v,%v", d, ok)
		}
	}
	var typed *subject.Detail
	d, ok := subject.FindDetail(typed)
	if d != nil || !ok {
		t.Errorf("typed nil got%v,%v", d, ok)
	}
}

func TestContractCase04(t *testing.T) {
	want := &subject.Detail{Code: 9, Message: "custom"}
	got, ok := subject.FindDetail(&projectedError{detail: want})
	if got != want || !ok {
		t.Errorf("custom As got%v,%v", got, ok)
	}
}
