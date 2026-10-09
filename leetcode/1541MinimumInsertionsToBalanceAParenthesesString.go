package leetcode

func minInsertions1541(s string) int {
	ret, need := 0, 0
	for i := range s {
		if s[i] == '(' {
			need += 2
			if need&1 != 0 {
				ret++
				need--
			}
			continue
		}
		need--
		if need < 0 {
			ret++
			need += 2
		}
	}
	return ret + need
}
