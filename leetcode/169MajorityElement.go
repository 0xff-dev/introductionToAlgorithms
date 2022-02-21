package leetcode

//
func majorityElement(nums []int) int {
	times := 1
	target := nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == target {
			times++
			continue
		}
		times--
		if times == 0 {
			target = nums[idx]
			times = 1
		}
	}
	return target
}
