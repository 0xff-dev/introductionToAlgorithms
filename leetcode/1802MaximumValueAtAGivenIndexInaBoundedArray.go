package leetcode

func getSum(index, value, n int) int {
	count := 0
	if value > index {
		count += (value + value - index) * (index + 1) / 2
	} else {
		count += (value+1)*value/2 + index - value + 1
	}
	if value >= n-index {
		count += (value + value - n + 1 + index) * (n - index) / 2
	} else {
		count += (value+1)*value/2 + n - index - value
	}
	return count - value
}
func maxValue(n int, index int, maxSum int) int {
	left, right := 1, maxSum
	for left < right {
		m := (left + right + 1) / 2
		if getSum(index, m, n) <= maxSum {
			left = m
		} else {
			right = m - 1
		}
	}
	return left
}
