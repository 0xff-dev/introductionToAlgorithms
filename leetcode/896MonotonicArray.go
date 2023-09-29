package leetcode

func isMonotonic(nums []int) bool {
	var judger func(func(a, b int) bool) bool
	judger = func(f func(a, b int) bool) bool {
		for idx := 1; idx < len(nums); idx++ {
			if !f(nums[idx], nums[idx-1]) {
				return false
			}
		}
		return true
	}
	return judger(func(a, b int) bool { return a >= b }) || judger(func(a, b int) bool { return a <= b })
}

func isMonotonic2(nums []int) bool {
	a, b := -1, -1
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == nums[idx-1] {
			continue
		}
		if nums[idx] > nums[idx-1] {
			if b != -1 {
				return false
			}
			a = 1
			continue
		}
		if nums[idx] < nums[idx-1] {
			if a != -1 {
				return false
			}
			b = 1
		}
	}
	return true
}
