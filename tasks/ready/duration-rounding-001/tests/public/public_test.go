package subject_test

import (
	"task.local/duration-rounding-001/subject"
	"testing"
	"time"
)

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		v    time.Duration
		want int64
	}{{0, 0}, {time.Minute - time.Nanosecond, 0}, {time.Minute, 1}, {89 * time.Second, 1}, {119 * time.Second, 1}, {2 * time.Minute, 2}, {24 * time.Hour, 1440}} {
		n, e := subject.WholeMinutes(c.v)
		if n != c.want || e != nil {
			t.Errorf("duration %v got%d,%v want%d,nil", c.v, n, e, c.want)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	sentinel := originalSentinel
	if sentinel == nil {
		t.Fatal("nil duration sentinel")
	}
	for _, v := range []time.Duration{-time.Nanosecond, 24*time.Hour + time.Nanosecond, time.Duration(1<<63 - 1)} {
		n, e := subject.WholeMinutes(v)
		if n != 0 || e != sentinel || subject.ErrDuration != sentinel {
			t.Errorf("invalid duration%v got%d,%v want0,original", v, n, e)
		}
	}
}

var originalSentinel = subject.ErrDuration
