package leetcode

var dirs = [][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	r := 0
	rows, cols := len(grid), len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				bfs(grid, row, col, rows, cols)
				r++
			}
		}
	}
	return r
}

func bfs(grid [][]byte, x, y, rows, cols int) {
	if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == '1' {
		grid[x][y] = '0'
		for _, dir := range dirs {
			nextX, nextY := x+dir[0], y+dir[1]
			bfs(grid, nextX, nextY, rows, cols)
		}
	}
}
