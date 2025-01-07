package leetcode

import "strings"

func stringMatching(words []string) []string {
	ans := make([]string, 0)
	for i := range words {
		for j := range words {
			if j == i {
				continue
			}
			if strings.Contains(words[j], words[i]) {
				ans = append(ans, words[i])
				break
			}
		}
	}
	return ans
}
