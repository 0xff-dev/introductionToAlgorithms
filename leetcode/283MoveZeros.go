package leetcode

func moveZeroes(nums []int) {

	if len(nums) == 0 {
		return
	}
	zeroIdx := 0
	notZeroIdx := 0
	for zeroIdx < len(nums) {
		if nums[zeroIdx] != 0 {
			zeroIdx++
			notZeroIdx = zeroIdx
			continue
		}
		next := notZeroIdx + 1
		for ; next < len(nums) && nums[next] == 0; next++ {
		}

		if next == len(nums) {
			break
		}

		nums[zeroIdx], nums[next] = nums[next], nums[zeroIdx]
		notZeroIdx = next
		zeroIdx++
	}
}
