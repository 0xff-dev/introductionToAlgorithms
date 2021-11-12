package leetcode

func minStartValue(nums []int) int {
	_min := 1
	nums[0] += 1
	if nums[0] < 1 {
		_min = nums[0]
	}
	for idx := 1; idx < len(nums); idx++ {
		nums[idx] += nums[idx-1]
		if nums[idx] < 1 && _min > nums[idx] {
			_min = nums[idx]
		}
	}
	if _min == 1 {
		return 1
	}
	return 2 - _min
}
