package leetcode

// 最坏的方法是，将所有的substring找出来，然后判断差值
func largestVariance(s string) int {
	// 找的过程中，滑动窗口
	l := len(s)
	/*
		// 表示在i之前有多少a，b，c，d
		count := make([][]int, l+1)
		for i := 0; i <= l; i++ {
			count[i] = make([]int, 26)
		}
		for i := 1; i <= l; i++ {
			copy(count[i], count[i-1])
			count[i][s[i-1]-'a']++
		}

		ans := 0
		// a a b a b b b
		for ws := 2; ws <= l; ws++ {
			for end := ws; end <= l; end += 1 {
				start := end - ws
				max, min := 0, 0
				firstMin := true
				for i := 0; i < 26; i++ {
					if count[end][i] == 0 {
						continue
					}
					diff := count[end][i] - count[start][i]
					if diff == 0 {
						// 元素不应该存在
						continue
					}
					if diff > max {
						max = diff
					}
					if firstMin || diff < min {
						firstMin = false
						min = diff
					}
				}
				if r := max - min; r > ans {
					ans = r
				}
			}
		}
		return ans
	*/
	count := make([]int, 26)
	for i := 0; i < l; i++ {
		count[s[i]-'a']++
	}
	gm := 0
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i == j || count[i] == 0 || count[j] == 0 {
				continue
			}
			major := 'a' + i
			minor := 'a' + j
			majorCount := 0
			minorCount := 0

			restMinor := count[j]
			for _, ch := range s {
				if byte(ch) == byte(major) {
					majorCount++
				}
				if byte(ch) == byte(minor) {
					minorCount++
					restMinor--
				}
				if minorCount > 0 && majorCount-minorCount > 0 {
					gm = majorCount - minorCount
				}
				if majorCount < minorCount && restMinor > 0 {
					majorCount = 0
					minorCount = 0
				}
			}
		}
	}
	return gm
}
