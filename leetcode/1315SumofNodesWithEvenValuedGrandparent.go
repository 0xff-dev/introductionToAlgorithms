package leetcode

func sumEvenGrandparent(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	var dfs func(*TreeNode, *TreeNode, *int)
	dfs = func(current, parent *TreeNode, ans *int) {
		if current == nil {
			return
		}

		// 这也就是说，当前
		if current.Left != nil {
			if parent.Val&1 == 0 {
				*ans += current.Left.Val
			}
			dfs(current.Left, current, ans)
		}
		if current.Right != nil {
			if parent.Val&1 == 0 {
				*ans += current.Right.Val
			}
			dfs(current.Right, current, ans)
		}
	}
	dfs(root.Left, root, &ans)
	dfs(root.Right, root, &ans)
	return ans
}
