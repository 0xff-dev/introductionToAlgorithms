package leetcode

func largeGroupPositions(s string) [][]int {
	ans := make([][]int, 0)
	start, end := 0, 0
	for ; end < len(s); end++ {
		if s[end] == s[start] {
			continue
		}
		if end-start >= 3 {
			ans = append(ans, []int{start, end - 1})
		}
		start = end
	}
	if end-start >= 3 {
		ans = append(ans, []int{start, end - 1})
	}
	return ans
}
