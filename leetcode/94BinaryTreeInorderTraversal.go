package leetcode

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := make([]int, 0)
	inorderTraversalAux(root, &r)
	return r
}

func inorderTraversalAux(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorderTraversalAux(root.Left, data)
	}
	*data = append(*data, root.Val)
	if root.Right != nil {
		inorderTraversalAux(root.Right, data)
	}
}
