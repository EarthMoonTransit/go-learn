package subject_test

import (
	"sync"
	"task.local/concurrent-counter-read-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	var c subject.Counter
	if c.Value() != 0 {
		t.Error("zero read")
	}
	c.Inc()
	c.Inc()
	c.Inc()
	if c.Value() != 3 || c.Value() != 3 {
		t.Error("stable read after three increments")
	}
}

func TestPublicCase02(t *testing.T) {
	c := subject.NewCounter()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range 1000 {
			c.Inc()
		}
	}()
	go func() {
		defer wg.Done()
		for range 1000 {
			if v := c.Value(); v < 0 || v > 1000 {
				t.Errorf("concurrent read %d outside0..1000", v)
			}
		}
	}()
	wg.Wait()
	if c.Value() != 1000 || c.Value() != 1000 {
		t.Error("final count must be stable1000")
	}
}

func TestPublicCase03(t *testing.T) {
	a, b := subject.NewCounter(), subject.NewCounter()
	a.Inc()
	if a.Value() != 1 || b.Value() != 0 {
		t.Error("counter state is shared")
	}
}
