package subject_test

import (
	"task.local/exact-dynamic-number-001/subject"
	"testing"
)

type alternateNumber struct{}

func (alternateNumber) Value() string { panic("unsupported Value must not be called") }

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		text string
		want int
	}{{"11", 11}, {"0", 0}, {"-12", -12}, {"+007", 7}, {"two", 0}, {" 10", 0}, {"1.5", 0}, {"999999999999999999999999999999999999", 0}} {
		if got := subject.Extract(subject.FancyNumber{Text: c.text}); got != c.want {
			t.Errorf("Extract %q got %d want %d", c.text, got, c.want)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	for _, v := range []subject.FancyNumberBox{nil, &subject.FancyNumber{Text: "4"}, alternateNumber{}, (*subject.FancyNumber)(nil)} {
		if got := subject.Extract(v); got != 0 {
			t.Errorf("unsupported %T got %d want0", v, got)
		}
	}
}
