package leetcode

// 我的方法就是直接硬算了，虽然过了，但是时间和速度确实过慢
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

// 这个思路参考halfrost
// 要想两头的⻓度最少，也就是中间这段的⻓度最大。
// 这样就转换成直接在数组 上使用滑动窗口求解，累加和等于一个固定值的连续最⻓的子数组
func minOperations1(nums []int, x int) int {
	total := 0
	length := len(nums)
	for idx := 0; idx < length; idx++ {
		total += nums[idx]
	}
	diff := total - x
	if diff < 0 {
		return -1
	}
	if diff == 0 {
		return length
	}

	left, right, sum, ans := 0, 0, 0, -1
	for right < length {
		if sum < diff {
			sum += nums[right]
			right++
		}
		for sum >= diff {
			if sum == diff && ans < right-left {
				ans = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if ans == -1 {
		return -1
	}
	return length - ans
}
