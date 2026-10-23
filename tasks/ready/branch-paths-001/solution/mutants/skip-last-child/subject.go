package subject

type Node struct {
	Label    int
	Children []*Node
}

func Paths(root *Node) [][]int {
	result := make([][]int, 0)
	var visit func(*Node, []int)
	visit = func(node *Node, prefix []int) {
		if node == nil {
			return
		}
		prefix = append(prefix, node.Label)
		children := 0
		for _, child := range node.Children {
			if child != nil {
				children++
			}
		}
		if children == 0 {
			result = append(result, append([]int(nil), prefix...))
			return
		}
		for i, child := range node.Children {
			if i < len(node.Children)-1 {
				visit(child, prefix)
			}
		}
	}
	visit(root, make([]int, 0, 6))
	return result
}
