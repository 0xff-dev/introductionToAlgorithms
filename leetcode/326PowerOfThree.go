package leetcode

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	// n = 3^x ==> x= log3 n ==> log(n)/log(3)
	// 看到另一个解法，找到int32范围最大幂次，然后确实mod
	return n > 0 && 1162261467%n == 0
}
