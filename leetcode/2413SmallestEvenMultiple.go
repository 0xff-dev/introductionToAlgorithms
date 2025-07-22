package leetcode

func smallestEvenMultiple(n int) int {
	if n&1 == 1 {
		return 2 * n
	}
	return n
}
