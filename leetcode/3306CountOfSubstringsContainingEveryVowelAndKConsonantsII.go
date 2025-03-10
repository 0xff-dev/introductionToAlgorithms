package leetcode

func isVowel3306(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func countOfSubstrings(word string, k int) int64 {
	var numValidSubstrings int64 = 0
	start := 0
	end := 0
	vowelCount := make(map[byte]int)
	consonantCount := 0

	// 预计算 nextConsonant 数组
	nextConsonant := make([]int, len(word))
	nextConsonantIndex := len(word)
	for i := len(word) - 1; i >= 0; i-- {
		nextConsonant[i] = nextConsonantIndex
		if !isVowel3306(word[i]) {
			nextConsonantIndex = i
		}
	}

	for end < len(word) {
		newLetter := word[end]

		if isVowel3306(newLetter) {
			vowelCount[newLetter]++
		} else {
			consonantCount++
		}

		// 缩小窗口，直到符合条件（consonantCount <= k）
		for consonantCount > k {
			startLetter := word[start]
			if isVowel3306(startLetter) {
				vowelCount[startLetter]--
				if vowelCount[startLetter] == 0 {
					delete(vowelCount, startLetter)
				}
			} else {
				consonantCount--
			}
			start++
		}

		// 满足所有元音出现且辅音数量为 k
		for start < len(word) && len(vowelCount) == 5 && consonantCount == k {
			numValidSubstrings += int64(nextConsonant[end] - end)

			startLetter := word[start]
			if isVowel3306(startLetter) {
				vowelCount[startLetter]--
				if vowelCount[startLetter] == 0 {
					delete(vowelCount, startLetter)
				}
			} else {
				consonantCount--
			}

			start++
		}

		end++
	}

	return numValidSubstrings
}
