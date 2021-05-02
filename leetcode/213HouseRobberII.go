package leetcode

// 这个问题就是第一个房子选与不选的问题，然后选择两个的最大值
func rob1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	r := simpleCal(nums[1:])
	r1 := simpleCal(nums[:length-1])
	if r < r1 {
		r = r1
	}
	return r
}

// 0 表示抢，1表示不强
func simpleCal(nums []int) int {
	length := len(nums)
	dp := make([][2]int, length)
	dp[0][0], dp[0][1] = nums[0], 0

	for idx := 1; idx < length; idx++ {
		dp[idx][0] = dp[idx-1][1] + nums[idx]
		dp[idx][1] = dp[idx-1][0]
		if dp[idx-1][1] > dp[idx][1] {
			dp[idx][1] = dp[idx-1][1]
		}
	}
	r := dp[length-1][0]
	if dp[length-1][1] > r {
		r = dp[length-1][1]
	}

	return r
}
