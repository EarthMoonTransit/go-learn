package subject_test

import (
	"testing"

	"task.local/clamp-tests-001/subject"
)

func TestClamp(t *testing.T) {
	tests := []struct {
		name            string
		value, min, max int
		want            int
	}{
		{name: "below range", value: -4, min: 0, max: 10, want: 0},
		{name: "lower boundary", value: 0, min: 0, max: 10, want: 0},
		{name: "inside range", value: 5, min: 0, max: 10, want: 5},
		{name: "upper boundary", value: 10, min: 0, max: 10, want: 10},
		{name: "above range", value: 14, min: 0, max: 10, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subject.Clamp(tt.value, tt.min, tt.max); got != tt.want {
				t.Fatalf("Clamp(%d, %d, %d) = %d; want %d", tt.value, tt.min, tt.max, got, tt.want)
			}
		})
	}
}
