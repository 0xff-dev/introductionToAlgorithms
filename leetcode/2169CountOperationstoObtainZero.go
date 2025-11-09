package leetcode

func countOperations(num1 int, num2 int) int {
	var ret int
	for num1 > 0 && num2 > 0 {
		ret++
		if num1 > num2 {
			num1 -= num2
			continue
		}
		num2 -= num1
	}
	return ret
}
