package leetcode

import "strings"

func truncateSentence(s string, k int) string {
	sb := strings.Builder{}
	words := 0
	// aa bb cc k=3
	for idx := 0; idx < len(s); idx++ {
		if s[idx] == ' ' {
			words++
			if words == k {
				break
			}
		}
		sb.WriteByte(s[idx])
	}
	return sb.String()
}
