package leetcode

func kthSmallest230(root *TreeNode, k int) int {
	var inOrder func(*TreeNode)
	which := 1
	ans := -1
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inOrder(root.Left)
		}
		if which == k && ans == -1 {
			ans = root.Val
			return
		}
		which++
		if root.Right != nil {
			inOrder(root.Right)
		}
	}
	inOrder(root)
	return ans
}
