package subject_test

import (
	"context"
	"slices"
	"task.local/cancelable-forward-001/subject"
	"testing"
	"testing/synctest"
)

func TestForwardEmptyAndNegativeValues(t *testing.T) {
	for _, values := range [][]int{nil, {-9, -9, 0, 5}} {
		synctest.Test(t, func(t *testing.T) {
			input := make(chan int, len(values))
			for _, v := range values {
				input <- v
			}
			close(input)
			out := subject.Forward(context.Background(), input)
			defer func() { drainForward(out, len(values)+3); synctest.Wait() }()
			var got []int
			for range values {
				synctest.Wait()
				select {
				case v, open := <-out:
					if open {
						got = append(got, v)
					}
				default:
				}
			}
			if !slices.Equal(got, values) {
				t.Errorf("extended forwarded values = %v, want %v", got, values)
			}
			expectForwardClosed(t, out)
		})
	}
}
