package leetcode

const mod940 = 1000000007

func distinctSubseqII(s string) int {
	// 记录以x结尾的子序列长度，对于新增的来说，只要之前的就是重复
	dp := [26]int{}
	total := 0
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		// 出现s[i]后新增的不同子序列为total+1
		nc := (total + 1) % mod940
		// 如果s[i]出现过，那么就说明dp[c]多计算了
		total = (total + nc - dp[c] + mod940) % mod940
		dp[c] = nc
	}
	return total
}
