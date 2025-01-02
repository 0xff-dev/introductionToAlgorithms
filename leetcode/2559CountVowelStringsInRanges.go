package leetcode

func isVowel2559(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func vowelStrings(words []string, queries [][]int) []int {
	count := make([]int, len(words))
	pre := 0
	for i, w := range words {
		count[i] = pre
		if isVowel2559(w[0]) && isVowel2559(w[len(w)-1]) {
			count[i]++
		}
		pre = count[i]
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = count[q[1]]
		if q[0] > 0 {
			ans[i] -= count[q[0]-1]
		}
	}
	return ans
}
