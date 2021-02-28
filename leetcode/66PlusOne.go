package leetcode

func plusOne(digits []int) []int {
	length := len(digits) - 1
	digits[length] += 1
	cf := digits[length] / 10
	digits[length] %= 10
	for idx := length - 1; idx >= 0 && cf > 0; idx-- {
		digits[idx] = digits[idx] + cf
		cf = digits[idx] / 10

		digits[idx] %= 10
	}
	if cf != 0 {
		digits = append([]int{cf}, digits...)
	}
	return digits
}
