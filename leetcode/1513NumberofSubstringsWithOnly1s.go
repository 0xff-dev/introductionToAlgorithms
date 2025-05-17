package leetcode

const mod1513 = 1e9 + 7

func numSub(s string) int {
	one := 0
	ans := 0
	for _, b := range s {
		if b == '0' {
			ans = (ans + one*(one+1)/2) % mod1513
			one = 0
			continue
		}
		one++
	}
	ans = (ans + one*(one+1)/2) % mod1513
	return ans
}
