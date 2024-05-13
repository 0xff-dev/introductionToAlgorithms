package leetcode

func matrixScore(grid [][]int) int {
	// 前面的看到你过年时高位
	rows, cols := len(grid), len(grid[0])
	// 优先调整第一列
	for i := 0; i < rows; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < cols; j++ {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}
	for c := 1; c < cols; c++ {
		ones := 0
		for r := 0; r < rows; r++ {
			if grid[r][c] == 1 {
				ones++
			}
		}
		if ones < rows-ones {
			for r := 0; r < rows; r++ {
				grid[r][c] = 1 - grid[r][c]
			}
		}
	}

	ans := 0
	for i := 0; i < rows; i++ {
		cur := 0
		for j := 0; j < cols; j++ {
			cur = cur*2 + grid[i][j]
		}
		ans += cur
	}
	return ans
}
