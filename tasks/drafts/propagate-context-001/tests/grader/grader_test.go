package subject_test

import (
	"context"
	"task.local/propagate-context-001/subject"
	"testing"
	"testing/synctest"
)

func TestGraderCase01(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		release := make(chan struct{})
		released := false
		unblock := func() {
			if !released {
				close(release)
				released = true
			}
		}
		defer unblock()
		returned := make(chan int, 1)
		go func() {
			n, _ := subject.CallWithContext(ctx, func(context.Context) (int, error) { <-release; return 9, nil })
			returned <- n
		}()
		synctest.Wait()
		select {
		case <-returned:
			t.Error("returned before callback")
		default:
		}
		unblock()
		synctest.Wait()
		select {
		case n := <-returned:
			if n != 9 {
				t.Error("callback result lost")
			}
		default:
			t.Error("callback not joined")
		}
	})
}
