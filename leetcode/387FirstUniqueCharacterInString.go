package leetcode

func firstUniqChar(s string) int {
	chars := make([]int, 26)
	bs := []byte(s)
	for _, b := range bs {
		chars[b-'a']++
	}

	for idx := range bs {
		if chars[bs[idx]-'a'] == 1 {
			return idx
		}
	}

	return -1

}
