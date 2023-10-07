package leetcode

const mod1420 = 1000000007

func numOfArrays(n int, m int, k int) int {
	// 妈的，看着像是三维数组的dp呢?
	// 用前1-m的数字组成长度为1-n的数组，且满足k是1-k的搜索次数
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}
	// 动态规划
	// 1,2,3,4,5
	for i := 0; i <= m; i++ {
		dp[n][i][0] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := m; j >= 0; j-- {
			for o := 0; o <= k; o++ {
				ans := 0
				for kk := 1; kk <= j; kk++ {
					ans = (ans + dp[i+1][j][o]) % mod1420
				}
				if o > 0 {
					for kk := j + 1; kk <= m; kk++ {
						ans = (ans + dp[i+1][kk][o-1]) % mod1420
					}
				}
			}
		}
	}
	return dp[0][0][k]
}
