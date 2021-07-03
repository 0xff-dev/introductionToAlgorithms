package leetcode

func jump(nums []int) int {
	dp := make([]int, len(nums))
	const inf = 0xffff
	for idx := 0; idx < len(nums); idx++ {
		dp[idx] = inf
	}
	dp[0] = 0
	for idx := 1; idx < len(nums); idx++ {
		for j := idx - 1; j >= 0; j-- {
			if idx-j <= nums[j] {
				if dp[idx] > dp[j]+1 {
					dp[idx] = dp[j] + 1
				}
			}
		}
	}
	return dp[len(nums)-1]
}
