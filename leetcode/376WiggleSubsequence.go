package leetcode

func wiggleMaxLength(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	// 0, 以当前元素结尾切上升的情况，1以当前元素结尾且下降
	dp := make([][2]int, length)
	dp[0][0], dp[0][1] = 1, 1
	for idx := 1; idx < length; idx++ {
		if nums[idx] > nums[idx-1] {
			//上升的情况，依赖于前一个结尾是下降的
			dp[idx][0] = dp[idx-1][1] + 1
		} else {
			dp[idx][0] = dp[idx-1][0]
		}
		if nums[idx] < nums[idx-1] {
			dp[idx][1] = dp[idx-1][0] + 1
		} else {
			dp[idx][1] = dp[idx-1][1]
		}

	}
	r := dp[length-1][0]
	if dp[length-1][1] > r {
		r = dp[length-1][1]
	}
	return r
}
