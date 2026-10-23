package subject_test

import (
	"errors"
	"sync"
	"task.local/keyed-once-registry-001/subject"
	"testing"
	"testing/synctest"
)

func TestRegistryFiniteKeyExtension(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		keys := []string{"", "x", "repeat", "12345678"}
		marker := errors.New("odd length")
		var mu sync.Mutex
		counts := map[string]int{}
		registry := subject.NewRegistry(func(key string) (int, error) {
			mu.Lock()
			counts[key]++
			mu.Unlock()
			if len(key)%2 != 0 {
				return len(key) - 4, marker
			}
			return len(key) - 4, nil
		})
		if registry == nil {
			t.Error("NewRegistry returned nil")
			return
		}
		done := make(chan registryResult, 16)
		for i := 0; i < 16; i++ {
			key := keys[i%4]
			go func() { v, e := registry.Get(key); done <- registryResult{key, v, e} }()
		}
		synctest.Wait() // All callbacks return immediately; no held gates or locks.
		for i := 0; i < 16; i++ {
			select {
			case got := <-done:
				wantError := error(nil)
				if len(got.key)%2 != 0 {
					wantError = marker
				}
				if got.value != len(got.key)-4 || got.err != wantError {
					t.Errorf("finite key %q pair=(%d,%v); want (%d,exact %v)", got.key, got.value, got.err, len(got.key)-4, wantError)
				}
			default:
				t.Error("finite Get did not finish")
				return
			}
		}
		mu.Lock()
		defer mu.Unlock()
		for _, key := range keys {
			if counts[key] != 1 {
				t.Errorf("finite key %q calls=%d; want 1", key, counts[key])
			}
		}
	})
}
