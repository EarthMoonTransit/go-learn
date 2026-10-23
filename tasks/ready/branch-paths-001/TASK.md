# Construct paths through a branching tree
## Context
Return every root-to-leaf label path.
## API
```go
type Node struct { Label int; Children []*Node }
func Paths(root *Node) [][]int
```
## Requirements
- Nil root has no paths. Ignore nil children; a node without nonnil children is a leaf.
- Return one full path per leaf, in depth-first child-slice order. Keep repeated and negative labels.
- Preserve all input nodes and child slices. Returned paths must be independent of one another and subsequent calls.
- Use authored recursion whose subresults determine the returned paths; do not merely call an unused recursive helper.
- Explain terminal cases, a strictly decreasing measure and how every required branch contributes.
## Constraints
- Inputs are trees, never cycles or shared nodes: at most15nodes, depth6, labels[-9,9]. Nil and empty results both represent no paths. Method review is required alongside behavioral tests.
