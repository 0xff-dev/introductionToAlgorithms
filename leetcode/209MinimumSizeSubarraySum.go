package leetcode

// 8, 3
// 3
func minSubArrayLen(target int, nums []int) int {
	ans := 0
	start, end, sum := 0, 0, 0
	for end < len(nums) {
		if sum < target {
			sum += nums[end]
			end++
		}
		for sum >= target {
			if ans == 0 || ans > end-start {
				ans = end - start
			}
			sum -= nums[start]
			start++
		}
	}
	return ans
}
