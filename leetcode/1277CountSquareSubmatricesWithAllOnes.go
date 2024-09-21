package leetcode

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m+1)
	for i := range m + 1 {
		cache[i] = make([]int, n+1)
	}
	for r := m - 1; r >= 0; r-- {
		for c := n - 1; c >= 0; c-- {
			cache[r][c] = matrix[r][c] + cache[r+1][c] + cache[r][c+1] - cache[r+1][c+1]
		}
	}
	maxLen := min(m, n)
	ans := cache[0][0]
	for l := 2; l <= maxLen; l++ {
		target := l * l
		for i := 0; i <= m-l; i++ {
			for j := 0; j <= n-l; j++ {
				ni, nj := i+l, j+l
				if r := cache[i][j] - cache[ni][j] - cache[i][nj] + cache[ni][nj]; r == target {
					ans++
				}
			}
		}
	}

	return ans
}
