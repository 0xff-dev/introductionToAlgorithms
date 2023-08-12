package leetcode

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}
	a := flipEquiv(root1.Left, root2.Left)
	b := flipEquiv(root1.Right, root2.Right)

	if a && b {
		return true
	}
	a = flipEquiv(root1.Left, root2.Right)
	b = flipEquiv(root1.Right, root2.Left)
	if a && b {
		return true
	}
	return false
}
