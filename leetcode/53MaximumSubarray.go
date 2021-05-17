package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	_max := nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if 0 < nums[idx-1] {
			nums[idx] += nums[idx-1]
		}
		if nums[idx] > _max {
			_max = nums[idx]
		}
	}

	return _max
}
