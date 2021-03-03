package leetcode

// 1：1 2：2 3：3
func climbStairs(n int) int {
	first, second := 1, 2
	if n == 1 {
		return first
	}
	if n == 2 {
		return second
	}
	for idx := 3; idx <= n; idx++ {
		next := first + second
		first = second
		second = next
	}
	return second
}
