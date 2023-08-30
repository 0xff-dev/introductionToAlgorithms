package leetcode

func minimumReplacement(nums []int) int64 {
	ans := int64(0)
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			continue
		}
		a := int64(nums[i] + nums[i+1] - 1)
		var e int64 = a / int64(nums[i+1])

		ans += e - 1
		nums[i] = nums[i] / int(e)
	}
	return ans
}
