package leetcode

func smallestRangeI(nums []int, k int) int {
	_max, _min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		_min = min(_min, nums[i])
		_max = max(_max, nums[i])
	}
	return max(0, _max-_min-2*k)
}
