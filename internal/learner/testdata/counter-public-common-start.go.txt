package subject_test

import (
	"runtime"
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

func TestConcurrentReadCommonStart(t *testing.T) {
	previousProcs := runtime.GOMAXPROCS(0)
	if previousProcs < 2 {
		runtime.GOMAXPROCS(2)
		t.Cleanup(func() { runtime.GOMAXPROCS(previousProcs) })
	}

	const writers, readers, calls = 4, 4, 1000
	const finalCount = writers * calls
	c := subject.NewCounter()
	ready := make(chan struct{}, writers+readers)
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(writers + readers)
	type invalidRead struct {
		found bool
		value int
	}
	var invalid [readers]invalidRead
	for range writers {
		go func() {
			defer wg.Done()
			ready <- struct{}{}
			<-start
			for range calls {
				c.Inc()
				runtime.Gosched()
			}
		}()
	}
	for reader := range readers {
		go func(reader int) {
			defer wg.Done()
			ready <- struct{}{}
			<-start
			for range calls {
				value := c.Value()
				if (value < 0 || value > finalCount) && !invalid[reader].found {
					invalid[reader] = invalidRead{found: true, value: value}
				}
				runtime.Gosched()
			}
		}(reader)
	}
	for range writers + readers {
		<-ready
	}
	close(start)
	wg.Wait()
	for reader, observed := range invalid {
		if observed.found {
			t.Errorf("reader %d observed %d outside 0..%d after common start", reader, observed.value, finalCount)
		}
	}
	first, second := c.Value(), c.Value()
	if first != finalCount || second != finalCount {
		t.Errorf("common-start final reads got %d and %d, want stable %d", first, second, finalCount)
	}
}
