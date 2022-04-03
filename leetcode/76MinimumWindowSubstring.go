package leetcode

func minWindow(s string, t string) string {
	charCount := make(map[byte]int)

	result := ""
	if len(t) > len(s) {
		return result
	}
	for _, b := range t {
		charCount[byte(b)]++
	}

	windowCount := make(map[byte]int)
	left, right := 0, 0
	chars := 0
	for ; right < len(s); right++ {
		windowCount[s[right]]++
		if v, ok := charCount[s[right]]; ok && windowCount[s[right]] == v {
			chars++
		}
		for left <= right && chars == len(charCount) {
			if result == "" || right-left+1 < len(result) {
				result = s[left : right+1]
			}
			windowCount[s[left]]--
			if v, ok := charCount[s[left]]; ok {
				if windowCount[s[left]] < v {
					chars--
				}
			}
			left++
		}
	}
	return result
}
