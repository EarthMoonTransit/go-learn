package subject_test

import (
	"slices"
	"task.local/variadic-aliasing-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	values := []int{1, 2, 3}
	if !subject.ReplaceFirst(9, values...) || !slices.Equal(values, []int{9, 2, 3}) {
		t.Errorf("expanded values got%v", values)
	}
}

func TestPublicCase02(t *testing.T) {
	whole := []int{7, 1, 2, 8}
	view := whole[1:3]
	beforeCap := cap(view)
	if !subject.ReplaceFirst(-4, view...) || !slices.Equal(whole, []int{7, -4, 2, 8}) || len(view) != 2 || cap(view) != beforeCap {
		t.Errorf("subslice got%v len%d cap%d", whole, len(view), cap(view))
	}
	if subject.ReplaceFirst(1) || subject.ReplaceFirst(1, []int(nil)...) {
		t.Error("empty accepted")
	}
	if !subject.ReplaceFirst(0, 2, 3) {
		t.Error("direct arguments rejected")
	}
}
