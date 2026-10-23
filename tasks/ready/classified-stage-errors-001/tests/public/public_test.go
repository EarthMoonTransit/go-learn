package subject_test

import (
	"errors"
	"fmt"
	"reflect"
	"task.local/classified-stage-errors-001/subject"
	"testing"
)

type matchingCause struct{ target error }

func (m *matchingCause) Error() string        { return "custom match" }
func (m *matchingCause) Is(target error) bool { return target == m.target }

type projectedCause struct{ detail *subject.Detail }

func (p *projectedCause) Error() string { return "projected detail" }
func (p *projectedCause) As(target any) bool {
	slot, ok := target.(**subject.Detail)
	if ok {
		*slot = p.detail
	}
	return ok
}

type renderedCause struct{}

func (*renderedCause) Error() string { return "actual Error text" }
func (*renderedCause) Format(state fmt.State, verb rune) {
	fmt.Fprint(state, "different formatter text")
}

func checkFailure(t *testing.T, out subject.Outcome, value int, name string, cause error, expose, temp bool, detail *subject.Detail, found bool) {
	t.Helper()
	if out.Value != value {
		t.Errorf("failure value=%d; want %d", out.Value, value)
	}
	if out.Err == nil {
		t.Error("failure Err=nil; want contextual error")
		return
	}
	wantText := name + ": " + cause.Error()
	if out.Err.Error() != wantText {
		t.Errorf("context text=%q; want %q", out.Err.Error(), wantText)
	}
	wantCause := error(nil)
	if expose {
		wantCause = cause
	}
	if got := errors.Unwrap(out.Err); got != wantCause {
		t.Errorf("direct cause=%T %v; want original=%t (%T %v)", got, got, expose, wantCause, wantCause)
	}
	if !expose && errors.Is(out.Err, cause) {
		t.Error("opaque context exposes original cause")
	}
	if out.Temporary != temp {
		t.Errorf("Temporary=%t; want %t for returned error", out.Temporary, temp)
	}
	if out.HasDetail != found || out.Detail != detail {
		t.Errorf("detail=(%p,%t); want (%p,%t)", out.Detail, out.HasDetail, detail, found)
	}
}
func TestStageOrderAndValues(t *testing.T) {
	for _, fail := range []int{-1, 0, 1, 2} {
		t.Run(fmt.Sprint(fail), func(t *testing.T) {
			cause := errors.New("stop")
			var calls []int
			var stages []subject.Stage
			for i := 0; i < 3; i++ {
				stages = append(stages, subject.Stage{Name: fmt.Sprint(i), Expose: true, Run: func() (int, error) {
					calls = append(calls, i)
					if i == fail {
						return -10 - i, cause
					}
					return 20 + i, nil
				}})
			}
			out := subject.RunStages(stages)
			wantCalls := []int{0, 1, 2}
			if fail >= 0 {
				wantCalls = wantCalls[:fail+1]
			}
			if !reflect.DeepEqual(calls, wantCalls) {
				t.Errorf("stage calls=%v; want %v exactly once", calls, wantCalls)
			}
			if fail >= 0 {
				checkFailure(t, out, -10-fail, fmt.Sprint(fail), cause, true, false, nil, false)
			} else if out.Value != 22 || out.Err != nil || out.Temporary || out.Detail != nil || out.HasDetail {
				t.Errorf("success Outcome=%+v; want value22 and empty error classification", out)
			}
		})
	}
	for _, stages := range [][]subject.Stage{nil, {}} {
		out := subject.RunStages(stages)
		if out != (subject.Outcome{}) {
			t.Errorf("empty Outcome=%+v; want zero Outcome", out)
		}
	}
}
func TestContextTextAndIdentity(t *testing.T) {
	base := errors.New("base")
	nested := fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", base))
	for _, name := range []string{"", "  decode λ  ", "load"} {
		for _, cause := range []error{nested, &renderedCause{}} {
			out := subject.RunStages([]subject.Stage{{Name: name, Expose: true, Run: func() (int, error) { return 7, cause }}})
			checkFailure(t, out, 7, name, cause, true, false, nil, false)
		}
	}
}
func TestTemporaryClassification(t *testing.T) {
	cases := []struct {
		name  string
		cause error
		want  bool
	}{
		{"sentinel", subject.ErrTemporary, true}, {"wrapped", fmt.Errorf("inside: %w", subject.ErrTemporary), true},
		{"same-message", errors.New(subject.ErrTemporary.Error()), false}, {"custom", &matchingCause{subject.ErrTemporary}, true},
		{"custom-miss", &matchingCause{errors.New("another")}, false}, {"later-branch", errors.Join(errors.New("first"), subject.ErrTemporary), true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := subject.RunStages([]subject.Stage{{Name: "step", Expose: true, Run: func() (int, error) { return 3, c.cause }}})
			checkFailure(t, out, 3, "step", c.cause, true, c.want, nil, false)
		})
	}
}
func TestDetailClassification(t *testing.T) {
	left, right := &subject.Detail{Code: 2, Message: "left"}, &subject.Detail{Code: 8, Message: "right"}
	var nilDetail *subject.Detail
	cases := []struct {
		name  string
		cause error
		want  *subject.Detail
		found bool
	}{
		{"direct", left, left, true}, {"wrapped", fmt.Errorf("inside: %w", left), left, true},
		{"joined-first", errors.Join(fmt.Errorf("wrapped: %w", left), right), left, true},
		{"custom", &projectedCause{right}, right, true}, {"typed-nil", nilDetail, nil, true},
		{"custom-nil", &projectedCause{nil}, nil, true}, {"nil-first", errors.Join(nilDetail, right), nil, true},
		{"unrelated", errors.New("ordinary"), nil, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := subject.RunStages([]subject.Stage{{Name: "read", Expose: true, Run: func() (int, error) { return -2, c.cause }}})
			checkFailure(t, out, -2, "read", c.cause, true, false, c.want, c.found)
		})
	}
}
func TestOpaqueBoundary(t *testing.T) {
	detail := &subject.Detail{Code: 5, Message: "structured"}
	var nilDetail *subject.Detail
	causes := []error{errors.Join(subject.ErrTemporary, detail), &matchingCause{subject.ErrTemporary}, &projectedCause{detail}, nilDetail}
	for i, cause := range causes {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			out := subject.RunStages([]subject.Stage{{Name: " public ", Expose: false, Run: func() (int, error) { return 9, cause }}})
			checkFailure(t, out, 9, " public ", cause, false, false, nil, false)
		})
	}
}
