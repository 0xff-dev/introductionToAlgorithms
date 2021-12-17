package leetcode

func findTheDifference(s, t string) byte {
	letters := make([]int, 26)
	for _, b := range []byte(t) {
		letters[b-'a']++
	}
	for _, b := range []byte(s) {
		letters[b-'a']--
	}
	for idx := uint8(0); idx < 26; idx++ {
		if letters[idx] > 0 {
			return idx + 'a'
		}
	}
	return byte(' ')
}

func findTheDifference2(s, t string) byte {
	start := uint8(0)
	for _, b := range []byte(s) {
		start ^= b
	}
	for _, b := range []byte(t) {
		start ^= b
	}
	return byte(start)
}
