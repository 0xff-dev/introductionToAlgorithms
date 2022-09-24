package leetcode

func maxRepeating(sequence string, word string) int {
	step := len(word)

	match := 0
	ml := 0

	// abab ab
	for start := 0; start <= len(sequence)-step; start++ {
		i := start
		end := i + step
		match = 0
		for ; end <= len(sequence); end += step {
			now := sequence[i:end]
			if now == word {
				match, i = match+1, end
				continue
			}
			break
		}

		if match > ml {
			ml = match
		}
	}

	return ml
}
