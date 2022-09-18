package leetcode

func checkIfPangram(sentence string) bool {
	bucket := make([]bool, 26)
	for _, b := range sentence {
		bucket[b-'a'] = true
	}
	for idx := 0; idx < 26; idx++ {
		if !bucket[idx] {
			return false
		}
	}

	return true
}