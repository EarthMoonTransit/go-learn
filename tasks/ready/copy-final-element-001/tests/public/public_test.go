package subject_test

import (
	"slices"
	"task.local/copy-final-element-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	src := []string{"a", "b", "c", "last"}
	dst := []string{"a", "x", "b", "z", "?"}
	subject.CopyLast(src, dst)
	want := []string{"a", "x", "b", "z", "last"}
	if !slices.Equal(dst, want) || !slices.Equal(src, []string{"a", "b", "c", "last"}) {
		t.Errorf("source/destination got %q/%q want unchanged source/%q", src, dst, want)
	}
}

func TestPublicCase02(t *testing.T) {
	src := []string{""}
	dst := []string{"?"}
	view := dst[:]
	subject.CopyLast(src, dst)
	if view[0] != "" {
		t.Errorf("aliased final value got %q want empty", view[0])
	}
}
