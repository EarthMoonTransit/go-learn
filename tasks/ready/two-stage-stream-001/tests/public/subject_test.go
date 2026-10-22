package subject_test

import (
	"slices"
	subject "task.local/two-stage-stream-001/subject"
	"testing"
	"testing/synctest"
)

func checkStream(t *testing.T, values []int, delta int) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		input := make(chan int, 16)
		returned := make(chan (<-chan int), 1)
		go func() { returned <- subject.Transform(input, delta) }()
		synctest.Wait()
		var output <-chan int
		select {
		case output = <-returned:
		default:
			t.Error("Transform did not return output while input was open")
		}
		var got []int
		closed := false
		drain := func() {
			for {
				synctest.Wait()
				select {
				case v, ok := <-output:
					if !ok {
						closed = true
						return
					}
					got = append(got, v)
				default:
					return
				}
			}
		}
		want := []int{}
		for i, v := range values {
			input <- v
			if v+delta >= 0 {
				want = append(want, v+delta)
			}
			drain()
			if closed {
				t.Errorf("output closed before input at index %d", i)
			}
			if !slices.Equal(got, want) {
				t.Errorf("before EOF after input[%d]=%d delta=%d: output=%v; want streamed prefix %v", i, v, delta, got, want)
			}
		}
		close(input)
		synctest.Wait()
		if output == nil {
			select {
			case output = <-returned:
			default:
				t.Error("Transform never returned output after input close")
			}
		}
		drain()
		if !closed {
			t.Error("output must close after input exhaustion")
		}
		if !slices.Equal(got, want) {
			t.Errorf("final output=%v; want transformed nonnegative values %v", got, want)
		}
	})
}
func TestStreamingBeforeEOF(t *testing.T) { checkStream(t, []int{-4, 0, 2, -1, 2, 4}, 1) }
func TestEmptyAndAllFiltered(t *testing.T) {
	checkStream(t, nil, 0)
	checkStream(t, []int{-20, -3, -2}, -1)
}
func TestZeroNegativeDeltaAndDuplicates(t *testing.T) { checkStream(t, []int{3, 2, 3, 0, 5}, -3) }
