package leetcode

func partitionString(s string) int {
	/*
		chars := make([]int, 26)
		for _, b := range s {
			chars[b-'a']++
		}
		ans := 0
		for i := 0; i < 26; i++ {
			if chars[i] > 0 && ans < chars[i] {
				ans = chars[i]
			}
		}
		return ans
	*/
	// ab ac ab a
	chars := make([]bool, 26)
	ans := 1
	for _, b := range s {
		if !chars[b-'a'] {
			chars[b-'a'] = true
			continue
		}
		ans++
		for i := 0; i < 26; i++ {
			chars[i] = false
		}
		chars[b-'a'] = true
	}
	return ans
}
