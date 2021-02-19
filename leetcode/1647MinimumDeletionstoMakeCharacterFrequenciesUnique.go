package leetcode

func minDeletions(s string) int {
	const letters = 26
	deletions := 0
	chars := make([]int, letters)
	for _, c := range s {
		chars[c-'a']++
	}
	repeat := make(map[int]int)
	maxRepeatCount := 0
	// a:5, b:5, c:4, d:3, e:2, f:1
	for idx := 0; idx < letters; idx++ {
		if chars[idx] <= 0 {
			continue
		}
		if chars[idx] > maxRepeatCount {
			maxRepeatCount = chars[idx]
		}
		repeat[chars[idx]]++
	}

	// a:3 b:3, c:2
	for idx := 0; idx < letters; idx++ {
		count := chars[idx]
		if count == 0 {
			continue
		}
		repeatCount := repeat[count]
		if repeatCount == 1 {
			continue
		}
		i := minDefectIndex(repeat, count)
		deletions += count - i
		repeat[count]--
		if i != 0 {
			repeat[i] = 1
		}
	}

	return deletions
}

func minDefectIndex(repeat map[int]int, maxRepeatCount int) int {
	for idx := maxRepeatCount - 1; idx > 0; idx-- {
		if _, ok := repeat[idx]; !ok {
			return idx
		}
	}
	return 0
}
