package leetcode

func countCommas(n int) int {
	// 1000 -1999是多少个1000个逗号
	if n < 1000 {
		return 0
	}
	return n - 999
}
