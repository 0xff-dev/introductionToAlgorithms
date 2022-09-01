package leetcode

func goodNodes(root *TreeNode) int {
	ans := 0
	goodNodesDfs(root, &ans, root.Val)
	return ans
}

func goodNodesDfs(root *TreeNode, ans *int, cmp int) {
	if root == nil {
		return
	}
	if root.Val >= cmp {
		*ans++
		cmp = root.Val
	}
	goodNodesDfs(root.Left, ans, cmp)
	goodNodesDfs(root.Right, ans, cmp)
}
