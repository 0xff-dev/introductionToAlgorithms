package leetcode

func islandPerimeter(grid [][]int) int {
	ans := 0
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				continue
			}
			// 检查四个方向
			if i == 0 || grid[i-1][j] == 0 {
				ans++
			}
			if i == rows-1 || grid[i+1][j] == 0 {
				ans++
			}
			if j == 0 || grid[i][j-1] == 0 {
				ans++
			}
			if j == cols-1 || grid[i][j+1] == 0 {
				ans++
			}
		}
	}
	return ans
}
