package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	has := false
	hasPathSumAux(root, targetSum, 0, &has)
	return has
}

func hasPathSumAux(root *TreeNode, targetSum, now int, equal *bool) {
	if root == nil {
		return
	}

	now += root.Val
	if root.Left == nil && root.Right == nil && now == targetSum {
		*equal = true
		return
	}

	hasPathSumAux(root.Left, targetSum, now, equal)
	hasPathSumAux(root.Right, targetSum, now, equal)
}
