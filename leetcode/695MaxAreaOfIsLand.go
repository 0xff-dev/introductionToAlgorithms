package leetcode

func maxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var (
		dfs  func(int, int, *int)
		dirs = [][]int{
			{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		}
	)
	dfs = func(x, y int, ans *int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == 0 {
			return
		}
		*ans++
		grid[x][y] = 0
		for _, dir := range dirs {
			nx, ny := dir[0]+x, dir[1]+y
			dfs(nx, ny, ans)
		}
	}
	result := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				ans := 0
				dfs(r, c, &ans)
				if ans > result {
					result = ans
				}
			}
		}
	}
	return result
}
