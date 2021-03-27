package leetcode

func levelOrderBottom(root *TreeNode) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}
	levelOrderBottomAux([]*TreeNode{root}, &r)
	return append(r, []int{root.Val})
}

func levelOrderBottomAux(root []*TreeNode, r *[][]int) {
	if len(root) == 0 {
		return
	}
	tmpRoot := make([]*TreeNode, 0)
	for _, node := range root {
		if node.Left != nil {
			tmpRoot = append(tmpRoot, node.Left)
		}
		if node.Right != nil {
			tmpRoot = append(tmpRoot, node.Right)
		}
	}
	if len(tmpRoot) > 0 {
		levelOrderBottomAux(tmpRoot, r)
		values := make([]int, 0)
		for _, node := range tmpRoot {
			values = append(values, node.Val)
		}
		*r = append(*r, values)
	}
}
