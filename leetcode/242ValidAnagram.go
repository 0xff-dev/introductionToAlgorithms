package leetcode

func isAnagram(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	letters := make([]int, 26)
	for idx := 0; idx < ls; idx++ {
		letters[s[idx]-'a']++
		letters[t[idx]-'a']--
	}
	for idx := 0; idx < 26; idx++ {
		if letters[idx] != 0 {
			return false
		}
	}
	return true
}
