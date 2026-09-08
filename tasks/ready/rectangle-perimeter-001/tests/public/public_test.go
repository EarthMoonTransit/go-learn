package subject_test

import (
	"math"
	"task.local/rectangle-perimeter-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		r    subject.Rectangle
		want float64
	}{{subject.Rectangle{Width: 10, Height: 10}, 40}, {subject.Rectangle{Width: 12, Height: 6}, 36}, {subject.Rectangle{Width: .25, Height: 1.5}, 3.5}, {subject.Rectangle{Width: 3}, 6}, {subject.Rectangle{Height: 4}, 8}, {subject.Rectangle{}, 0}} {
		before := c.r
		got := subject.Perimeter(c.r)
		if math.IsNaN(got) || math.Abs(got-c.want) > 1e-9 || c.r != before {
			t.Errorf("rect%v got%g want%g", c.r, got, c.want)
		}
	}
}
