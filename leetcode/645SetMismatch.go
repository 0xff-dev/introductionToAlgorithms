package leetcode

// 1, 2, 2, 4 // 9
// 1, 2, 3, 4 // 10
// 1, 1, 3, 4 // 9
func findErrorNums(nums []int) []int {
	dup := -1
	for _, n := range nums {
		abs := n
		if abs < 0 {
			abs = -abs
		}
		if nums[abs-1] < 0 {
			dup = abs
			continue
		}
		nums[abs-1] *= -1
	}
	// 0, 1, 2, 3
	// 1, 2, 2, 4
	// -1,-2, 2, -4
	missing := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			missing = i + 1
			break
		}
	}
	return []int{dup, missing}
}
