package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	minLen := -1
	for _, str := range strs {
		if minLen == -1 || len(str) < minLen {
			minLen = len(str)
		}
	}
	buf := strings.Builder{}
	for i := 0; i < minLen; i++ {
		c := strs[0][i]
		j := 1
		for ; j < len(strs); j++ {
			if strs[j][i] != c {
				break
			}
		}
		if j != len(strs) {
			break
		}
		buf.WriteByte(c)
	}
	return buf.String()
}
