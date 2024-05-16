package leetcode

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val == 0 {
		return false
	}
	if root.Val == 1 {
		return true
	}
	l := evaluateTree(root.Left)
	r := evaluateTree(root.Right)
	if root.Val == 2 {
		return l || r
	}
	return l && r
}
