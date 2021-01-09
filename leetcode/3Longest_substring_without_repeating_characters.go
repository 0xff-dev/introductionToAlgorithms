package leetcode

//
func lengthOfLongestSubstring(s string) int {
	r := 0
	if len(s) == 0 {
		return r
	}

	m := map[byte]int{}
	m[s[0]] = 0
	start, end := 0, 1
	for end < len(s) {
		if idx, ok := m[s[end]]; !ok {
			m[s[end]] = end
		} else {
			for i := start; i < idx; i++ {
				delete(m, s[i])
			}
			start = idx + 1
			m[s[end]], m[s[start]] = end, start
		}
		if len(m) > r {
			r = len(m)
		}
		end++
	}
	if len(m) > r {
		r = len(m)
	}
	return r
}

func lengthOfLongestSubstring1(s string) int {
	r := 0
	if len(s) == 0 {
		return r
	}
	if len(s) == 1 {
		return 1
	}

	m := map[byte]int{}
	m[s[0]] = 0
	start, end := 0, 1
	for end < len(s) {
		if idx, ok := m[s[end]]; !ok {
			m[s[end]] = end
		} else {
			m[s[end]] = end
			if idx+1 > start {
				start = idx+1
			}
		}
		if end-start+1 > r {
			r= end-start+1
		}
		end++
	}
	return r
}