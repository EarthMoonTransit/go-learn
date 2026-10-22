package subject_test

import (
	"reflect"
	subject "task.local/first-pair-positions-001/subject"
	"testing"
)

func TestOrderedPairFirstPositions(t *testing.T) {
	cases := []struct {
		pairs [][2]string
		want  []int
	}{
		{[][2]string{{"a", "b"}, {"x", "y"}, {"a", "b"}, {"a", "b"}, {"x", "y"}}, []int{0, 1, 0, 0, 1}},
		{[][2]string{{"ab", "c"}, {"a", "bc"}, {"ab", "c"}}, []int{0, 1, 0}},
		{[][2]string{{"a\x00b", "c"}, {"a", "b\x00c"}, {"a\x00b", "c"}}, []int{0, 1, 0}},
		{[][2]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"z", "b"}, {"a", "b"}}, []int{0, 1, 2, 3, 0}},
		{[][2]string{{"\xff", "x"}, {"\xfe", "x"}, {"�", "x"}, {"x", "\xff"}, {"x", "\xfe"}, {"x", "�"}, {"\xff", "x"}, {"x", "\xfe"}}, []int{0, 1, 2, 3, 4, 5, 0, 4}},
		{[][2]string{{"", ""}, {"", "x"}, {"x", ""}, {"", ""}, {"界", "🙂"}, {"界", "🙂"}}, []int{0, 1, 2, 0, 4, 4}},
	}
	for _, c := range cases {
		before := append([][2]string(nil), c.pairs...)
		got := subject.FirstPositions(c.pairs)
		if !reflect.DeepEqual(got, c.want) {
			t.Fatalf("pairs=%q positions=%v; want %v", before, got, c.want)
		}
		if !reflect.DeepEqual(c.pairs, before) {
			t.Fatalf("input changed: %q; want %q", c.pairs, before)
		}
	}
}
func TestPairCallsAndEmptyInputs(t *testing.T) {
	for _, p := range [][][2]string{nil, {}, {{"same", "value"}}, {{"other", "key"}, {"same", "value"}, {"same", "value"}}} {
		got := subject.FirstPositions(p)
		want := []int{}
		if len(p) == 1 {
			want = []int{0}
		}
		if len(p) == 3 {
			want = []int{0, 1, 1}
		}
		if len(got) != len(want) {
			t.Fatalf("input=%q len=%d; want %d", p, len(got), len(want))
		}
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("independent call input=%q positions=%v; want %v", p, got, want)
			}
		}
	}
}
