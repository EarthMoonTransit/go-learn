package subject_test

import (
	"task.local/independent-policies-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		id   string
		v    int
		a, b bool
	}{{"alpha", 17, true, false}, {"beta", 18, true, true}, {"other", 18, false, true}, {"", -1, false, false}, {"Alpha", 99, false, true}, {" beta ", 18, false, true}, {"alpha", -5, true, false}} {
		a, b := subject.Classify(c.id, c.v)
		if a != c.a || b != c.b {
			t.Errorf("%q v%d got%v,%v want%v,%v", c.id, c.v, a, b, c.a, c.b)
		}
	}
}
