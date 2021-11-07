package leetcode

func nAryPreorder(root *MultiChildNode) []int {
	nodes := make([]int, 0)
	nAryPreorderAux(root, &nodes)
	return nodes
}

func nAryPreorderAux(root *MultiChildNode, nodes *[]int) {
	if root == nil {
		return
	}

	*nodes = append(*nodes, root.Val)
	if len(root.Children) > 0 {
		for idx := 0; idx < len(root.Children); idx++ {
			nAryPreorderAux(root.Children[idx], nodes)
		}
	}
}
