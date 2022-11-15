package leetcode

func addDigits(n int) int {
	if n == 0 {
		return 0
	}
	r := n % 9
	if r == 0 {
		return 9
	}
	return r
}
