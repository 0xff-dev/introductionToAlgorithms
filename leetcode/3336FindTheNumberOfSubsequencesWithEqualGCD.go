package leetcode

const MOD3336 = 1000000007
const MAXV = 200

var gcd3336Memo [MAXV + 1][MAXV + 1]int

func init() {
	for i := 0; i <= MAXV; i++ {
		for j := 0; j <= MAXV; j++ {
			gcd3336Memo[i][j] = gcd3336(i, j)
		}
	}
}

func gcd3336(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func subsequencePairCount(nums []int) int {
	dp := make([][]int, MAXV+1)
	for i := range dp {
		dp[i] = make([]int, MAXV+1)
	}
	dp[0][0] = 1

	for _, v := range nums {
		nextDP := make([][]int, MAXV+1)
		for i := range nextDP {
			nextDP[i] = make([]int, MAXV+1)
			copy(nextDP[i], dp[i])
		}

		for g1 := 0; g1 <= MAXV; g1++ {
			for g2 := 0; g2 <= MAXV; g2++ {
				count := dp[g1][g2]
				if count == 0 {
					continue
				}

				ng1 := gcd3336Memo[g1][v]
				nextDP[ng1][g2] = (nextDP[ng1][g2] + count) % MOD3336

				ng2 := gcd3336Memo[g2][v]
				nextDP[g1][ng2] = (nextDP[g1][ng2] + count) % MOD3336
			}
		}
		dp = nextDP
	}

	ans := 0
	for g := 1; g <= MAXV; g++ {
		ans = (ans + dp[g][g]) % MOD3336
	}

	return ans
}
