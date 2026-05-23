package leetcode

func check(nums []int) bool {
	size := len(nums)
	_break := false

	for i := 1; i < size; i++ {
		if nums[i] >= nums[i-1] {
			continue
		}

		if _break {
			return false
		}
		_break = true
	}

	if !_break {
		return true
	}
	return nums[size-1] <= nums[0]
}
