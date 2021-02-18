package leetcode

func findSubstring(s string, words []string) []int {
	r := make([]int, 0)
	if len(words) == 0 {
		return r
	}

	wordLen := len(words[0])
	allWordLen := wordLen * len(words)
	wordMap := make(map[string]int)
	for _, w := range words {
		wordMap[w]++
	}

	for index := 0; index <= len(s)-allWordLen; index++ {
		existWord := 0
		tmpWord := make(map[string]int)
		for start := index; start < index+allWordLen; start += wordLen {
			tw := s[start : start+wordLen]
			tmpWord[tw]++
			count, ok := wordMap[tw]
			if !ok {
				break
			}
			if count == tmpWord[tw] {
				existWord++
			}
		}
		if existWord == len(wordMap) {
			r = append(r, index)
		}
	}
	return r
}
