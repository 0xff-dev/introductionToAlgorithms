package leetcode

func isBalanced(root *TreeNode) bool {
	depth := 0
	return isBalancedAux(root, &depth)
}

func isBalancedAux(root *TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}

	left, right := 0, 0
	if isBalancedAux(root.Left, &left) && isBalancedAux(root.Right, &right) {
		diff := left - right
		if diff >= -1 && diff <= 1 {
			*depth = left + 1
			if diff < 0 {
				*depth = right + 1
			}
			return true
		}
	}
	return false
}
