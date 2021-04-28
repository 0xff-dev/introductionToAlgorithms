package leetcode

import (
	"strings"
)

func reverse(s []byte) {
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}

}
func reverseWords(s string) string {
	r := []byte(strings.Trim(s, " "))
	reverse(r)
	result := make([]string, 0)
	start, end := 0, 0
	for start < len(r) {
		if r[start] == ' ' {
			start++
			continue
		}
		end = start + 1
		for ; end < len(r) && r[end] != ' '; end++ {
		}
		reverse(r[start:end])
		result = append(result, string(r[start:end]))
		start = end + 1
	}
	return strings.Join(result, " ")
}
