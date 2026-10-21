package subject_test

import (
	"reflect"
	subject "task.local/guarded-stage-values-001/subject"
	"testing"
)

func TestStagesAtEveryStopPosition(t *testing.T) {
	for stop := -1; stop < 5; stop++ {
		trace := []int{}
		steps := make([]func(int) (int, bool), 5)
		for i := range steps {
			index := i
			steps[i] = func(v int) (int, bool) { trace = append(trace, index*100+v); return v + index + 1, index != stop }
		}
		want, count := 3, 0
		expected := []int{}
		for i := 0; i < 5; i++ {
			expected = append(expected, i*100+want)
			if i == stop {
				break
			}
			want += i + 1
			count++
		}
		got, done := subject.RunStages(3, steps)
		if got != want || done != count || !reflect.DeepEqual(trace, expected) {
			t.Fatalf("stop=%d: value/count=%d/%d calls=%v; want %d/%d %v", stop, got, done, trace, want, count, expected)
		}
	}
}
