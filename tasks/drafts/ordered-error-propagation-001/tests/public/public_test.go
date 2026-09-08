package subject_test

import (
	"errors"
	"math"
	"slices"
	"task.local/ordered-error-propagation-001/subject"
	"testing"
)

type calculator struct {
	amount, factor       float64
	amountErr, factorErr error
	calls                []string
	count                int
}

func (c *calculator) Amount(n int) (float64, error) {
	c.calls = append(c.calls, "amount")
	c.count = n
	return c.amount, c.amountErr
}
func (c *calculator) Factor() (float64, error) {
	c.calls = append(c.calls, "factor")
	return c.factor, c.factorErr
}

func TestPublicCase01(t *testing.T) {
	c := &calculator{amount: 60.5, factor: 1.3}
	n, e := subject.Divide(c, 5)
	if e != nil || math.IsNaN(n) || math.Abs(n-15.73) > 1e-9 || !slices.Equal(c.calls, []string{"amount", "factor"}) || c.count != 5 {
		t.Errorf("divide got %g,%v calls%v count%d", n, e, c.calls, c.count)
	}
}

func TestPublicCase02(t *testing.T) {
	sentinel := errors.New("amount failed")
	c := &calculator{amount: 99, amountErr: sentinel, factor: 2}
	n, e := subject.Divide(c, 4)
	if e != sentinel || math.IsNaN(n) || math.Abs(n) > 1e-9 || !slices.Equal(c.calls, []string{"amount"}) {
		t.Errorf("first error got %g,%v calls%v", n, e, c.calls)
	}
}

func TestPublicCase03(t *testing.T) {
	sentinel := errors.New("factor failed")
	c := &calculator{amount: 12, factor: 99, factorErr: sentinel}
	n, e := subject.Divide(c, 3)
	if e != sentinel || math.IsNaN(n) || math.Abs(n) > 1e-9 || !slices.Equal(c.calls, []string{"amount", "factor"}) {
		t.Errorf("second error got %g,%v calls%v", n, e, c.calls)
	}
}
