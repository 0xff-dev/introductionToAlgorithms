package leetcode

func countVowelStrings(n int) int {

	res := 0
	countVowelStringsAux(n, 1, &res)
	return res
}

func countVowelStringsAux(n, last int, res *int) {
	if n == 0 {
		*res++
		return
	}
	for idx := last; idx <= 5; idx++ {
		countVowelStringsAux(n-1, idx, res)
	}
}
