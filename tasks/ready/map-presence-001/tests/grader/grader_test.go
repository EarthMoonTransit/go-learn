package subject_test

import (
	"reflect"
	"task.local/map-presence-001/subject"
	"testing"
)

func TestLookupContract(t *testing.T) {
	for _, key := range []string{"", "a", "zero", "missing", "λ"} {
		for _, fallback := range []int{-1000, 0, 1000} {
			input := map[string]int{"": -3, "a": 1000, "zero": 0, "λ": -1000}
			before := map[string]int{"": -3, "a": 1000, "zero": 0, "λ": -1000}
			want, ok := before[key]
			if !ok {
				want = fallback
			}
			if got := subject.LookupOr(input, key, fallback); got != want {
				t.Fatal("wrong lookup result")
			}
			if !reflect.DeepEqual(input, before) {
				t.Fatal("read-only lookup mutated input")
			}
		}
	}
	for _, input := range []map[string]int{nil, {}} {
		if subject.LookupOr(input, "", 23) != 23 {
			t.Fatal("empty map lookup")
		}
	}
}
