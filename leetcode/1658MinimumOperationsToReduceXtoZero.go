package leetcode

func minOperations(nums []int, x int) int {
	left := make(map[uint64]int)
	right := make(map[uint64]int)
	length := len(nums)
	base := uint64(0)
	for idx := 0; idx < length; idx++ {
		base += uint64(nums[idx])
		left[base] = idx
	}
	base = uint64(0)
	for idx := length - 1; idx >= 0; idx-- {
		base += uint64(nums[idx])
		right[base] = idx
	}
	ans := -1
	for leftSum, leftSteps := range left {
		if uint64(x) < leftSum {
			if rightSteps, ok := right[uint64(x)]; ok && (ans == -1 || length-rightSteps < ans) {
				ans = length - rightSteps
			}
			continue
		}
		if leftSum == uint64(x) && (ans == -1 || ans > leftSteps+1) {
			ans = leftSteps + 1
			continue
		}
		if rightSteps, ok := right[uint64(x)]; ok && (ans == -1 || length-rightSteps < ans) {
			ans = length - rightSteps
			continue
		}
		diff := uint64(x) - leftSum
		if rightSteps, ok := right[diff]; ok && rightSteps > leftSteps && (ans == -1 || ans > length-rightSteps+leftSteps+1) {
			ans = length - rightSteps + leftSteps + 1
		}
	}
	return ans
}
