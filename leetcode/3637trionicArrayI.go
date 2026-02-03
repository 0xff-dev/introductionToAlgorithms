package leetcode

func isTrionic(nums []int) bool {
	l := len(nums)
	incrEnd, decrStart := 1, l-1
	for ; incrEnd < l && nums[incrEnd] > nums[incrEnd-1]; incrEnd++ {
	}
	incrEnd--
	for ; decrStart > incrEnd && nums[decrStart-1] < nums[decrStart]; decrStart-- {

	}
	if incrEnd == 0 || decrStart == l-1 {
		return false
	}
	if incrEnd == decrStart {
		return false
	}
	for ; incrEnd < decrStart && nums[incrEnd] > nums[incrEnd+1]; incrEnd++ {

	}
	return incrEnd == decrStart
}
