package leetcode

func maxDepth1614(s string) int {
	ans := 0
	cur := 0
	for _, b := range s {
		if b == '(' {
			cur++
			if cur > ans {
				ans = cur
			}
		}
		if b == ')' {
			cur--
		}
	}
	return ans
}
