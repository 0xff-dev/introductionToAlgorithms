package leetcode

// [apple], a, e
//func Constructor745(words []string) WordFilter {
//	return WordFilter{words: words, length: len(words)}
//}
//
//func (this *WordFilter) F(prefix string, suffix string) int {
//	for idx := this.length - 1; idx >= 0; idx-- {
//		if hasPrefix(this.words[idx], prefix) && hasSuffix(this.words[idx], suffix) {
//			return idx
//		}
//	}
//	return -1
//}
//
//func hasPrefix(str, prefix string) bool {
//	if len(str) >= len(prefix) {
//		i, j := 0, 0
//		for ; j < len(prefix) && str[i] == prefix[j]; i, j = i+1, j+1 {
//		}
//		return j == len(prefix)
//	}
//	return false
//}
//func hasSuffix(str, suffix string) bool {
//	if len(str) >= len(suffix) {
//		i, j := len(str)-1, len(suffix)-1
//		for ; j >= 0 && str[i] == suffix[j]; i, j = i-1, j-1 {
//		}
//		return j == -1
//	}
//	return false
//}

// 硬编码，应该思考利用前缀树，晚点在做一下。
type wordTrie struct {
	suffix map[string]struct{}
	prefix map[string]struct{}
}
type WordFilter struct {
	words []wordTrie
	cache map[string]int
}

func Constructor745(words []string) WordFilter {
	storeWords := make([]wordTrie, len(words))
	for idx, word := range words {
		w := wordTrie{suffix: make(map[string]struct{}, 10), prefix: make(map[string]struct{}, 10)}
		for split := 0; split < len(word); split++ {
			w.prefix[word[:split+1]] = struct{}{}
			w.suffix[word[split:]] = struct{}{}
		}
		storeWords[idx] = w
	}
	return WordFilter{words: storeWords, cache: map[string]int{}}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	if i, ok := this.cache[prefix+"-"+suffix]; ok {
		return i
	}
	for idx := len(this.words) - 1; idx >= 0; idx-- {
		_, prefixOk := this.words[idx].prefix[prefix]
		_, suffixOk := this.words[idx].suffix[suffix]
		if prefixOk && suffixOk {
			this.cache[prefix+"-"+suffix] = idx
			return idx
		}
	}
	this.cache[prefix+"-"+suffix] = -1
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
