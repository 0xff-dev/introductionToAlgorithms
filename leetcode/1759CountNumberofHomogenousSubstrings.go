package leetcode

const mod1795 = 1000000007

func countHomogenous(s string) int {
	ans := 0
	count := 1
	// a b b c c c
	for idx := 1; idx < len(s); idx++ {
		ans = (ans + count) % mod1795
		if s[idx] == s[idx-1] {
			count++
		} else {
			count = 1
		}
	}
	ans = (ans + count) % mod1795
	return ans
}
