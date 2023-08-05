package leetcode

// 1, 3, 4, 2, 2
// -, ,  -,
func findDuplicate(nums []int) int {
	for idx := range nums {
		x := nums[idx]
		if x < 0 {
			x = -x
		}
		index := x - 1
		r := nums[index]
		if r < 0 {
			return x
		}

		nums[index] = -nums[index]
	}
	return -1
}
