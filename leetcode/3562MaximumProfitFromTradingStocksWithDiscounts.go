package leetcode

func maxProfit3562(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	// 构建邻接表
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, 0)
	}
	for _, e := range hierarchy {
		g[e[0]-1] = append(g[e[0]-1], e[1]-1)
	}

	var dfs func(int) result
	dfs = func(u int) result {
		cost := present[u]
		dCost := present[u] / 2
		// dp[u][state][budget]
		// state = 0: Do not purchase parent node, state = 1: Must purchase parent node
		dp0 := make([]int, budget+1)
		dp1 := make([]int, budget+1)
		// subProfit[state][budget]
		// state = 0: discount not available, state = 1: discount available
		subProfit0 := make([]int, budget+1)
		subProfit1 := make([]int, budget+1)

		uSize := cost
		for _, v := range g[u] {
			childResult := dfs(v)
			uSize += childResult.size
			for i := budget; i >= 0; i-- {
				for sub := 0; sub <= min(childResult.size, i); sub++ {
					if i-sub >= 0 {
						subProfit0[i] = max(subProfit0[i], subProfit0[i-sub]+childResult.dp0[sub])
						subProfit1[i] = max(subProfit1[i], subProfit1[i-sub]+childResult.dp1[sub])
					}
				}
			}
		}

		for i := 0; i <= budget; i++ {
			dp0[i] = subProfit0[i]
			dp1[i] = subProfit0[i]
			if i >= dCost {
				dp1[i] = max(subProfit0[i], subProfit1[i-dCost]+future[u]-dCost)
			}
			if i >= cost {
				dp0[i] = max(subProfit0[i], subProfit1[i-cost]+future[u]-cost)
			}
		}

		return result{dp0, dp1, uSize}
	}

	return dfs(0).dp0[budget]
}

type result struct {
	dp0  []int
	dp1  []int
	size int
}
