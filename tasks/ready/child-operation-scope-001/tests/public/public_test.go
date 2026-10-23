package subject_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"task.local/child-operation-scope-001/subject"
)

type scopeKey struct{}

func scopeDone(ctx context.Context) bool {
	if ctx == nil {
		return false
	}
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func checkScopeSequence(t *testing.T, values []int, errorAt int, preCanceled, deadline bool) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.WithValue(context.Background(), scopeKey{}, "marker"))
		defer cancel()
		if deadline {
			var release context.CancelFunc
			parent, release = context.WithDeadline(parent, time.Date(2000, 1, 1, 1, 0, 0, 0, time.UTC))
			defer release()
		}
		if preCanceled {
			cancel()
		}
		sentinel := errors.New("operation failure")
		expected := make([]subject.Result, len(values))
		operations := make([]func(context.Context) (int, error), len(values))
		captured := make([]context.Context, len(values))
		var mu sync.Mutex
		var order []int
		for index, value := range values {
			index, value := index, value
			var failure error
			if index == errorAt {
				failure = sentinel
			}
			expected[index] = subject.Result{Value: value, Err: failure}
			operations[index] = func(child context.Context) (int, error) {
				mu.Lock()
				defer mu.Unlock()
				order = append(order, index)
				if child == nil {
					t.Errorf("operation %d received nil context", index)
					return value, failure
				}
				if child.Value(scopeKey{}) != "marker" {
					t.Errorf("operation %d lost parent value", index)
				}
				gotDeadline, gotOK := child.Deadline()
				wantDeadline, wantOK := parent.Deadline()
				if gotOK != wantOK || (gotOK && !gotDeadline.Equal(wantDeadline)) {
					t.Errorf("operation %d lost parent deadline: got (%v, %v), want (%v, %v)", index, gotDeadline, gotOK, wantDeadline, wantOK)
				}
				if preCanceled {
					if !scopeDone(child) || child.Err() != context.Canceled {
						t.Errorf("operation %d lost parent cancellation: Done=%v Err=%v", index, scopeDone(child), child.Err())
					}
				} else {
					if child.Done() == nil || child.Done() == parent.Done() {
						t.Errorf("operation %d did not receive a fresh child", index)
					}
					if scopeDone(child) {
						t.Errorf("operation %d child canceled before invocation", index)
					}
					for before, previous := range captured {
						if previous != nil && child.Done() == previous.Done() {
							t.Errorf("operation %d did not receive a fresh child; reused operation %d scope", index, before)
						}
					}
				}
				if index > 0 {
					previous := captured[index-1]
					if previous == nil || !scopeDone(previous) {
						t.Errorf("previous child still live or absent at next callback %d", index)
					}
				}
				captured[index] = child
				if !preCanceled && scopeDone(parent) {
					t.Errorf("parent unexpectedly canceled during operation %d", index)
				}
				return value, failure
			}
		}
		results := subject.RunScoped(parent, operations)
		synctest.Wait()
		mu.Lock()
		gotOrder := append([]int(nil), order...)
		children := append([]context.Context(nil), captured...)
		mu.Unlock()
		wantOrder := make([]int, len(values))
		for index := range wantOrder {
			wantOrder[index] = index
		}
		if len(gotOrder) != len(wantOrder) || (len(wantOrder) > 0 && !reflect.DeepEqual(gotOrder, wantOrder)) {
			t.Errorf("callback order/count = %v; want %v", gotOrder, wantOrder)
		}
		if len(results) != len(expected) {
			t.Errorf("result count = %d; want %d", len(results), len(expected))
		}
		for index, want := range expected {
			if index >= len(results) {
				break
			}
			got := results[index]
			if got.Value != want.Value {
				t.Errorf("result[%d].Value = %d; want %d", index, got.Value, want.Value)
			}
			if got.Err != want.Err {
				t.Errorf("result[%d].Err identity changed: got %v (%T); want original error %v (%T)", index, got.Err, got.Err, want.Err, want.Err)
			}
		}
		for index, child := range children {
			if child != nil && !scopeDone(child) {
				t.Errorf("child %d remains live after return", index)
			}
		}
		if !preCanceled && scopeDone(parent) {
			t.Error("parent unexpectedly canceled after return")
		}
	})
}

func TestPublicScopedSequence(t *testing.T) {
	cases := []struct {
		name     string
		values   []int
		errorAt  int
		canceled bool
		deadline bool
	}{
		{"nil", nil, -1, false, false},
		{"empty", []int{}, -1, false, false},
		{"single-success", []int{9}, -1, false, false},
		{"single-error", []int{-4}, 0, false, false},
		{"siblings-after-error", []int{12, 0, -8, 12}, 1, false, true},
		{"already-canceled", []int{-2, 6, 0, 9}, 1, true, true},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) { checkScopeSequence(t, test.values, test.errorAt, test.canceled, test.deadline) })
	}
}

func TestPublicParentCancellationWaitsForCleanup(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.WithValue(context.Background(), scopeKey{}, "marker"))
		defer cancel()
		rescue := make(chan struct{})
		cleanup := make(chan struct{})
		entered := make(chan struct{}, 32)
		observed := make(chan struct{}, 32)
		completed := make(chan struct{}, 32)
		next := make(chan struct{}, 32)
		returned := make(chan []subject.Result, 1)
		sentinel := errors.New("held operation failed")
		operations := []func(context.Context) (int, error){
			func(child context.Context) (int, error) {
				entered <- struct{}{}
				var done <-chan struct{}
				if child != nil {
					done = child.Done()
				} else {
					t.Error("held operation received nil context")
				}
				select {
				case <-done:
					observed <- struct{}{}
				case <-rescue:
				}
				<-cleanup
				completed <- struct{}{}
				return 31, sentinel
			},
			func(child context.Context) (int, error) {
				next <- struct{}{}
				return -7, nil
			},
		}
		go func() { returned <- subject.RunScoped(parent, operations) }()
		synctest.Wait()
		if len(entered) == 0 {
			t.Error("first callback not invoked before waiting for cleanup")
		}
		if len(returned) != 0 {
			t.Error("returned before callback cleanup")
		}
		if len(next) != 0 {
			t.Error("next callback entered before previous cleanup")
		}
		cancel()
		synctest.Wait()
		if len(observed) == 0 {
			t.Error("held callback lost parent cancellation")
		}
		if len(returned) != 0 {
			t.Error("returned before callback cleanup after parent cancellation")
		}
		if len(next) != 0 {
			t.Error("next callback entered before previous cleanup after cancellation")
		}
		// All useful observations precede the independent rescue and release.
		close(rescue)
		close(cleanup)
		synctest.Wait()
		select {
		case results := <-returned:
			if len(results) != 2 {
				t.Errorf("joined result count = %d; want 2", len(results))
			} else {
				if results[0].Value != 31 || results[0].Err != sentinel {
					t.Errorf("joined result[0] = %#v; want exact held result", results[0])
				}
				if results[1].Value != -7 || results[1].Err != nil {
					t.Errorf("joined result[1] = %#v; want (-7, nil)", results[1])
				}
			}
		default:
			t.Error("wrapper did not finish after fixture rescue and cleanup release")
		}
		if len(entered) != 1 || len(completed) != 1 || len(next) != 1 {
			t.Errorf("callback completion counts = (%d, %d, %d); want (1, 1, 1)", len(entered), len(completed), len(next))
		}
	})
}

func TestPublicSeparateCalls(t *testing.T) {
	for index := 0; index < 2; index++ {
		t.Run(fmt.Sprint(index), func(t *testing.T) { checkScopeSequence(t, []int{index + 3, -index - 8}, -1, false, false) })
	}
}
