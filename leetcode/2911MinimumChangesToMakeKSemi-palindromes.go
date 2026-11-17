package leetcode

import "math"

func minimumChanges(s string, k int) int {
	n := len(s)

	// 1. 预处理：找出每个长度的所有合法因数 d
	divisors := make([][]int, n+1)
	for length := 1; length <= n; length++ {
		for d := 1; d < length; d++ {
			if length%d == 0 {
				divisors[length] = append(divisors[length], d)
			}
		}
	}

	// 2. 预处理 g[i][j]：子串 s[i...j] 变成半回文串的最小修改次数
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt32 / 2
		}
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			for _, d := range divisors[length] {
				changes := 0
				// 按步长 d 分组，检查每组是否能通过修改变成回文串
				for offset := 0; offset < d; offset++ {
					l, r := 0, (length/d)-1
					for l < r {
						idx1 := i + offset + l*d
						idx2 := i + offset + r*d
						if s[idx1] != s[idx2] {
							changes++
						}
						l++
						r--
					}
				}
				if changes < g[i][j] {
					g[i][j] = changes
				}
			}
		}
	}

	// 3. 动态规划：dp[i][j] 表示前 i 个字符划分为 j 个半回文串所需的最小修改次数
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32 / 2
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			// 前 j-1 个部分至少需要 2 * (j - 1) 个字符，最后一部分至少需要 2 个字符（即 h <= i - 2）
			minH := 2 * (j - 1)
			if minH > i-2 {
				continue
			}
			for h := minH; h <= i-2; h++ {
				if dp[h][j-1] != math.MaxInt32/2 {
					cost := dp[h][j-1] + g[h][i-1]
					if cost < dp[i][j] {
						dp[i][j] = cost
					}
				}
			}
		}
	}

	return dp[n][k]
}
