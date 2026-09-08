package subject_test

import (
	"task.local/afternoon-window-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		text string
		want bool
	}{{"Thursday, May 13, 2010 20:32:00", false}, {"Friday, March 8, 1974 12:02:02", true}, {"Friday, September 9, 2112 11:59:59", false}, {"Monday, January 1, 2024 12:00:00", true}, {"Monday, January 1, 2024 17:59:59.999", true}, {"Monday, January 1, 2024 18:00:00", false}} {
		got, e := subject.IsAfternoon(c.text)
		if got != c.want || e != nil {
			t.Errorf("window%q got%v,%v want%v,nil", c.text, got, e, c.want)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	sentinel := originalSentinel
	if sentinel == nil {
		t.Fatal("nil date sentinel")
	}
	for _, text := range []string{"", "Monday, February 30, 2024 13:00:00", "Monday, January 1, 2024 25:00:00", "1/1/2024 13:00:00"} {
		v, e := subject.IsAfternoon(text)
		if v || e != sentinel || subject.ErrDate != sentinel {
			t.Errorf("invalid %q got%v,%v", text, v, e)
		}
	}
}

var originalSentinel = subject.ErrDate
