package subject_test

import (
	"errors"
	"task.local/embedded-cache-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
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

func TestContractCase02(t *testing.T) {
	f := &subject.Feed{Total: 42, Lookup: func(k string) (subject.Document, error) { return subject.Document{Key: "different", Title: k}, nil }}
	c := subject.NewCache(f)
	if c == nil {
		t.Fatal("nil cache")
	}
	for _, k := range []string{"a", "a", "A", "A"} {
		d, e := c.Fetch(k)
		if e != nil || d.Title != k || d.Key != "different" {
			t.Errorf("key%q got%v,%v", k, d, e)
		}
	}
	if f.Calls != 2 {
		t.Errorf("lookup calls got%d want2", f.Calls)
	}
	if c.Count() != 42 {
		t.Errorf("promoted count got%d want42", c.Count())
	}
	f.Total = 9
	if c.Count() != 9 {
		t.Errorf("live count got%d want9", c.Count())
	}
	other := subject.NewCache(f)
	_, _ = other.Fetch("a")
	if f.Calls != 3 {
		t.Errorf("independent cache calls%d want3", f.Calls)
	}
}

func TestContractCase03(t *testing.T) {
	sentinel := errors.New("missing")
	attempts := 0
	f := &subject.Feed{Lookup: func(string) (subject.Document, error) {
		attempts++
		if attempts < 3 {
			return subject.Document{Key: "ignored"}, sentinel
		}
		return subject.Document{}, nil
	}}
	c := subject.NewCache(f)
	if c == nil {
		t.Fatal("nil cache")
	}
	for i := 0; i < 2; i++ {
		d, e := c.Fetch("")
		if d != (subject.Document{}) || e != sentinel {
			t.Errorf("failure got%v,%v wantzero,original", d, e)
		}
	}
	for i := 0; i < 2; i++ {
		d, e := c.Fetch("")
		if d != (subject.Document{}) || e != nil {
			t.Errorf("empty success got%v,%v", d, e)
		}
	}
	if attempts != 3 || f.Calls != 3 {
		t.Errorf("retry/empty cache attempts%d calls%d want3", attempts, f.Calls)
	}
}
