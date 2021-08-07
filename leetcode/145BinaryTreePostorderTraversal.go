package leetcode

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
			stack = append(stack, root)
		}
		length := len(stack)
		if length > 0 {
			top := stack[length-1]
			stack = stack[:length-1]
			if length == 1 || top != stack[len(stack)-1] {
				res = append(res, top.Val)
				root = nil
				continue
			}
			root = top.Right
		}
	}

	return res
}
