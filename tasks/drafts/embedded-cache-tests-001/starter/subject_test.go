package subject_test

import (
	"task.local/embedded-cache-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	f := &subject.Feed{Total: 42, Lookup: func(k string) (subject.Document, error) {
		return subject.Document{Key: k, Title: "Title for " + k}, nil
	}}
	c := subject.NewCache(f)
	if c == nil {
		t.Fatal("nil cache")
	}
	d, e := c.Fetch("a")
	if e != nil || d != (subject.Document{Key: "a", Title: "Title for a"}) || f.Calls != 1 {
		t.Errorf("fetch got%v,%v calls%d", d, e, f.Calls)
	}
}

// Add cases for the remaining requirements in TASK.md.
