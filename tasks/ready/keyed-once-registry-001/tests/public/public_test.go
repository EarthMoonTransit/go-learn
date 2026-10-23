package subject_test

import (
	"errors"
	"sync"
	"sync/atomic"
	"task.local/keyed-once-registry-001/subject"
	"testing"
	"testing/synctest"
)

type registryResult struct {
	key   string
	value int
	err   error
}

func constructRegistry(t *testing.T, load func(string) (int, error), rescue func()) *subject.Registry {
	t.Helper()
	made := make(chan *subject.Registry, 1)
	go func() { made <- subject.NewRegistry(load) }()
	synctest.Wait() // No Get caller or mutex contender exists in this phase.
	var registry *subject.Registry
	select {
	case registry = <-made:
	default:
		t.Error("constructor invoked loader before Get and did not return lazily")
		rescue()
		synctest.Wait()
		select {
		case registry = <-made:
		default:
			t.Error("constructor did not return after independent loader release")
		}
	}
	if registry == nil {
		t.Error("NewRegistry returned nil")
	}
	return registry
}

func TestRegistryLazyConstruction(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		gate := make(chan struct{})
		opened := false
		rescue := func() {
			if !opened {
				close(gate)
				opened = true
			}
		}
		defer rescue()
		var calls atomic.Int32
		constructRegistry(t, func(string) (int, error) { calls.Add(1); <-gate; return 1, nil }, rescue)
		if got := calls.Load(); got != 0 {
			t.Errorf("construction loader calls=%d; want 0", got)
		}
		rescue()
		synctest.Wait()
	})
}

func TestRegistryHeldAndConcurrentKeys(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		gates := map[string]chan struct{}{"": make(chan struct{}), "A": make(chan struct{}), "B": make(chan struct{}), "C": make(chan struct{})}
		fallback := make(chan struct{})
		opened := false
		rescue := func() {
			if !opened {
				for _, gate := range gates {
					close(gate)
				}
				close(fallback)
				opened = true
			}
		}
		defer rescue()
		markerA, markerC := errors.New("A failure"), errors.New("C failure")
		want := map[string]registryResult{"": {"", -5, nil}, "A": {"A", 7, markerA}, "B": {"B", 0, nil}, "C": {"C", -3, markerC}}
		var countsMu sync.Mutex
		counts := make(map[string]int)
		load := func(key string) (int, error) {
			countsMu.Lock()
			counts[key]++
			countsMu.Unlock()
			gate, ok := gates[key]
			if !ok {
				gate = fallback
			}
			<-gate
			expected, ok := want[key]
			if !ok {
				return 99, errors.New("unexpected key")
			}
			return expected.value, expected.err
		}
		registry := constructRegistry(t, load, rescue)
		if registry == nil {
			return
		}
		done := make(chan registryResult, 16)
		invoke := func(key string) { value, err := registry.Get(key); done <- registryResult{key, value, err} }
		go invoke("A")
		synctest.Wait() // Only one Get exists; no held Once/mutex contenders.
		var results []registryResult
		if !opened {
			select {
			case got := <-done:
				results = append(results, got)
				t.Errorf("first Get returned (%d,%v) before loader completed", got.value, got.err)
			default:
			}
			countsMu.Lock()
			got := counts["A"]
			countsMu.Unlock()
			if got != 1 {
				t.Errorf("held key A loader calls=%d; want 1", got)
			}
		}
		keys := []string{"A", "B", "", "C", "A", "B"}
		started := make(chan struct{}, len(keys))
		for _, key := range keys {
			go func() { started <- struct{}{}; invoke(key) }()
		}
		for range keys {
			<-started
		}
		// Acknowledgements precede Get. They do not prove entry into Once.
		// Release ALL callback gates before Wait, even for bad global-lock variants.
		rescue()
		synctest.Wait()
		for len(results) < 1+len(keys) {
			select {
			case got := <-done:
				results = append(results, got)
			default:
				t.Error("Get missing after independent release of every loader")
				return
			}
		}
		for _, got := range results {
			expected := want[got.key]
			if got.value != expected.value || got.err != expected.err {
				t.Errorf("key %q complete pair=(%d,%v); want (%d,exact %v)", got.key, got.value, got.err, expected.value, expected.err)
			}
		}
		for _, key := range []string{"", "A", "B", "C"} {
			value, err := registry.Get(key)
			expected := want[key]
			if value != expected.value || err != expected.err {
				t.Errorf("late key %q=(%d,%v); want cached (%d,exact %v)", key, value, err, expected.value, expected.err)
			}
		}
		synctest.Wait()
		countsMu.Lock()
		defer countsMu.Unlock()
		for _, key := range []string{"", "A", "B", "C"} {
			if counts[key] != 1 {
				t.Errorf("key %q loader calls=%d; want 1 including cached failures", key, counts[key])
			}
		}
		for key, n := range counts {
			if _, ok := want[key]; !ok {
				t.Errorf("loader received unexpected key %q (%d calls)", key, n)
			}
		}
	})
}

func TestRegistrySequentialKeyLifetimes(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		marker := errors.New("cached failure")
		var mu sync.Mutex
		counts := map[string]int{}
		registry := subject.NewRegistry(func(key string) (int, error) {
			mu.Lock()
			counts[key]++
			mu.Unlock()
			if key == "bad" {
				return 8, marker
			}
			return len(key), nil
		})
		if registry == nil {
			t.Error("NewRegistry returned nil")
			return
		}
		for _, key := range []string{"bad", "", "ok", "bad", "ok", ""} {
			value, err := registry.Get(key)
			wantValue, wantErr := len(key), error(nil)
			if key == "bad" {
				wantValue, wantErr = 8, marker
			}
			if value != wantValue || err != wantErr {
				t.Errorf("key %q result=(%d,%v); want (%d,exact %v)", key, value, err, wantValue, wantErr)
			}
		}
		synctest.Wait()
		mu.Lock()
		defer mu.Unlock()
		for _, key := range []string{"bad", "", "ok"} {
			if counts[key] != 1 {
				t.Errorf("sequential key %q loader calls=%d; want 1", key, counts[key])
			}
		}
	})
}

func TestRegistryInstanceIndependence(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		firstError, secondError := errors.New("first"), errors.New("second")
		var firstCalls, secondCalls atomic.Int32
		first := subject.NewRegistry(func(string) (int, error) { firstCalls.Add(1); return 13, firstError })
		second := subject.NewRegistry(func(string) (int, error) { secondCalls.Add(1); return -9, secondError })
		if first == nil || second == nil {
			t.Error("NewRegistry returned nil")
			return
		}
		for range 2 {
			av, ae := first.Get("same")
			bv, be := second.Get("same")
			if av != 13 || ae != firstError || bv != -9 || be != secondError {
				t.Errorf("independent registries=(%d,%v),(%d,%v); want (13,first),(-9,second)", av, ae, bv, be)
			}
		}
		synctest.Wait()
		if firstCalls.Load() != 1 || secondCalls.Load() != 1 {
			t.Errorf("independent loader counts=(%d,%d); want (1,1)", firstCalls.Load(), secondCalls.Load())
		}
	})
}
