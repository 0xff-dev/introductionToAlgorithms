package leetcode

func lengthOfLastWord(s string) int {
	r := 0
	idx := len(s) - 1
	for ; idx >= 0 && s[idx] == ' '; idx-- {
	}
	if idx < 0 {
		return r
	}
	for ; idx >= 0 && s[idx] != ' '; idx-- {
		r++
	}
	return r
}
