package leetcode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	if left > right {
		return right + 1
	}
	return left + 1
	/*
		if root.Left == nil && root.Right == nil {
			return 1
		}
		min := 100000
		if root.Left != nil {
			left := minDepth(root.Left)
			if left < min {
				min = left
			}
		}
		if root.Right != nil {
			right := minDepth(root.Right)
			if right < min {
				min = right
			}
		}
		return min + 1
	*/
}
