package leetcode

func maxPower(s string) int {
	pre := byte(' ')
	preIndex := -1
	ans := 0
	i := 0
	for ; i < len(s); i++ {
		if s[i] != pre {
			ans = max(ans, i-preIndex-1)
			pre = s[i]
			preIndex = i
		}
	}
	ans = max(ans, i-preIndex-1)
	return ans + 1
}
