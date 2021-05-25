package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	lenPattern := len(pattern)
	splitArr := strings.Split(s, " ")
	if len(splitArr) != lenPattern {
		return false
	}
	// abba, d d d d
	check := make(map[string]byte)
	patternCheck := make(map[byte]string)
	for idx := 0; idx < lenPattern; idx++ {
		word := splitArr[idx]
		p, ok := check[word]
		if !ok {
			tmpWord, exists := patternCheck[pattern[idx]]
			if !exists {
				check[word] = pattern[idx]
				patternCheck[pattern[idx]] = word
				continue
			}
			if tmpWord != word {
				return false
			}
			continue
		}
		if p != pattern[idx] {
			return false
		}
	}
	return true
}
