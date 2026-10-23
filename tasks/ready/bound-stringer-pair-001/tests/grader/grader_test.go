package subject_test

import (
	"fmt"
	"task.local/bound-stringer-pair-001/subject"
	"testing"
)

func TestFiniteBindingExtensions(t *testing.T) {
	for _, value := range []int{-1000, -1, 1, 1000} {
		for _, text := range []string{"a b", "界", "  ", "01234567890123456789012345678901"} {
			t.Run(fmt.Sprintf("%d/%q", value, text), func(t *testing.T) { binding(t, text, value) })
		}
	}
	counter := &subject.Counter{Value: 2}
	oldLabel, oldLive := subject.Bind(subject.Label{Text: "old"}, counter)
	counter.Value = 9
	newLabel, newLive := subject.Bind(subject.Label{Text: "new"}, counter)
	a, oka := oldLabel.(subject.Label)
	b, okb := newLabel.(subject.Label)
	p, okp := oldLive.(*subject.Counter)
	q, okq := newLive.(*subject.Counter)
	if !oka || !okb || !okp || !okq || p == nil || q == nil {
		t.Error("rebinding dynamic types or pointers wrong")
		return
	}
	if a.Text != "old" || b.Text != "new" || p != counter || q != counter {
		t.Error("rebinding replaced stored values or pointer identity")
	}
	formats(t, oldLive, "count=9")
	formats(t, newLive, "count=9")
}
