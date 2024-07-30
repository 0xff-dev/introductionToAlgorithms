package leetcode

import "math"

func minimumDeletions1653(s string) int {
	/*
		func minimumDeletions(s string) int {
			// 首先找到A的最右边index
			// aababba b
			ll := len(s)
			ans := 0
			l, r := 0, len(s)-1
			for ; l < ll && s[l] == 'a'; l++ {
			}
			// aabbb
			// bbaaaaa bb
			for ; r >= 0 && s[r] == 'b'; r-- {
			}
			if r < l {
				return ans
			}
			a, b := 0, 0
			for i := l; i <= r; i++ {
				if s[i] == 'a' {
					a++
					continue
				}
				b++
			}
			return min(a, b)
		}
	*/
	l := len(s)
	ac, bc := make([]int, l), make([]int, l)
	b := 0
	for i := 0; i < l; i++ {
		if s[i] == 'b' {
			b++
		}
		bc[i] = b
	}
	a := 0
	for i := l - 1; i >= 0; i-- {
		if s[i] == 'a' {
			a++
		}
		ac[i] = a
	}
	ans := math.MaxInt
	for i := 0; i < l; i++ {
		ans = min(ans, ac[i]+bc[i])
	}
	return ans - 1
}
