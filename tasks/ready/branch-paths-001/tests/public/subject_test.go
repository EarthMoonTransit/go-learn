package subject_test

import (
	"reflect"
	subject "task.local/branch-paths-001/subject"
	"testing"
)

type nodeState struct {
	node     *subject.Node
	label    int
	children []*subject.Node
}

func capture(node *subject.Node) []nodeState {
	if node == nil {
		return nil
	}
	states := []nodeState{{node, node.Label, append([]*subject.Node(nil), node.Children...)}}
	for _, child := range node.Children {
		states = append(states, capture(child)...)
	}
	return states
}
func preserved(t *testing.T, states []nodeState) {
	t.Helper()
	for _, s := range states {
		if s.node.Label != s.label || len(s.node.Children) != len(s.children) {
			t.Error("input node or child slice changed")
			continue
		}
		for i, c := range s.children {
			if s.node.Children[i] != c {
				t.Error("input child pointer changed")
			}
		}
	}
}
func equalPaths(t *testing.T, got, want [][]int) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("path count=%d; want%d; paths=%v", len(got), len(want), got)
		return
	}
	for i := range want {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Errorf("path[%d]=%v; want full ordered path%v", i, got[i], want[i])
		}
	}
}
func branchTree() *subject.Node {
	return &subject.Node{Label: 1, Children: []*subject.Node{{Label: 2}, {Label: 3, Children: []*subject.Node{{Label: 4}, nil, {Label: -1}}}, nil, {Label: 0}}}
}
func TestPathShapeAndOrder(t *testing.T) {
	cases := []struct {
		node *subject.Node
		want [][]int
	}{
		{nil, nil}, {&subject.Node{Label: 0}, [][]int{{0}}}, {&subject.Node{Label: -9, Children: []*subject.Node{nil, nil}}, [][]int{{-9}}},
		{branchTree(), [][]int{{1, 2}, {1, 3, 4}, {1, 3, -1}, {1, 0}}},
		{&subject.Node{Label: -2, Children: []*subject.Node{{Label: 2}, {Label: 2}}}, [][]int{{-2, 2}, {-2, 2}}},
	}
	for i, c := range cases {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			before := capture(c.node)
			got := subject.Paths(c.node)
			equalPaths(t, got, c.want)
			preserved(t, before)
		})
	}
}
func TestPathsOwnTheirStorage(t *testing.T) {
	root := &subject.Node{Label: 5, Children: []*subject.Node{{Label: 1}, {Label: 2}, {Label: 3}}}
	want := [][]int{{5, 1}, {5, 2}, {5, 3}}
	before := capture(root)
	first := subject.Paths(root)
	equalPaths(t, first, want)
	second := subject.Paths(root)
	equalPaths(t, second, want)
	for i, path := range first {
		if len(path) > 0 {
			path[0] = 100 + i
		}
	}
	equalPaths(t, second, want)
	equalPaths(t, subject.Paths(root), want)
	preserved(t, before)
	// One chosen path cannot share even its prefix storage with siblings.
	third := subject.Paths(root)
	if len(third) == 3 && len(third[0]) > 0 {
		third[0][0] = 222
		equalPaths(t, third[1:], want[1:])
	}
	other := subject.Paths(&subject.Node{Label: 8})
	equalPaths(t, other, [][]int{{8}})
	equalPaths(t, second, want)
}
