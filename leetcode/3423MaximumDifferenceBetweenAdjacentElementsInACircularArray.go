package leetcode

func maxAdjacentDistance(nums []int) int {
	var ans, tmp int
	l := len(nums)
	for i := 0; i < l; i++ {
		tmp = nums[i] - nums[(i-1+l)%l]
		if tmp < 0 {
			tmp = -tmp
		}
		ans = max(ans, tmp)
	}
	return ans
}
