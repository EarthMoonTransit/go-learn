package subject_test

import (
	"math"
	"task.local/shape-contract-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	var shape subject.Shape = subject.Rectangle{Width: 12, Height: 6}
	got := subject.AreaOf(shape)
	if math.IsNaN(got) || math.Abs(got-72) > 1e-9 {
		t.Errorf("rectangle got %g want72", got)
	}
}

func TestContractCase02(t *testing.T) {
	cases := []struct {
		s    subject.Shape
		want float64
	}{{subject.Circle{Radius: 10}, math.Pi * 100}, {subject.Triangle{Base: 12, Height: 6}, 36}, {subject.Rectangle{Width: .5, Height: 3}, 1.5}, {subject.Circle{Radius: .5}, math.Pi / 4}, {subject.Triangle{Base: 1, Height: 1}, .5}, {subject.Rectangle{}, 0}, {subject.Circle{}, 0}, {subject.Triangle{Base: 4}, 0}}
	for i, c := range cases {
		got := subject.AreaOf(c.s)
		if math.IsNaN(got) || math.Abs(got-c.want) > 1e-9 {
			t.Errorf("shape%d got %g want %g", i, got, c.want)
		}
	}
}

type customShape struct{ value float64 }

func (s customShape) Area() float64 { return s.value }
func TestCustomAreaContract(t *testing.T) {
	for _, v := range []float64{37, .25, 0} {
		got := subject.AreaOf(customShape{value: v})
		if math.IsNaN(got) || math.Abs(got-v) > 1e-9 {
			t.Errorf("custom dispatch got %g want %g", got, v)
		}
	}
}
