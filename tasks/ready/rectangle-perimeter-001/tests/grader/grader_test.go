package subject_test

import (
	"math"
	"task.local/rectangle-perimeter-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	got := subject.Perimeter(subject.Rectangle{Width: 1000, Height: 1})
	if math.IsNaN(got) || math.Abs(got-2002) > 1e-9 {
		t.Errorf("large got%g want2002", got)
	}
}
