package subject_test

import (
	subject "task.local/branch-paths-001/subject"
	"testing"
)

func generated(depth, label int) (*subject.Node, [][]int) {
	n := &subject.Node{Label: label}
	if depth == 0 {
		return n, [][]int{{label}}
	}
	a, pa := generated(depth-1, label-1)
	b, pb := generated(depth-1, label+1)
	n.Children = []*subject.Node{a, nil, b}
	var want [][]int
	for _, paths := range [][][]int{pa, pb} {
		for _, p := range paths {
			want = append(want, append([]int{label}, p...))
		}
	}
	return n, want
}
func TestBoundedTreeMatrix(t *testing.T) {
	for depth := 0; depth <= 3; depth++ {
		for seed := -3; seed <= 3; seed++ {
			root, want := generated(depth, seed)
			before := capture(root)
			equalPaths(t, subject.Paths(root), want)
			preserved(t, before)
		}
	}
	root := &subject.Node{Label: 1}
	at := root
	for i := 2; i <= 6; i++ {
		child := &subject.Node{Label: i}
		at.Children = []*subject.Node{nil, child}
		at = child
	}
	equalPaths(t, subject.Paths(root), [][]int{{1, 2, 3, 4, 5, 6}})
}
