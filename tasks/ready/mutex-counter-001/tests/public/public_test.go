package subject_test

import (
	"sync"
	"task.local/mutex-counter-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	c := subject.NewCounter()
	if c == nil || c.Value() != 0 {
		t.Fatal("new counter must start at zero")
	}
	c.Inc()
	c.Inc()
	c.Inc()
	if c.Value() != 3 {
		t.Errorf("three increments got%d want3", c.Value())
	}
}

func TestPublicCase02(t *testing.T) {
	c := subject.NewCounter()
	var wg sync.WaitGroup
	wg.Add(1000)
	for range 1000 {
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	if c.Value() != 1000 {
		t.Errorf("concurrent increments got%d want1000", c.Value())
	}
}

func TestPublicCase03(t *testing.T) {
	var a, b subject.Counter
	a.Inc()
	b.Inc()
	b.Inc()
	if a.Value() != 1 || b.Value() != 2 {
		t.Errorf("independent zero counters got%d,%d", a.Value(), b.Value())
	}
}
