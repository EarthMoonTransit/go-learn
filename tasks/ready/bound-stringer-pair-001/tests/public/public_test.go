package subject_test

import (
	"fmt"
	"task.local/bound-stringer-pair-001/subject"
	"testing"
)

func formats(t *testing.T, value fmt.Stringer, want string) {
	t.Helper()
	for _, got := range []string{value.String(), fmt.Sprint(value), fmt.Sprintf("%s", value)} {
		if got != want {
			t.Errorf("formatted text=%q; want %q", got, want)
		}
	}
}
func binding(t *testing.T, text string, initial int) {
	t.Helper()
	label := subject.Label{Text: text}
	counter := &subject.Counter{Value: initial}
	left, right := subject.Bind(label, counter)
	stored, ok := left.(subject.Label)
	if !ok {
		t.Errorf("snapshot dynamic type=%T; want subject.Label", left)
		return
	}
	if stored != label {
		t.Errorf("snapshot value=%+v; want %+v", stored, label)
	}
	live, ok := right.(*subject.Counter)
	if !ok || live == nil {
		t.Errorf("live binding=%T; want nonnil *Counter", right)
		return
	}
	if live != counter {
		t.Errorf("live pointer=%p; want original %p", live, counter)
	}
	if counter.Value != initial {
		t.Errorf("Bind changed counter=%d; want %d", counter.Value, initial)
	}
	formats(t, left, "label="+text)
	formats(t, right, fmt.Sprintf("count=%d", initial))
	if counter.Value != initial {
		t.Errorf("formatting changed counter=%d; want %d", counter.Value, initial)
	}
	if label.Text != text {
		t.Errorf("formatting changed label=%q; want %q", label.Text, text)
	}
	label.Text = "changed after binding"
	counter.Value = -initial
	formats(t, left, "label="+text)
	formats(t, right, fmt.Sprintf("count=%d", -initial))
	if counter.Value != -initial {
		t.Errorf("repeated formatting changed counter=%d; want %d", counter.Value, -initial)
	}
	formats(t, left, "label="+text)
}
func TestSnapshotAndLiveBindings(t *testing.T) {
	for _, c := range []struct {
		text  string
		value int
	}{{"  λ  ", -7}, {"", 0}, {"count", 12}} {
		t.Run(fmt.Sprint(c.value), func(t *testing.T) { binding(t, c.text, c.value) })
	}
}
func TestAbsentCounterIsNilInterface(t *testing.T) {
	for _, text := range []string{"", " pad "} {
		left, right := subject.Bind(subject.Label{Text: text}, nil)
		if right != nil {
			t.Errorf("absent counter returned nonnil interface with dynamic type %T", right)
		}
		snapshot, ok := left.(subject.Label)
		if !ok {
			t.Errorf("snapshot dynamic type=%T; want subject.Label", left)
			continue
		}
		if snapshot.Text != text {
			t.Errorf("snapshot text=%q; want %q", snapshot.Text, text)
		}
		formats(t, left, "label="+text)
	}
}
func TestDirectStableStrings(t *testing.T) {
	for _, text := range []string{"", " padded ", "雪 λ"} {
		label := subject.Label{Text: text}
		formats(t, label, "label="+text)
		if label.Text != text {
			t.Error("String changed label")
		}
	}
	for _, value := range []int{-12, 0, 45} {
		counter := &subject.Counter{Value: value}
		formats(t, counter, fmt.Sprintf("count=%d", value))
		if counter.Value != value {
			t.Errorf("formatting changed counter=%d; want %d", counter.Value, value)
		}
	}
}
func TestSeparateBindingsStayIndependent(t *testing.T) {
	first, second := &subject.Counter{Value: 3}, &subject.Counter{Value: 8}
	a, x := subject.Bind(subject.Label{Text: "first"}, first)
	b, y := subject.Bind(subject.Label{Text: "second"}, second)
	ax, okx := a.(subject.Label)
	by, oky := b.(subject.Label)
	px, okp := x.(*subject.Counter)
	py, okq := y.(*subject.Counter)
	if !okx || !oky || !okp || !okq || px == nil || py == nil {
		t.Error("independent bindings have wrong dynamic types or nil live values")
		return
	}
	if ax.Text != "first" || by.Text != "second" || px != first || py != second {
		t.Error("independent bindings lost snapshots or original pointers")
	}
	first.Value = -4
	formats(t, x, "count=-4")
	formats(t, y, "count=8")
	formats(t, a, "label=first")
	formats(t, b, "label=second")
}
