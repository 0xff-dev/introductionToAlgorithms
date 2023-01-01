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

func wordPattern290(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	b2w := make(map[byte]string)
	w2b := make(map[string]byte)
	b, w := pattern[0], words[0]
	b2w[b] = w
	w2b[w] = b
	// abba, dog cat cat dog
	// abba dog, dog, dog, dog
	for i := 1; i < len(pattern); i++ {
		if pattern[i] == b {
			if words[i] != w {
				return false
			}
			continue
		}
		b, w = pattern[i], words[i]
		v, ok := b2w[pattern[i]]
		v1, ok1 := w2b[words[i]]
		if (ok && words[i] != v) || (ok1 && pattern[i] != v1) {
			return false
		}
		b2w[b] = w
	}
	return true
}
