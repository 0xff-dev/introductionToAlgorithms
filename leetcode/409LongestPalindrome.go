package leetcode

func longestPalindrome(s string) int {
	Up := make([]int, 26)
	Low := make([]int, 26)
	for _, b := range []byte(s) {
		if b >= 65 && b <= 90 {
			Up[b-'A']++
			continue
		}
		Low[b-'a']++
	}

	count := 0

	hasOdd := false
	for _, u := range Up {
		isEven := u & 1
		if u > 0 {
			if isEven == 0 {
				count += u
				continue
			}
			hasOdd = true
			count += u - 1
		}
	}
	for _, l := range Low {
		isEven := l & 1
		if l > 0 {
			if isEven == 0 {
				count += l
				continue
			}
			hasOdd = true
			count += l - 1
		}
	}
	if hasOdd {
		count++
	}
	return count
}
