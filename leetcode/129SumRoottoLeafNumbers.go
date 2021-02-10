package leetcode

func sumNumbers(root *TreeNode) int {
	r := 0
	sumNumbersAux(root, 0, &r)
	return r
}

func sumNumbersAux(root *TreeNode, now int, sum *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*sum += now*10 + root.Val
		return
	}

	sumNumbersAux(root.Left, now*10+root.Val, sum)
	sumNumbersAux(root.Right, now*10+root.Val, sum)
}
