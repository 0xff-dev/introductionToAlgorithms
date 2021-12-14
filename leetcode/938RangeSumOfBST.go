package leetcode

func rangeSumBST(root *TreeNode, low, high int) int {
	sum := 0
	innerTraversal(root, low, high, &sum)
	return sum
}

func innerTraversal(root *TreeNode, low, high int, sum *int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		innerTraversal(root.Left, low, high, sum)
	}
	if root.Val >= low && root.Val <= high {
		*sum += root.Val
	}
	if root.Right != nil {
		innerTraversal(root.Right, low, high, sum)
	}
}
