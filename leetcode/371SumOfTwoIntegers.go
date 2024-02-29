package leetcode

func getSum371(a int, b int) int {
	for b != 0 {
		// 找到相加应该进位得点
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}
