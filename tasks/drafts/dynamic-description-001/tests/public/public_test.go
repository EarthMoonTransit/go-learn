package subject_test

import (
	"task.local/dynamic-description-001/subject"
	"testing"
)

type numberHolder struct{ n int }

func (n numberHolder) Number() int { return n.n }

type otherFancy struct{}

func (otherFancy) Value() string { panic("Value must not be called") }

type bothBox struct{}

func (bothBox) Number() int   { return 9 }
func (bothBox) Value() string { panic("NumberBox must take precedence") }

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		v    any
		want string
	}{{subject.FancyNumber{Text: "9999999999999999999999999999999999"}, "fancy 0.0"}, {42, "number 42.0"}, {7.11, "number 7.1"}, {-3.2, "number -3.2"}, {numberHolder{16}, "box 16.0"}, {subject.FancyNumber{Text: "16"}, "fancy 16.0"}, {subject.FancyNumber{Text: "bad"}, "fancy 0.0"}, {otherFancy{}, "fancy 0.0"}, {bothBox{}, "box 9.0"}, {nil, "unsupported"}, {"42", "unsupported"}, {int64(42), "unsupported"}} {
		if got := subject.Describe(c.v); got != c.want {
			t.Errorf("Describe %T got %q want %q", c.v, got, c.want)
		}
	}
}
