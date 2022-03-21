package leetcode

func partitionLabels(s string) []int {
	partLength := make([]int, 0)
	charsLastIndex := make([]int, 26)
	for idx, b := range s {
		charsLastIndex[b-'a'] = idx
	}

	lastIdx := -1
	length := 0
	// e, cc, bb
	// cc e bb
	// cec bb
	for idx := 0; idx < len(s); idx++ {
		if charsLastIndex[s[idx]-'a'] > lastIdx {
			lastIdx = charsLastIndex[s[idx]-'a']
		}
		if idx == lastIdx {
			partLength = append(partLength, length+1)
			length = 0
			continue
		}
		length++
	}
	return partLength
}
