package leetcode

import "fmt"

// a=97
func smallestFromLeaf(root *TreeNode) string {
	var dfs func(*TreeNode, string)
	ans := ""
	dfs = func(now *TreeNode, total string) {
		if now == nil {
			return
		}
		self := fmt.Sprintf("%c", now.Val+97)

		if now.Left == nil && now.Right == nil {
			r := self + total
			if ans == "" || ans > r {
				ans = r
			}
		}
		dfs(now.Left, self+total)
		dfs(now.Right, self+total)
	}
	dfs(root, "")
	return ans
}
