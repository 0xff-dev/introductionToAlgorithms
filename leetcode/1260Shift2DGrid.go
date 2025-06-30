package leetcode

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	var do func()
	do = func() {
		newFirstColumn := []int{grid[m-1][n-1]}
		for i := 0; i < m-1; i++ {
			newFirstColumn = append(newFirstColumn, grid[i][n-1])
		}
		for r := m - 1; r >= 0; r-- {
			for c := n - 2; c >= 0; c-- {
				grid[r][c+1] = grid[r][c]
			}
		}
		for i := 0; i < len(newFirstColumn); i++ {
			grid[i][0] = newFirstColumn[i]
		}
	}
	for range k {
		do()
	}
	return grid
}
