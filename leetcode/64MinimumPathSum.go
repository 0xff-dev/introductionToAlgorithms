package leetcode

const INF = 0xffff

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	r := [2][]int{}
	for i := 0; i < 2; i++ {
		r[i] = make([]int, len(grid[0])+1)
		for j := 0; j < n+1; j++ {
			r[i][j] = INF
		}
	}
	r[0][0] = 0
	for row := 0; row < m; row++ {
		loop := row & 1
		for col := 1; col < n+1; col++ {
			min := r[loop][col-1]
			if min > r[1-loop][col] {
				min = r[1-loop][col]
			}
			r[loop][col] = min + grid[row][col-1]
		}
		r[0][0] = INF
	}
	return r[1-m&1][n]
}
