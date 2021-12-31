package leetcode

func construct2DArray(original []int, m, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}

	result := make([][]int, m)
	idx := 0
	for row := 0; row < m; row++ {
		result[row] = make([]int, n)
		for in := 0; in < n; in++ {
			result[row][in] = original[idx+in]
		}
		idx += n
	}

	return result
}
