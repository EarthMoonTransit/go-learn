package gradertests

import "testing"

type API struct {
	Sum      func(...int) int
	SumSlice func([]int) int
}

func Run(t *testing.T, api API) {
	t.Helper()
	values := []int{10, -3, 7, 0, -4, 11, 2}
	if got, want := api.Sum(values...), 23; got != want {
		t.Errorf("Sum with more than three values: got %d, want %d", got, want)
	}
	if got := api.SumSlice(nil); got != 0 {
		t.Errorf("SumSlice(nil): got %d, want 0", got)
	}
	input := []int{5, 6, 7, 8}
	if got, want := api.SumSlice(input), 26; got != want {
		t.Errorf("SumSlice with four values: got %d, want %d", got, want)
	}
	if want := []int{5, 6, 7, 8}; !equal(input, want) {
		t.Errorf("SumSlice modified input: got %v, want %v", input, want)
	}
}

func equal(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
