package subject_test

import (
	subject "task.local/first-pair-positions-001/subject"
	"testing"
)

func TestPairValueEquality(t *testing.T) {
	alphabet := []string{"", "a", "b", "a\x00", "\x00a"}
	pairs := make([][2]string, 0)
	for _, a := range alphabet {
		for _, b := range alphabet {
			pairs = append(pairs, [2]string{a, b})
		}
	}
	pairs = append(pairs, pairs...)
	got := subject.FirstPositions(pairs)
	if len(got) != len(pairs) {
		t.Fatalf("length=%d; want %d", len(got), len(pairs))
	}
	for i := range pairs {
		want := i
		for j := 0; j < i; j++ {
			if pairs[j] == pairs[i] {
				want = j
				break
			}
		}
		if got[i] != want {
			t.Fatalf("pair[%d]=%q position=%d; want %d", i, pairs[i], got[i], want)
		}
	}
}
