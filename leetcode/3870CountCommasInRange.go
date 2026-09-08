package leetcode

func countCommas(n int) int {
	// 1000 -1999是多少个1000个逗号
	if n < 1000 {
		return 0
	}
	// 9999-1000+1 = 9000个
	bits, s := 0, n
	for ; n > 0; n /= 10 {
		bits++
	}
	base := 9000
	if bits == 4 {
		return s - 1000 + 1
	}
	return base + s - 10000 + 1
}
