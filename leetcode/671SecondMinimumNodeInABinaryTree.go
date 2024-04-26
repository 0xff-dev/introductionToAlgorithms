package leetcode

import "sort"

// 应该可以递归解决
func findSecondMinimumValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	v := make([]int, 0)
	uv := map[int]struct{}{}

	for len(queue) > 0 {
		nq := make([]*TreeNode, 0)
		for _, cur := range queue {
			if _, ok := uv[cur.Val]; !ok {
				v = append(v, cur.Val)
				uv[cur.Val] = struct{}{}
			}
			if cur.Left != nil {
				nq = append(nq, cur.Left)
			}
			if cur.Right != nil {
				nq = append(nq, cur.Right)
			}
		}
		queue = nq
	}
	sort.Ints(v)
	if len(v) < 2 {
		return -1
	}
	return v[1]
}
