package subject_test

import (
	"context"
	"errors"
	"task.local/timeout-operation-001/subject"
	"testing"
	"testing/synctest"
	"time"
)

type timeoutKey struct{}

var operationError = errors.New("operation result")

func bubbleOrigin() time.Time { return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) }
func assertResult(t *testing.T, n int, err error, want int, wantErr error) {
	t.Helper()
	if n != want || err != wantErr {
		t.Errorf("operation result = (%d,%v), want unchanged (%d,%v)", n, err, want, wantErr)
	}
}
func assertDeadline(t *testing.T, ctx context.Context, want time.Time) {
	t.Helper()
	got, ok := ctx.Deadline()
	if !ok || !got.Equal(want) {
		t.Errorf("child deadline = (%v,%t), want (%v,true)", got, ok, want)
	}
}
func assertReleased(t *testing.T, ctx context.Context) {
	t.Helper()
	if ctx == nil {
		t.Error("operation context = nil, want derived context")
		return
	}
	select {
	case <-ctx.Done():
	default:
		t.Error("child after return = live, want canceled/released")
	}
}
func TestTimeoutFastResultAndRelease(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.WithValue(context.Background(), timeoutKey{}, "marker"))
		defer cancel()
		calls := 0
		var child context.Context
		n, err := subject.RunWithTimeout(parent, 5*time.Second, func(ctx context.Context) (int, error) {
			calls++
			child = ctx
			if ctx == parent {
				t.Error("operation received parent, want distinct derived context")
			}
			if got := ctx.Value(timeoutKey{}); got != "marker" {
				t.Errorf("child value = %v, want marker", got)
			}
			if got := ctx.Err(); got != nil {
				t.Errorf("child before operation returns = %v, want live context", got)
			}
			assertDeadline(t, ctx, bubbleOrigin().Add(5*time.Second))
			return 41, operationError
		})
		synctest.Wait()
		assertResult(t, n, err, 41, operationError)
		if calls != 1 {
			t.Errorf("operation calls = %d, want 1", calls)
		}
		assertReleased(t, child)
		if got := parent.Err(); got != nil {
			t.Errorf("parent after child release = %v, want live", got)
		}
	})
}
func TestTimeoutAlreadyExpiredOrCanceled(t *testing.T) {
	for _, budget := range []time.Duration{0, -2 * time.Second} {
		synctest.Test(t, func(t *testing.T) {
			parent, cancel := context.WithCancel(context.Background())
			defer cancel()
			calls := 0
			var child context.Context
			n, err := subject.RunWithTimeout(parent, budget, func(ctx context.Context) (int, error) {
				calls++
				child = ctx
				assertDeadline(t, ctx, bubbleOrigin().Add(budget))
				if got := ctx.Err(); !errors.Is(got, context.DeadlineExceeded) {
					t.Errorf("nonpositive budget error = %v, want DeadlineExceeded", got)
				}
				return 17, operationError
			})
			synctest.Wait()
			assertResult(t, n, err, 17, operationError)
			assertReleased(t, child)
			if calls != 1 {
				t.Errorf("expired operation calls = %d, want 1", calls)
			}
			if got := parent.Err(); got != nil {
				t.Errorf("nonpositive child parent error = %v, want nil", got)
			}
		})
	}
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		cancel()
		calls := 0
		n, err := subject.RunWithTimeout(parent, 5*time.Second, func(ctx context.Context) (int, error) {
			calls++
			if got := ctx.Err(); !errors.Is(got, context.Canceled) {
				t.Errorf("already canceled parent child error = %v, want Canceled", got)
			}
			return 19, operationError
		})
		synctest.Wait()
		assertResult(t, n, err, 19, operationError)
		if calls != 1 {
			t.Errorf("already canceled operation calls = %d, want 1", calls)
		}
	})
}
func TestTimeoutBudgetAndEarlierParent(t *testing.T) {
	for _, earlierParent := range []bool{false, true} {
		synctest.Test(t, func(t *testing.T) {
			origin := bubbleOrigin()
			parent, cancel := context.WithCancel(context.Background())
			defer cancel()
			boundary := 5 * time.Second
			if earlierParent {
				var stop context.CancelFunc
				parent, stop = context.WithDeadline(parent, origin.Add(3*time.Second))
				defer stop()
				boundary = 3 * time.Second
			}
			guard := time.NewTimer(6 * time.Second)
			defer guard.Stop()
			calls := 0
			n, err := subject.RunWithTimeout(parent, 5*time.Second, func(ctx context.Context) (int, error) {
				calls++
				assertDeadline(t, ctx, origin.Add(boundary))
				select {
				case <-ctx.Done():
				case <-guard.C:
					t.Error("virtual guard reached with live child, want deadline cancellation first")
				}
				if got := time.Since(origin); got != boundary {
					t.Errorf("virtual expiry elapsed = %v, want %v", got, boundary)
				}
				if got := ctx.Err(); !errors.Is(got, context.DeadlineExceeded) {
					t.Errorf("expiry child error = %v, want DeadlineExceeded", got)
				}
				return 53, ctx.Err()
			})
			assertResult(t, n, err, 53, context.DeadlineExceeded)
			if calls != 1 {
				t.Errorf("expiry operation calls = %d, want 1", calls)
			}
			if !earlierParent && parent.Err() != nil {
				t.Errorf("child-only timeout parent error = %v, want nil", parent.Err())
			}
		})
	}
}

type timeoutResult struct {
	n   int
	err error
}

func TestTimeoutWaitsForCooperativeCleanup(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		release, stop := make(chan struct{}), make(chan struct{})
		released := false
		unblock := func() {
			if !released {
				close(release)
				released = true
			}
		}
		defer func() { cancel(); close(stop); unblock(); synctest.Wait() }()
		entered, observed := make(chan struct{}, 2), make(chan struct{}, 2)
		returned := make(chan timeoutResult, 1)
		go func() {
			n, err := subject.RunWithTimeout(parent, time.Hour, func(ctx context.Context) (int, error) {
				entered <- struct{}{}
				select {
				case <-ctx.Done():
					observed <- struct{}{}
				case <-stop:
				}
				<-release
				return 67, operationError
			})
			returned <- timeoutResult{n, err}
		}()
		synctest.Wait()
		select {
		case <-entered:
		default:
			t.Error("operation entry = absent, want one invocation")
		}
		cancel()
		synctest.Wait()
		select {
		case <-observed:
		default:
			t.Error("operation parent cancellation = absent, want propagated Done")
		}
		var result timeoutResult
		early := false
		select {
		case result = <-returned:
			early = true
			t.Error("wrapper returned before operation cleanup, want pending until operation returns")
		default:
		}
		unblock()
		synctest.Wait()
		if !early {
			select {
			case result = <-returned:
			default:
				t.Error("wrapper result after cleanup = absent, want completed operation result")
			}
		}
		assertResult(t, result.n, result.err, 67, operationError)
	})
}

func TestTimeoutFastSuccessAndRelease(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithDeadline(context.WithValue(context.Background(), timeoutKey{}, "marker"), bubbleOrigin().Add(3*time.Second))
		defer cancel()
		calls := 0
		var child context.Context
		n, err := subject.RunWithTimeout(parent, 5*time.Second, func(ctx context.Context) (int, error) {
			calls++
			child = ctx
			if ctx == parent {
				t.Error("operation received parent, want distinct derived context")
			}
			if got := ctx.Value(timeoutKey{}); got != "marker" {
				t.Errorf("child value = %v, want marker", got)
			}
			if got := ctx.Err(); got != nil {
				t.Errorf("child before operation returns = %v, want live context", got)
			}
			assertDeadline(t, ctx, bubbleOrigin().Add(3*time.Second))
			return 41, nil
		})
		synctest.Wait()
		assertResult(t, n, err, 41, nil)
		if calls != 1 {
			t.Errorf("operation calls = %d, want 1", calls)
		}
		assertReleased(t, child)
		if got := parent.Err(); got != nil {
			t.Errorf("parent after child release = %v, want live", got)
		}
	})
}
