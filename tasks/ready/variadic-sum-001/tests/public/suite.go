package publictests

import "testing"

type API struct {
	Sum      func(...int) int
	SumSlice func([]int) int
}

func Run(t *testing.T, api API) {
	t.Helper()
	tests := []struct {
		name string
		got  int
		want int
	}{
		{name: "no arguments", got: api.Sum(), want: 0},
		{name: "separate arguments", got: api.Sum(1, 2, 3), want: 6},
		{name: "expanded slice", got: api.SumSlice([]int{4, -1, 2}), want: 5},
	}
	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s: got %d, want %d", test.name, test.got, test.want)
		}
	}
}
