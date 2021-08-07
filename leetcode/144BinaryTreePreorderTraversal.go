package leetcode

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for ; root != nil; root = root.Left {
			res = append(res, root.Val)
			stack = append(stack, root)
		}

		if len(stack) > 0 {
			top := stack[len(stack)-1]
			root = top.Right
			stack = stack[:len(stack)-1]
		}
	}

	return res
}
