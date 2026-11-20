package leetcode

import "sort"

func maxSelectedElements(nums []int) int {
	sort.Ints(nums)

	dp := make(map[int]int, 4)
	ans := 0

	for _, x := range nums {
		useX := dp[x-1] + 1 // 决策一：x 保持 x，前面必须结束在 x-1
		useX1 := dp[x] + 1  // 决策二：x 变成 x+1，前面必须结束在 x

		// 两个转移都基于【旧】dp 算完再写回，
		// 保证同一个 x 不会既当 x 又当 x+1（+1 只能用一次）。
		if v := dp[x]; useX > v {
			dp[x] = useX
		}
		if v := dp[x+1]; useX1 > v {
			dp[x+1] = useX1
		}

		if dp[x] > ans {
			ans = dp[x]
		}
		if dp[x+1] > ans {
			ans = dp[x+1]
		}

		// 优化：处理完 x 之后，所有键 < x-1 的 dp 永远不会再被访问
		//（之后只会查 x'-1, x', x'+1，而 x' >= x）。map 只会有 ≤4 个键。
		for k := range dp {
			if k < x-1 {
				delete(dp, k)
			}
		}
	}
	return ans
}
