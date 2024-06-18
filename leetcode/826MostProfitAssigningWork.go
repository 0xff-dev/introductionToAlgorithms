package leetcode

import (
	"sort"
)

type workProfit struct {
	d, p int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	// 需要找的是，小于x但是value最大的
	l := len(difficulty)
	dp := make([]workProfit, l)
	for i := 0; i < len(difficulty); i++ {
		dp[i] = workProfit{d: difficulty[i], p: profit[i]}
	}
	// 晚上回去改成线段树
	sort.Slice(dp, func(i, j int) bool {
		return dp[i].d < dp[j].d
	})

	ans := 0
	for _, n := range worker {
		idx := sort.Search(l, func(i int) bool {
			return dp[i].d > n
		})
		// [0, idx)在个区间找到最大值
		// 不行就上先段树
		m := 0
		for i := 0; i < idx; i++ {
			m = max(m, dp[i].p)
		}
		ans += m
	}

	return ans
}
