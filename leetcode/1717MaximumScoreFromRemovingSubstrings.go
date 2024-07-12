package leetcode

func maximumGain(s string, x int, y int) int {
	// Ensure "ab" always has more points than "ba"
	if x < y {
		// Swap points
		x, y = y, x
		// Reverse the string to maintain logic
		s = reverseString1717(s)
	}

	aCount := 0
	bCount := 0
	totalPoints := 0

	for _, currentChar := range s {
		if currentChar == 'a' {
			aCount++
		} else if currentChar == 'b' {
			if aCount > 0 {
				// Can form "ab", remove it and add points
				aCount--
				totalPoints += x
			} else {
				// Can't form "ab", keep 'b' for potential future "ba"
				bCount++
			}
		} else {
			// Non 'a' or 'b' character encountered
			// Calculate points for any remaining "ba" pairs
			totalPoints += min(bCount, aCount) * y
			// Reset counters for next segment
			aCount = 0
			bCount = 0
		}
	}
	// Calculate points for any remaining "ba" pairs at the end
	totalPoints += min(bCount, aCount) * y

	return totalPoints
}

func reverseString1717(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

/*
func maximumGain(s string, x int, y int) int {
		start, end := 0, 0
		ans := 0
		cache := map[string]int{
			"": 0,
		}
		var dfs func(string) int
		dfs = func(str string) int {
			if v, ok := cache[str]; ok {
				return v
			}
			ans := 0
			for i := 0; i < len(str)-1; i++ {
				if str[i] == 'a' && str[i+1] == 'b' {
					ans = max(ans, dfs(str[:i]+str[i+2:])+x)
				}
			}
			for i := 0; i < len(str)-1; i++ {
				if str[i] == 'b' && str[i+1] == 'a' {
					ans = max(ans, dfs(str[:i]+str[i+2:])+y)
				}
			}
			cache[str] = ans
			return ans
		}

		for end < len(s) {
			if s[end] == 'a' || s[end] == 'b' {
				end++
				continue
			}
			if end-start > 1 {
				ans += dfs(s[start:end])
			}
			end++
			start = end
		}
		if end-start > 1 {
			ans += dfs(s[start:end])
		}

		return ans
}
*/
