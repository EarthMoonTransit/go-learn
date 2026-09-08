package subject_test

import (
	"slices"
	"task.local/deferred-release-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	events := []string{}
	n := subject.WithRelease(func() int { events = append(events, "work"); return 7 }, func() { events = append(events, "release") })
	if n != 7 || !slices.Equal(events, []string{"work", "release"}) {
		t.Errorf("normal got%d events%v", n, events)
	}
}

func TestPublicCase02(t *testing.T) {
	token := &struct{ n int }{3}
	events := []string{}
	func() {
		defer func() {
			if got := recover(); got != token {
				t.Errorf("panic got%v want original", got)
			}
		}()
		subject.WithRelease(func() int { events = append(events, "work"); panic(token) }, func() { events = append(events, "release") })
	}()
	if !slices.Equal(events, []string{"work", "release"}) {
		t.Errorf("panic events%v want work,release", events)
	}
}
