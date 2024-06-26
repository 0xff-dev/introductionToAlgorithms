package leetcode

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ok := true
	if root.Left != nil {
		if root.Val != root.Left.Val {
			return false
		}
		ok = ok && isUnivalTree(root.Left)
	}
	if root.Right != nil {
		if root.Val != root.Right.Val {
			return false
		}
		ok = ok && isUnivalTree(root.Right)
	}
	return ok
}
