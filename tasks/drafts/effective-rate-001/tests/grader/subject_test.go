package subject_test

import (
	"testing"

	"task.local/effective-rate-001/subject"
)

func TestEffectiveRate(t *testing.T) {
	tests := []struct {
		name       string
		items      int
		percentage float64
		want       int
	}{
		{name: "zero percentage", items: 200, percentage: 0, want: 0},
		{name: "half truncates", items: 3, percentage: 50, want: 1},
		{name: "calculation before conversion", items: 99, percentage: 50, want: 49},
		{name: "full rate", items: 17, percentage: 100, want: 17},
		{name: "fractional percentage", items: 101, percentage: 99.5, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subject.EffectiveRate(tt.items, tt.percentage); got != tt.want {
				t.Fatalf("got %d; want %d", got, tt.want)
			}
		})
	}
}
