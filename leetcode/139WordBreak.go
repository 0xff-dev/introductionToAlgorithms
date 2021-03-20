package leetcode

import "strings"

//func wordBreak(s string, wordDict []string) bool {
//	wordMap := make(map[string]struct{})
//	for _, word := range wordDict {
//		wordMap[word] = struct{}{}
//	}
//	return wordBreakAux(s, wordMap)
//}

//func wordBreakAux(s string, wordMap map[string]struct{}) bool {
//	if _, ok := wordMap[s]; ok {
//		return true
//	}
//
//	for k := range wordMap {
//		if strings.HasPrefix(s, k) {
//			if wordBreakAux(s[len(k):], wordMap) {
//				return true
//			}
//		}
//	}
//	return false
//}
//
//
func wordBreak(s string, wordDict []string) bool {
	checkMap := make(map[int]bool)
	return wordBreakAux(s, 0, wordDict, checkMap)
}

// 递归的时候考虑缓存问题
func wordBreakAux(s string, idx int, words []string, check map[int]bool) bool {
	if idx == len(s) {
		check[idx] = true
		return true
	}
	if val, ok := check[idx]; ok {
		return val
	}
	for _, word := range words {
		if strings.HasPrefix(s[idx:], word) && wordBreakAux(s, idx+len(word), words, check) {
			check[idx] = true
			return true
		}
	}
	check[idx] = false
	return false
}
