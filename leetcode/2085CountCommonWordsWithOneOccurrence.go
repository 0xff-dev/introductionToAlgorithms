package leetcode

func countWords(words1 []string, words2 []string) int {
	w1 := make(map[string]int)
	w2 := make(map[string]int)
	for _, word := range words1 {
		w1[word]++
	}
	for _, word := range words2 {
		w2[word]++
	}
	ans := 0
	for word, count := range w1 {
		if c, ok := w2[word]; ok && c == count && c == 1 {
			ans++
		}
	}
	return ans
}
