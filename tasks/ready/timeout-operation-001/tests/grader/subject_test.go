package subject_test

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)
import "task.local/timeout-operation-001/subject"

func TestTimeoutNilErrorAndIndependentCalls(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		defer cancel()
		var previous context.Context
		for _, want := range []int{0, -7, 91} {
			var child context.Context
			n, err := subject.RunWithTimeout(parent, 7*time.Second, func(ctx context.Context) (int, error) {
				child = ctx
				if previous != nil && ctx == previous {
					t.Error("reused child context, want independent call")
				}
				return want, nil
			})
			synctest.Wait()
			assertResult(t, n, err, want, nil)
			assertReleased(t, child)
			previous = child
			if parent.Err() != nil {
				t.Error("independent child canceled parent, want live parent")
			}
		}
	})
}
