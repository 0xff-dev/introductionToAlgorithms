package leetcode

var dir980 = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func uniquePathsIII(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]int, rows)
	steps := 0
	x, y := -1, -1
	ex, ey := -1, -1
	for row := 0; row < rows; row++ {
		visited[row] = make([]int, cols)
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 {
				steps++
			}
			if grid[row][col] == 1 {
				x, y = row, col
			}
			if grid[row][col] == 2 {
				ex, ey = row, col
			}
		}
	}
	visited[x][y] = 1
	ans := 0
	var dfs func(int, int, int)
	dfs = func(step, x, y int) {
		if step == steps+1 && x == ex && y == ey {
			ans++
			return
		}

		for _, dir := range dir980 {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols || visited[nx][ny] == 1 || grid[nx][ny] == -1 {
				continue
			}
			visited[nx][ny] = 1
			dfs(step+1, nx, ny)
			visited[nx][ny] = 0
		}
	}

	dfs(0, x, y)
	return ans
}
