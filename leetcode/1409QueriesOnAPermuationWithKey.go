package leetcode

func processQueries(queries []int, m int) []int {
	position := make([]int, m+1)
	for i := 1; i <= m; i++ {
		position[i] = i - 1
	}

	ans := make([]int, len(queries))
	for idx, query := range queries {
		ans[idx] = position[query]
		for i := 1; i <= m; i++ {
			if position[i] < ans[idx] {
				position[i]++
			}
		}
		position[query] = 0
	}
	return ans
}
