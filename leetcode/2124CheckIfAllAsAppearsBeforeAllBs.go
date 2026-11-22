package leetcode

func checkString(s string) bool {
	lastA, earliestB := -1, 1000000007
	for i := range s {
		if s[i] == 'a' {
			lastA = i
			continue
		}
		earliestB = min(earliestB, i)
	}
	return lastA < earliestB
}
