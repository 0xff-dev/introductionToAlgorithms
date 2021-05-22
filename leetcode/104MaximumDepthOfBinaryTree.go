package leetcode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1
	if left > right {
		return left
	}
	return right
}
