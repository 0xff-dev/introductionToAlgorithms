package leetcode

const mod629 = 1000000007

/* 这个已经ac，但是计算的部分c-j, j-i+1-j有大量重复计算，利用前缀
下面重新用滚动数据优化
func kInversePairs(n int, k int) int {
	if k == 0 {
		return 1
	}
	if k == 1 {
		return n - 1
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[1][0] = 1
	//dp[1][k+1] = 1 //计算综合
	// n=4，k=3
	// 看4放到哪里
	// **3多少种摆法已经计算来， 而4向左动一下，数量+1**
	// 另外一个问题
	// n=4,k=5
	// 这个时候，我们要计算dp[3]的5，4，3
	// 先多维，然后优化成滚动数组
	for i := 2; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if i > j {
				for c := 0; c <= j; c++ {
					dp[i][j] = (dp[i][j] + dp[i-1][c]) % mod629
				}
			} else {
				for c := j - i + 1; c <= j; c++ {
					dp[i][j] = (dp[i][j] + dp[i-1][c]) % mod629

				}
				// 例如4最多往前蹦3为，在蹦就到负数了,所以注意
				// j-i+1:j
			}
		}
	}
	return dp[n][k]
}
*/

func kInversePairs(n int, k int) int {
	dp := [2][]int{}
	sum := [2][]int{}
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, k+1)
		sum[i] = make([]int, k+1)
	}
	dp[1][0] = 1
	for i := 0; i <= k; i++ {
		sum[1][i] = 1
	}
	loop := 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= k; j++ {
			dp[1-loop][j] = 0
			if i > j {
				dp[1-loop][j] = (dp[1-loop][j] + sum[loop][j]) % mod629
			} else {
				dp[1-loop][j] = (dp[1-loop][j] + sum[loop][j] - sum[loop][j-i]) % mod629
			}

			add := 0
			if j > 0 {
				add = sum[1-loop][j-1]
			}
			sum[1-loop][j] = dp[1-loop][j] + add
		}
		loop = 1 - loop
	}
	return dp[loop][k]
}
