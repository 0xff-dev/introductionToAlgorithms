package leetcode

// 只包含a，b两个字符, 可以不连续的选择回文串，所以最坏的情况就是2，单选a+单选b
// 最好的情况就是s本身就是回文，直接返回1
func removePalindromeSub(s string) int {
	for s1, e := 0, len(s)-1; s1 < e; s1, e = s1+1, e-1 {
		if s[s1] != s[e] {
			return 2
		}
	}
	return 1
}
