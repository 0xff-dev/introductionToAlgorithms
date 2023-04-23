package challengeprogrammingdatastructure

func LongestIncreasingSubsequence(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	ans := 1
	for i := 1; i < len(nums); i++ {
		for pre := i - 1; pre >= 0; pre-- {
			if nums[pre] > nums[i] {
				continue
			}
			if r := dp[pre] + 1; r > dp[i] {
				dp[i] = r
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
